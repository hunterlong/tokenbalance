![TokenBalance](http://i.imgur.com/43Blvht.jpg)

# TokenBalance
Fetch the balance of a token on the ethereum network.

```bash
https://api.tokenbalance.com/balance/CONTRACT_HERE/ETHER_ADDRESS
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
  "name": "Golem Network Token",
  "wallet": "0xda0aed568d9a2dbdcbafc1576fedc633d28eee9a",
  "balance": 5.581543382215305,
  "eth_balance": 53.723423456,
  "symbol": "GNT",
  "decimals": 18
}
```

<p align="center">
  <img width="420" src="https://github.com/hunterlong/tokenbalance.com/blob/master/images/website.png?raw=true" alt="tokenbalance eth token contracts"/>
</p>


![Screen1](http://i.imgur.com/252w4tG.png)
