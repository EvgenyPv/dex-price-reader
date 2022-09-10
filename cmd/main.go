package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
	"text/tabwriter"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	"dex-price-reader/contract-api/erc20"
	"dex-price-reader/contract-api/unifactory"
	"dex-price-reader/contract-api/unipair"
)

type dexStruct struct {
	dex0Name     string
	dex0PairAddr common.Address
	dex1Name     string
	dex1PairAddr common.Address
}
type tokenStruct struct {
	tkn0Addr        common.Address
	tkn0Symbol      string
	tkn0Decimals    uint8
	tkn0Denominator *big.Float
	tkn1Addr        common.Address
	tkn1Symbol      string
	tkn1Decimals    uint8
	tkn1Denominator *big.Float
}

type swapSides int64

const (
	sell swapSides = iota
	buy
)

type tradeStruct struct {
	price    float64
	size     float64
	swapSide swapSides
}

type blocksStruct struct {
	blocks map[uint64]uint64
	mu     sync.Mutex
}

func initParams() (*ethclient.Client, dexStruct, tokenStruct) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appKey := os.Getenv("ETH_APPKEY")
	rpcUrl := os.Getenv("ETH_APIADDRESS") + appKey

	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal(err)
	}

	//factory contracts instances are needed to find respective pair pool addresses
	dex0, err := unifactory.NewUnifactory(common.HexToAddress(os.Getenv("ETH_DEX0_FACTORY")), client)
	if err != nil {
		log.Fatal(err)
	}
	dex1, err := unifactory.NewUnifactory(common.HexToAddress(os.Getenv("ETH_DEX1_FACTORY")), client)
	if err != nil {
		log.Fatal(err)
	}
	//Tokens contract addresses to be analysed
	var (
		tokens tokenStruct
		dexes  dexStruct
	)
	if os.Getenv("ETH_TOKEN1") > os.Getenv("ETH_TOKEN0") {
		tokens.tkn0Addr = common.HexToAddress(os.Getenv("ETH_TOKEN0"))
		tokens.tkn1Addr = common.HexToAddress(os.Getenv("ETH_TOKEN1"))
	} else {
		tokens.tkn0Addr = common.HexToAddress(os.Getenv("ETH_TOKEN1"))
		tokens.tkn1Addr = common.HexToAddress(os.Getenv("ETH_TOKEN0"))
	}

	tkn0, err := erc20.NewErc20(tokens.tkn0Addr, client)
	if err != nil {
		log.Fatal(err)
	}
	tkn1, err := erc20.NewErc20(tokens.tkn1Addr, client)
	if err != nil {
		log.Fatal(err)
	}
	tokens.tkn0Symbol, err = tkn0.Symbol(nil)
	if err != nil {
		log.Fatal(err)
	}
	tokens.tkn1Symbol, err = tkn1.Symbol(nil)
	if err != nil {
		log.Fatal(err)
	}
	tokens.tkn0Decimals, err = tkn0.Decimals(nil)
	if err != nil {
		log.Fatal(err)
	}
	tokens.tkn0Denominator = new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(tokens.tkn0Decimals)), nil))
	tokens.tkn1Decimals, err = tkn1.Decimals(nil)
	if err != nil {
		log.Fatal(err)
	}
	tokens.tkn1Denominator = new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(tokens.tkn1Decimals)), nil))

	//get contract addresses of the pair pool at decentralized exchanges to read logs of swaps
	dexes.dex0PairAddr, err = dex0.GetPair(nil, tokens.tkn0Addr, tokens.tkn1Addr)
	if err != nil {
		log.Fatal(err)
	}
	dexes.dex0Name = os.Getenv("ETH_DEX0_NAME")
	dexes.dex1PairAddr, err = dex1.GetPair(nil, tokens.tkn0Addr, tokens.tkn1Addr)
	if err != nil {
		log.Fatal(err)
	}
	dexes.dex1Name = os.Getenv("ETH_DEX1_NAME")

	return client, dexes, tokens
}

