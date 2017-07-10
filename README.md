![TokenBalance](http://i.imgur.com/43Blvht.jpg)

# TokenBalance API
TokenBalance is an easy to use public API and application that will output your [ERC20 Token](https://github.com/ConsenSys/Tokens/blob/master/Token_Contracts/contracts/Token.sol) balance without any troubles. You can run TokenBalance on your local computer or you can use api.tokenbalance.com to easily parse your erc20 token balances.
Connects to your local geth IPC and prints out a simple JSON response for ethereum token balances. Runs on port *8080* by default if you wish to run locally.

## Token Balance and Token Info (/token)
To fetch information about your balance, token details, and ETH balance use the follow API call in a simple HTTP GET or CURL. The response is in JSON so you can easily parse what you need. Replace TOKEN_ADDRESS with the contract address of the ERC20 token, and replace ETH_ADDRESS with your address.

```bash
https://api.tokenbalance.com/token/TOKEN_ADDRESS/ETH_ADDRESS
```
- ###### Response (JSON)
```bash
{
    "name": "TenX Pay Token",
    "wallet": "0x15b9360330e7be48d949c4710f01321735653f20",
    "symbol": "PAY",
    "balance": "10956.366853",
    "eth_balance": "3.75",
    "block": 4001224
}
```

## Only Token Balance (/balance)
This API response will only show you the ERC20 token balance in plain text. Perfect for ultra simple parsing.

```bash
https://api.tokenbalance.com/balance/TOKEN_ADDRESS/ETH_ADDRESS
```
- ###### Response (PLAIN TEXT)
```bash
1022.503000
```

## Examples

- [Fetch Balance and Token Details for Status Coin](https://api.tokenbalance.com/token/0x744d70FDBE2Ba4CF95131626614a1763DF805B9E/0x242f3f8cffecc870bdb30165a0cb3c1f06f32949)
- [Fetch Balance and Token Details for Gnosis](https://api.tokenbalance.com/token/0x6810e776880c02933d47db1b9fc05908e5386b96/0x97b47ffde901107303a53630d28105c6a7af1c3e)
- [Fetch Balance and Token Details for Storj](https://api.tokenbalance.com/token/0xb64ef51c888972c908cfacf59b47c1afbc0ab8ac/0x29b532092fd5031b9ee1e5fe07d627abedd5eda8)
- [Only Token Balance for Augur](https://api.tokenbalance.com/balance/0x48c80F1f4D53D5951e5D5438B54Cba84f29F32a5/0x90fbfc09db2f4b6e8b65b7a237e15bba9dc5db0c)
- [Only Token Balance for Golem](https://api.tokenbalance.com/balance/0xa74476443119A942dE498590Fe1f2454d7D4aC0d/0xe42b94dc4b02edef833556ede32757cf2b6cc455)


## Implement in your App
Feel free to use the TokenBalance API server to fetch ERC20 token balances and details. We do have a header set that will allow you to call the API via AJAX. `Access-Control-Allow-Origin "*"` The server may limit your requests if you do more than 60 hits per minute.

# Run Your Own Server
TokenBalance isn't just an API, it's an opensource HTTP server that you can run on your own computer or server.

## Installation
##### Ubuntu 16.04
```bash
git clone https://github.com/hunterlong/tokenbalance
cd tokenbalance
go get && go build .
```

## Start TokenBalance Server
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

curl https://api.tokenbalance.com/token/$CONTRACT/$ETH_ADDRESS
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
