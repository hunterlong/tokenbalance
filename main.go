package tokenbalance

import (
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
	// Create an IPC based RPC connection to a remote node

	//
	// MAC: /Users/username/Library/Ethereum/geth.ipc
	//
	// LINUX: /root/.ethereum/geth.ipc
	//

	conn, err = ethclient.Dial(GethLocation)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// CJX
	//reqAddress := "0x38122af92aa2819Cdb2bE029f2Fa325959B52Ec8";
	//reqContract := "0x2ce349291b8365f8d12c4cedf992969f680c726e";

	// DAO
	//reqAddress := "0x21314ff1669aced72b3c72ad912102186cf5e1cd";
	//reqContract := "0x48c80F1f4D53D5951e5D5438B54Cba84f29F32a5";

}

func GetAccount(contract string, wallet string) (string, float64, string, uint8, float64) {
	var err error

	token, err := NewTokenCaller(common.HexToAddress(contract), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	address := common.HexToAddress(wallet)
	balance, err := token.BalanceOf(nil, address)
	symbol, err := token.Symbol(nil)
	decimals, err := token.Decimals(nil)
	name, err := token.Name(nil)

	ethAmount := 0

	if err != nil {
		log.Fatalf("Failed to retrieve token name: %v", err)
	}

	z := math.Pow(0.1, float64(decimals))
	newBalance := float64(balance.Int64()) * z

	q := math.Pow(0.1, 10)
	newEthBalance := float64(ethAmount) * q

	return name, newBalance, symbol, decimals, newEthBalance

}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