func getBlockByTimestamp(client *ethclient.Client, targetTimestamp uint64) (*big.Int, error) {
	headerFirst, err := client.HeaderByNumber(context.Background(), big.NewInt(1))
	if err != nil {
		return nil, err
	}
	headerCurrent, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	averageTime := float64(headerCurrent.Time-headerFirst.Time) / float64(headerCurrent.Number.Uint64())

	for headerCurrent.Time > targetTimestamp {
		decreaseBlocks := int64(math.Round(float64(headerCurrent.Time-targetTimestamp) / averageTime))
		if decreaseBlocks < 1 {
			break
		}
		headerCurrent, err = client.HeaderByNumber(context.Background(), big.NewInt(headerCurrent.Number.Int64()-decreaseBlocks))
		if err != nil {
			return nil, err
		}

	}
	for (headerCurrent.Time + uint64(averageTime)) < targetTimestamp {
		headerCurrent, err = client.HeaderByNumber(context.Background(), big.NewInt(headerCurrent.Number.Int64()+1))
		if err != nil {
			return nil, err
		}
	}

	return headerCurrent.Number, nil
}

func getLogs(client *ethclient.Client, pairAddr common.Address, tokens tokenStruct,
	fromBlock *big.Int) (map[uint64][]tradeStruct, error) {
	//Query all Swap events (without filterting by sender/to) for a given pair pool address
	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		Addresses: []common.Address{
			pairAddr,
		},
		Topics: [][]common.Hash{
			{crypto.Keccak256Hash([]byte("Swap(address,uint256,uint256,uint256,uint256,address)"))},
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, err
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(unipair.UnipairABI)))
	if err != nil {
		return nil, err
	}

	var tradingData map[uint64][]tradeStruct
	tradingData = make(map[uint64][]tradeStruct)

	for _, vLog := range logs {
		swapEvent, err := contractAbi.Unpack("Swap", vLog.Data)
		if err != nil {
			return nil, err
		}
		//Below we cast amounts to big.Float, divide them using token denominator and then cast to float64
		amount0In, _ := new(big.Float).Quo(new(big.Float).SetInt(swapEvent[0].(*big.Int)), tokens.tkn0Denominator).Float64()
		amount1In, _ := new(big.Float).Quo(new(big.Float).SetInt(swapEvent[1].(*big.Int)), tokens.tkn1Denominator).Float64()
		amount0Out, _ := new(big.Float).Quo(new(big.Float).SetInt(swapEvent[2].(*big.Int)), tokens.tkn0Denominator).Float64()
		amount1Out, _ := new(big.Float).Quo(new(big.Float).SetInt(swapEvent[3].(*big.Int)), tokens.tkn1Denominator).Float64()
		var tradeInfo tradeStruct
		if amount0In > 0 {
			tradeInfo.swapSide = sell
			if tokens.tkn0Decimals > tokens.tkn1Decimals {
				tradeInfo.price = math.Round(amount1Out/amount0In*100) / 100
				tradeInfo.size = math.Round(amount0In*100) / 100
			} else {
				tradeInfo.price = math.Round(amount0In/amount1Out*100) / 100
				tradeInfo.size = math.Round(amount1Out*100) / 100
			}

		} else {
			tradeInfo.swapSide = buy
			if tokens.tkn0Decimals > tokens.tkn1Decimals {
				tradeInfo.price = math.Round(amount1In/amount0Out*100) / 100
				tradeInfo.size = math.Round(amount0Out*100) / 100
			} else {
				tradeInfo.price = math.Round(amount0Out/amount1In*100) / 100
				tradeInfo.size = math.Round(amount1In*100) / 100
			}
		}

		tradingData[vLog.BlockNumber] = append(tradingData[vLog.BlockNumber], tradeInfo)

	}

	return tradingData, nil

}

func getBlocksTime(client *ethclient.Client, dex0Trades map[uint64][]tradeStruct,
	dex1Trades map[uint64][]tradeStruct) map[uint64]uint64 {

	var wg sync.WaitGroup
	blocksTime := blocksStruct{blocks: make(map[uint64]uint64)}
	for blockNum := range dex0Trades {
		if _, ok := dex1Trades[blockNum]; ok {
			wg.Add(1)
			go func(blockNum uint64, blocksTime *blocksStruct, wg *sync.WaitGroup) {
				defer wg.Done()
				block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(blockNum)))
				blocksTime.mu.Lock()
				if err != nil {
					blocksTime.blocks[blockNum] = 0
				} else {
					blocksTime.blocks[blockNum] = block.Time()
				}
				blocksTime.mu.Unlock()
			}(blockNum, &blocksTime, &wg)
		}
	}

	wg.Wait()

	return blocksTime.blocks
}

