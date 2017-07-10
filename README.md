![TokenBalance](http://i.imgur.com/43Blvht.jpg)

# TokenBalance Server
TokenBalance is an easy to use application to give you your ERC20 token balance without any troubles.
Connects to your local geth IPC and prints out a simple JSON response for ethereum token balances. Runs on port *8080*. Fetch the balance of a token on the ethereum network.

## TokenBalance API

#### Token Balance and Token Info
```bash
https://api.tokenbalance.com/balance/CONTRACT_HERE/ETHEREUM_ADDRESS
```
- ###### Response (JSON)
```bash
{
    "name": "TenX Pay Token",
    "wallet": "0x15b9360330e7be48d949c4710f01321735653f20",
    "symbol": "PAY",
    "balance": "10956.366853",
    "eth_balance": "0",
    "block": 4001224
}
```

#### Only Token Balance
```bash
https://api.tokenbalance.com/token/CONTRACT_HERE/ETHEREUM_ADDRESS
```
- ###### Response (PLAIN TEXT)
```bash
1022.503000
```

#### Run Your Own Server
TokenBalance isn't just an API, it's an opensource HTTP server that you can run on your own computer or server.

## Installation for Ubuntu
##### Ubuntu 16.04
```bash
wget -qO - https://deb.packager.io/key | sudo apt-key add -
echo "deb https://deb.packager.io/gh/hunterlong/tokenbalance xenial master" | sudo tee /etc/apt/sources.list.d/tokenbalance.list
```
##### Ubuntu 14.04
```bash
wget -qO - https://deb.packager.io/key | sudo apt-key add -
echo "deb https://deb.packager.io/gh/hunterlong/tokenbalance trusty master" | sudo tee /etc/apt/sources.list.d/tokenbalance.list
```

#### Install tokenbalance
```bash
sudo apt-get update
sudo apt-get install tokenbalance
```

## Run Your Own Server
```bash
tokenbalance start --geth="/ethereum/geth.ipc"
```
This will create a light weight HTTP server will respond balance information about a ethereum contract token.

## Optional Config
```bash
tokenbalance start --geth="/ethereum/geth.ipc" --port 8080 --ip 127.0.0.1
```

#### CURL Request
```bash
CONTRACT=0xa74476443119A942dE498590Fe1f2454d7D4aC0d
ETH_ADDRESS=0xda0aed568d9a2dbdcbafc1576fedc633d28eee9a

curl https://api.tokenbalance.com/balance/$CONTRACT/$ETH_ADDRESS
```

#### Response
```bash
{
    "name": "TenX Pay Token",
    "wallet": "0x15b9360330e7be48d949c4710f01321735653f20",
    "symbol": "PAY",
    "balance": "10956.366853",
    "eth_balance": "0",
    "block": 4001224
}
```

<p align="center">
  <img width="420" src="https://github.com/hunterlong/tokenbalance.com/blob/master/images/website.png?raw=true" alt="tokenbalance eth token contracts"/>
</p>

## Uninstalling with Ubuntu
```bash
sudo apt-get remove tokenbalance --purge
```
