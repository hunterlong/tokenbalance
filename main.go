package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mkideal/cli"
	"log"
	"math/big"
	"os"
	"strconv"
)

var (
	conn         *ethclient.Client
	GethLocation string
	UsePort      string
	UseIP        string
	VERSION      string
	decimals     uint8
)

var help = cli.HelpCommand("display help information")

func main() {

	if err := cli.Root(root,
		cli.Tree(help),
		cli.Tree(child),
		cli.Tree(versionCli),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

type rootT struct {
	cli.Helper
}

var root = &cli.Command{
	Desc: "\n      #######################\n" +
		"           TokenBalance\n" +
		"      #######################\n\n" +
		"TokenBalance is an easy to use server that \n" +
		"give you your ERC20 token balance without \n" +
		"any troubles. Connects to your local geth \n" +
		"IPC and prints out a simple JSON response \n" +
		"for ethereum token balances.",
	Argv: func() interface{} { return new(rootT) },
	Fn: func(ctx *cli.Context) error {

		ctx.String("To start the tokenbalance server, use command:\ntokenbalance start --geth \"/root/ethereum/geth.ipc\" --port 8080 --ip 0.0.0.0\n * replace geth location with your own *\n")
		return nil
	},
}

var child = &cli.Command{
	Name: "start",
	Desc: "run the tokenbalance http server",
	Argv: func() interface{} { return new(argT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		GethLocation = argv.Geth
		UsePort = argv.Port
		UseIP = argv.IP
		ConnectGeth()
		StartServer()

		return nil
	},
}

var versionCli = &cli.Command{
	Name: "version",
	Desc: "get the version of tokenbalance server",
	Argv: func() interface{} { return new(argT) },
	Fn: func(ctx *cli.Context) error {
		ctx.String(VERSION + "\n")
		return nil
	},
}

type argT struct {
	cli.Helper
	Geth string `cli:"*g,geth" usage:"attach geth IPC or HTTP location"`
	IP   string `cli:"ip" usage:"Bind to IP Address" dft:"0.0.0.0"`
	Port string `cli:"p,port" usage:"HTTP server port for token information in JSON" dft:"8080"`
}

func ConnectGeth() {
	var err error
	conn, err = ethclient.Dial(GethLocation)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v\n", err)
	} else {
		log.Printf("Connected to Geth at: %v\n", GethLocation)
	}
}

func GetAccount(contract string, wallet string) (*BalanceResponse, error) {
	var err error

	response := new(BalanceResponse)

	response.Wallet = common.HexToAddress(wallet)

	token, err := NewTokenCaller(common.HexToAddress(contract), conn)
	if err != nil {
		log.Printf("Failed to instantiate a Token contract: %v\n", err)
		return nil, err
	}

	response.Block, err = conn.BlockByNumber(context.TODO(), nil)
	if err != nil {
		log.Printf("Failed to get current block number: %v\n", err)
		response.Block = nil
	}

	response.Decimals, err = token.Decimals(nil)
	if err != nil {
		log.Printf("Failed to get decimals from contract: %v \n", contract)
		return nil, err
	}

	response.EthBalance, err = conn.BalanceAt(context.TODO(), response.Wallet, nil)
	if err != nil {
		log.Printf("Failed to get ethereum balance from address: %v \n", response.Wallet)
	}

	response.Balance, err = token.BalanceOf(nil, response.Wallet)
	if err != nil {
		log.Printf("Failed to get balance from contract: %v %v\n", contract, err)
	}

	response.Symbol, err = token.Symbol(nil)
	if err != nil {
		log.Printf("Failed to get symbol from contract: %v \n", contract)
		response.Symbol = SymbolFix(contract)
	}

	response.Name, err = token.Name(nil)
	if err != nil {
		log.Printf("Failed to retrieve token name from contract: %v | %v\n", contract, err)
		response.Name = "MISSING"
	}

	return response, err
}

func SymbolFix(contract string) string {
	switch common.HexToAddress(contract).String() {
	case "0x86Fa049857E0209aa7D9e616F7eb3b3B78ECfdb0":
		return "EOS"
	}
	return "MISSING"
}

type jsonResponse struct {
	Name       string `json:"name,omitempty"`
	Wallet     string `json:"wallet,omitempty"`
	Symbol     string `json:"symbol,omitempty"`
	Balance    string `json:"balance"`
	EthBalance string `json:"eth_balance,omitempty"`
	Decimals   int64  `json:"decimals,omitempty"`
	Block      int64  `json:"block,omitempty"`
}

func (b *BalanceResponse) Format() *jsonResponse {
	return &jsonResponse{
		b.Name,
		b.Wallet.String(),
		b.Symbol,
		BigIntDecimal(b.Balance, b.Decimals.Int64()),
		BigIntDecimal(b.EthBalance, 18),
		b.Decimals.Int64(),
		b.Block.Number().Int64(),
	}
}

func (b *BalanceResponse) Ok() bool {
	if b.Decimals.Sign() >= 0 && b.Balance.Sign() >= 0 {
		return true
	}
	return false
}

func BigIntDecimal(balance *big.Int, decimals int64) string {
	if balance.Sign() == 0 {
		return "0"
	}
	bal := big.NewFloat(0)
	bal.SetInt(balance)
	pow := BigPow(10, decimals)
	p := big.NewFloat(0)
	p.SetInt(pow)
	bal.Quo(bal, p)
	deci := strconv.Itoa(int(decimals))
	dec := "%." + deci + "f"
	newNum := Clean(fmt.Sprintf(dec, bal))
	return newNum
}

func Clean(newNum string) string {
	stringBytes := bytes.TrimRight([]byte(newNum), "0")
	newNum = string(stringBytes)
	if stringBytes[len(stringBytes)-1] == 46 {
		newNum += "0"
	}
	if stringBytes[0] == 46 {
		newNum = "0" + newNum
	}
	return newNum
}

func BigPow(a, b int64) *big.Int {
	r := big.NewInt(a)
	return r.Exp(r, big.NewInt(b), nil)
}
