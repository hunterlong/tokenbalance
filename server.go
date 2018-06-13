package main

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gorilla/mux"
	"log"
	"math/big"
	"net/http"
	"time"
)

type BalanceResponse struct {
	Name       string         `json:"name,omitempty"`
	Wallet     common.Address `json:"wallet,omitempty"`
	Symbol     string         `json:"symbol,omitempty"`
	Balance    *big.Int       `json:"balance"`
	EthBalance *big.Int       `json:"eth_balance,omitempty"`
	Decimals   *big.Int       `json:"decimals,omitempty"`
	Block      *types.Block   `json:"block,omitempty"`
}

type ErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func getTokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	contract := vars["contract"]
	wallet := vars["wallet"]

	log.Println("Fetching /token for Wallet:", wallet, "at Contract:", contract)

	response, err := GetAccount(contract, wallet)

	if err != nil {
		m := ErrorResponse{
			Error:   true,
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusNotFound)
		msg, _ := json.Marshal(m)
		w.Write(msg)
	} else {
		w.WriteHeader(http.StatusOK)
		jsoned, _ := json.Marshal(response.Format())
		w.Write(jsoned)
	}
}

func getBalanceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	vars := mux.Vars(r)
	contract := vars["contract"]
	wallet := vars["wallet"]

	log.Println("Fetching /balance for Wallet:", wallet, "at Contract:", contract)

	response, err := GetAccount(contract, wallet)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("0.0"))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response.Format().Balance))
	}
}

func StartServer() {
	log.Printf("TokenBalance Server Running: http://%v:%v", UseIP, UsePort)
	routes := Router()
	srv := &http.Server{
		Addr:         fmt.Sprintf("%v:%v", UseIP, UsePort),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      routes,
	}
	srv.ListenAndServe()
}

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/balance/{contract}/{wallet}", getBalanceHandler).Methods("GET")
	r.HandleFunc("/token/{contract}/{wallet}", getTokenHandler).Methods("GET")
	return r
}
