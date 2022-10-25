# Crypto DEX-price-reader
Sample project for DEX prices reading on Go.

The tool shows trades on different decentralized exchanges within one block.

Big difference between prcies on different exchanges shows potential opportunity for arbitrage trading.

# .env file example
```shell
ETH_APIADDRESS = "https://eth-mainnet.g.alchemy.com/v2/"
ETH_APPKEY = "enter-your-alchemy-app-key"
ETH_DEX0_NAME = "Sushiswap"
ETH_DEX0_FACTORY = "0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac"
ETH_DEX1_NAME = "Uniswap"
ETH_DEX1_FACTORY = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
ETH_TOKEN0 = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
ETH_TOKEN1 = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
```

# Output example
```shell
 Sep 10 16:42:26|       DEX|   Price| Size|
             Buy| Sushiswap| 1719.06| 1.37|
             Buy|   Uniswap| 1718.95| 0.16|
 Sep 10 16:46:26|       DEX|   Price| Size|
             Buy| Sushiswap| 1718.62| 0.16|
             Buy|   Uniswap| 1718.15| 0.57|
 Sep 10 17:22:32|       DEX|   Price| Size|
            Sell| Sushiswap| 1719.95| 0.35|
            Sell|   Uniswap| 1724.97| 0.47|
 Sep 10 16:44:10|       DEX|   Price| Size|
             Buy| Sushiswap| 1718.78| 1.21|
             Buy|   Uniswap| 1718.99| 1.21|
 Sep 10 17:10:38|       DEX|   Price| Size|
            Sell| Sushiswap| 1717.57| 0.40|
            Sell|   Uniswap| 1724.70| 0.05|
 Sep 10 17:18:09|       DEX|   Price| Size|
            Sell| Sushiswap| 1719.78| 0.05|
            Sell|   Uniswap| 1724.91| 0.58|
 Sep 10 17:01:41|       DEX|   Price| Size|
            Sell| Sushiswap| 1724.22| 0.06|
            Sell|   Uniswap| 1723.89| 2.88|
 Sep 10 17:15:00|       DEX|   Price| Size|
            Sell| Sushiswap| 1719.75| 0.20|
            Sell|   Uniswap| 1724.73| 0.08|
 Sep 10 17:23:04|       DEX|   Price| Size|
            Sell| Sushiswap| 1720.11| 1.05|
            Sell| Sushiswap| 1720.61| 3.24|
            Sell| Sushiswap| 1721.23| 2.05|
            Sell|   Uniswap| 1725.02| 0.37|
 Sep 10 17:10:53|       DEX|   Price| Size|
             Buy| Sushiswap| 1707.10| 2.00|
             Buy|   Uniswap| 1714.34| 0.52|
 Sep 10 16:43:11|       DEX|   Price| Size|
             Buy| Sushiswap| 1718.90| 0.01|
             Buy|   Uniswap| 1719.05| 0.04|
 Sep 10 17:21:32|       DEX|   Price| Size|
            Sell| Sushiswap| 1719.89| 0.19|
            Sell|   Uniswap| 1724.90| 0.03|
 Sep 10 17:36:04|       DEX|   Price| Size|
            Sell| Sushiswap| 1722.11| 0.09|
            Sell|   Uniswap| 1724.80| 0.01|
            Sell|   Uniswap| 1724.81| 0.17|
 ```
