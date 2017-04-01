![TokenBalance](http://i.imgur.com/43Blvht.jpg)

# TokenBalance Server
Connects to your local geth IPC and prints out a simple JSON response for ethereum token balances. Runs on port *19705*. Fetch the balance of a token on the ethereum network.

#### Run Your Own Server

# Installation
#### Ubuntu 14.04/16.04 and Debian
```bash
wget -qO - https://deb.packager.io/key | sudo apt-key add -
echo "deb https://deb.packager.io/gh/hunterlong/tokenbalance trusty master" | sudo tee /etc/apt/sources.list.d/tokenbalance.list
sudo apt-get update
sudo apt-get install tokenbalance
```

#### CentOS
```bash
sudo rpm --import https://rpm.packager.io/key
echo "[tokenbalance]
name=Repository for hunterlong/tokenbalance application.
baseurl=https://rpm.packager.io/gh/hunterlong/tokenbalance/centos6/master
enabled=1" | sudo tee /etc/yum.repos.d/tokenbalance.repo
sudo yum install tokenbalance
```

# Run Server
```bash
tokenbalance --geth="/ethereum/geth.ipc" --port 8888 --ip 127.0.0.1
```
This will create a light weight HTTP server will respond balance information about a ethereum contract token.


#### Use TokenBalance's Sever
```bash
https://api.tokenbalance.com/balance/CONTRACT_HERE/ETHER_ADDRESS
```

#### CURL Request
```bash
CONTRACT=0xa74476443119A942dE498590Fe1f2454d7D4aC0d
ETH_ADDRESS=0xda0aed568d9a2dbdcbafc1576fedc633d28eee9a

curl https://api.tokenbalance.com/balance/$CONTRACT/$ETH_ADDRESS
```

![Screen1](http://i.imgur.com/252w4tG.png)


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
