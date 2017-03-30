package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/mkideal/cli"
)

type BalanceResponse struct {
	Name       string  `json:"name"`
	Wallet     string  `json:"wallet"`
	Balance    float64 `json:"balance"`
	EthBalance float64 `json:"eth_balance"`
	Symbol     string  `json:"symbol"`
	Decimals   uint8   `json:"decimals"`
}

func getMembersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	contract := vars["contract"]
	wallet := vars["wallet"]

	fmt.Println("Fetching Wallet ", wallet, " at contract ", contract)

	name, balance, token, decimals, ethAmount := GetAccount(contract, wallet)

	new := BalanceResponse{
		Name:       name,
		Wallet:     wallet,
		Balance:    balance,
		EthBalance: ethAmount,
		Symbol:     token,
		Decimals:   decimals,
	}

	j, _ := json.Marshal(new)
	w.Write(j)
}



type argT struct {
	cli.Helper
	Geth  string    `cli:"g,geth" usage:"geth IPC location" dft:"~/.ethereum/geth.ipc"`
	IP    string `cli:"ip" usage:"Bind to IP Address" dft:"0.0.0.0"`
	Port    string `cli:"p,port" usage:"HTTP port for JSON" dft:"19705"`
}


func main() {

	cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		GethLocation = argv.Geth
		UsePort = argv.Port
		UseIP = argv.IP
		ConnectGeth()
		StartServer()
		return nil
	})

}


func StartServer() {


	r := mux.NewRouter()
	r.HandleFunc("/balance/{contract}/{wallet}", getMembersHandler).Methods("GET")

	fmt.Println("TokenBalance Server Running: "+UseIP+":"+UsePort)

	http.Handle("/", r)
	http.ListenAndServe(UseIP+":"+UsePort, nil)

}