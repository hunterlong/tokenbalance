package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	tb "github.com/hunterlong/tokenbalance"
	"log"
	"net/http"
	"time"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/balance/{contract}/{wallet}", getBalanceHandler).Methods("GET")
	r.HandleFunc("/token/{contract}/{wallet}", getTokenHandler).Methods("GET")
	return r
}

func StartServer() error {
	log.Printf("TokenBalance Server Running: http://%v:%v", configs.UseIP, configs.UsePort)
	routes := Router()
	srv := &http.Server{
		Addr:         fmt.Sprintf("%v:%v", configs.UseIP, configs.UsePort),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      routes,
	}
	return srv.ListenAndServe()
}

func getTokenHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contract := vars["contract"]
	wallet := vars["wallet"]

	log.Println("Fetching /token for Wallet:", wallet, "at Contract:", contract)

	query, err := tb.NewTokenBalance(contract, wallet)
	if err != nil {
		m := errorResponse{
			Error:   true,
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(m)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(query.ToJSON()))
	}

}

func getBalanceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contract := vars["contract"]
	wallet := vars["wallet"]

	log.Println("Fetching /balance for Wallet:", wallet, "at Contract:", contract)

	query, err := tb.NewTokenBalance(contract, wallet)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("0.0"))
	} else {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(query.BalanceString()))
	}
}
