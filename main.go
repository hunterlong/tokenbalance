package main

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
)

var conn *ethclient.Client

var GethLocation string
var UsePort string
var UseIP string

func ConnectGeth() {

	var err error

	//
	// MAC: /Users/username/Library/Ethereum/geth.ipc
	//
	// LINUX: /root/.ethereum/geth.ipc
	//

	conn, err = ethclient.Dial(GethLocation)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	} else {
		log.Println("Connected to Geth at: ", GethLocation)
	}

}

func GetAccount(contract string, wallet string) (string, float64, string, uint8, float64, uint64, error) {
	var err error

	token, err := NewTokenCaller(common.HexToAddress(contract), conn)
	if err != nil {
		log.Println("Failed to instantiate a Token contract: %v", err)
		return "error", 0, "error", 0, 0, 0, err
	}

	getBlock, err := conn.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Println("Failed to get current block number: ", err)
		return "error", 0, "error", 0, 0, 0, err
	}

	maxBlock := getBlock.NumberU64()

	address := common.HexToAddress(wallet)
	if err != nil {
		log.Println("Failed hex address: "+wallet, err)
		return "error", 0, "error", 0, 0, 0, err
	}

	ethAmount, err := conn.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Println("Failed to get ethereum balance from address: ", address, err)
		return "error", 0, "error", 0, 0, 0, err
	}

	balance, err := token.BalanceOf(nil, address)
	if err != nil {
		log.Println("Failed to get balance from contract: "+contract, err)
		return "error", 0, "error", 0, 0, 0, err
	}
	symbol, err := token.Symbol(nil)
	if err != nil {
		log.Println("Failed to get symbol from contract: "+contract, err)
		return "error", 0, "error", 0, 0, 0, err
	}
	decimals, err := token.Decimals(nil)
	if err != nil {
		log.Println("Failed to get decimals from contract: "+contract, err)
		return "error", 0, "error", 0, 0, 0, err
	}
	name, err := token.Name(nil)
	if err != nil {
		log.Println("Failed to retrieve token name from contract: "+contract, err)
		return "error", 0, "error", 0, 0, 0, err
	}

	z := math.Pow(0.1, float64(decimals))
	newBalance := float64(balance.Int64()) * z

	q := math.Pow(0.1, 18)
	newEthBalance := float64(ethAmount.Int64()) * q

	return name, newBalance, symbol, decimals, newEthBalance, maxBlock, err

}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
