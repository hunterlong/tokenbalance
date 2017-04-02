package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
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

	log.Println("Fetching Wallet:", wallet, "at Contract:", contract)

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

func StartServer() {

	r := mux.NewRouter()
	r.HandleFunc("/", getMembersHandler).Methods("GET")
	r.HandleFunc("/balance/{contract}/{wallet}", getMembersHandler).Methods("GET")

	log.Println("TokenBalance Server Running: http://" + UseIP + ":" + UsePort)

	http.Handle("/", r)
	http.ListenAndServe(UseIP+":"+UsePort, nil)

}