func logSynchronousSwaps(dex0Trades map[uint64][]tradeStruct, dex0Name string,
	dex1Trades map[uint64][]tradeStruct, dex1Name string, blocksTime map[uint64]uint64) {

	var (
		buyStringDEX0  string
		buyStringDEX1  string
		sellStringDEX0 string
		sellStringDEX1 string
	)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	for blockNum, dex0Slice := range dex0Trades {
		if dex1Slice, ok := dex1Trades[blockNum]; ok {
			buyStringDEX0, buyStringDEX1, sellStringDEX0, sellStringDEX1 = "", "", "", ""
			for _, dex0Swap := range dex0Slice {
				if dex0Swap.swapSide == buy {
					buyStringDEX0 = buyStringDEX0 + fmt.Sprintf("Buy\t"+dex0Name+"\t%.2f\t%.2f\t\r\n", dex0Swap.price, dex0Swap.size)
				} else {
					sellStringDEX0 = sellStringDEX0 + fmt.Sprintf("Sell\t"+dex0Name+"\t%.2f\t%.2f\t\r\n", dex0Swap.price, dex0Swap.size)
				}
			}
			for _, dex1Swap := range dex1Slice {
				if dex1Swap.swapSide == buy {
					buyStringDEX1 = buyStringDEX1 + fmt.Sprintf("Buy\t"+dex1Name+"\t%.2f\t%.2f\t\r\n", dex1Swap.price, dex1Swap.size)
				} else {
					sellStringDEX1 = sellStringDEX1 + fmt.Sprintf("Sell\t"+dex1Name+"\t%.2f\t%.2f\t\r\n", dex1Swap.price, dex1Swap.size)
				}
			}
			if (len(buyStringDEX0) > 0 && len(buyStringDEX1) > 0) || (len(sellStringDEX0) > 0 && len(sellStringDEX1) > 0) {
				fmt.Fprintln(w, time.Unix(int64(blocksTime[blockNum]), 0).Format(time.Stamp)+"\tDEX\tPrice\tSize\t")
				if len(buyStringDEX0) > 0 && len(buyStringDEX1) > 0 {
					fmt.Fprintf(w, buyStringDEX0+buyStringDEX1)
				}
				if len(sellStringDEX0) > 0 && len(sellStringDEX1) > 0 {
					fmt.Fprintf(w, sellStringDEX0+sellStringDEX1)
				}
			}
		}
	}

	w.Flush()

}

func main() {

	fmt.Println("Initializing DEX and tokens data")
	client, dexes, tokens := initParams()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter analysis depth in hours: ")
	input, _ := reader.ReadString('\n')
	inputNum := strings.ReplaceAll(input, "\r\n", "")
	duration, err := strconv.ParseInt(inputNum, 10, 64)
	if err != nil || duration <= 0 {
		log.Fatal("Input must be postive integer")
	}
	targetTimestamp := uint64(time.Now().Unix() - duration*60*60) //user has input duration in hours

	//we will analyse blocks from startBlock (defined based on the input from user) to the latest
	fmt.Println("Finding block number by timestamp")
	startBlock, err := getBlockByTimestamp(client, targetTimestamp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Reading swap logs")
	dex0Trades, err := getLogs(client, dexes.dex0PairAddr, tokens, startBlock)
	if err != nil {
		log.Fatal(err)
	}
	dex1Trades, err := getLogs(client, dexes.dex1PairAddr, tokens, startBlock)
	if err != nil {
		log.Fatal(err)
	}

	blocksTime := getBlocksTime(client, dex0Trades, dex1Trades)

	logSynchronousSwaps(dex0Trades, dexes.dex0Name, dex1Trades, dexes.dex1Name, blocksTime)

}
