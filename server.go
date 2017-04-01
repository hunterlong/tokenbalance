package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mkideal/cli"
	"log"
	"net/http"
)

type BalanceResponse struct {
	Name       string  `json:"name"`
	Wallet     string  `json:"wallet"`
	Symbol     string  `json:"symbol"`
	Balance    float64 `json:"balance"`
	EthBalance float64 `json:"eth_balance"`
	Decimals   uint8   `json:"decimals"`
	Block      uint64  `json:"block"`
}

type ErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func getMembersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	contract := vars["contract"]
	wallet := vars["wallet"]

	log.Println("Fetching Wallet: ", wallet, "at contract:", contract)

	name, balance, token, decimals, ethAmount, block, err := GetAccount(contract, wallet)

	if err != nil {
		m := ErrorResponse{
			Error:   true,
			Message: "could not find contract address",
		}
		msg, _ := json.Marshal(m)
		w.Write(msg)
		return
	}

	new := BalanceResponse{
		Name:       name,
		Symbol:     token,
		Wallet:     wallet,
		Balance:    balance,
		EthBalance: ethAmount,
		Decimals:   decimals,
		Block:      block,
	}

	j, err := json.Marshal(new)

	if err == nil {
		w.Write(j)
	}
}

type argT struct {
	cli.Helper
	Geth string `cli:"g,geth" usage:"geth IPC location" dft:"~/.ethereum/geth.ipc"`
	IP   string `cli:"ip" usage:"Bind to IP Address" dft:"0.0.0.0"`
	Port string `cli:"p,port" usage:"HTTP port for JSON" dft:"19705"`
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

	log.Println("TokenBalance Server Running: " + UseIP + ":" + UsePort)

	log.Println("Try it out! Go to: http://" + UseIP + ":" + UsePort + "/balance/0xa74476443119A942dE498590Fe1f2454d7D4aC0d/0xda0aed568d9a2dbdcbafc1576fedc633d28eee9a")

	http.Handle("/", r)
	http.ListenAndServe(UseIP+":"+UsePort, nil)

}
