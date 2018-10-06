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

func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/balance/{contract}/{wallet}", getBalanceHandler).Methods("GET")
	r.HandleFunc("/token/{contract}/{wallet}", getTokenHandler).Methods("GET")
	return r
}

func startServer(ip string, port int) error {
	log.Printf("TokenBalance Server Running: http://%v:%v", ip, port)
	srv := &http.Server{
		Addr:         fmt.Sprintf("%v:%v", ip, port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router(),
	}
	return srv.ListenAndServe()
}

func collectVars(r *http.Request) (string, string) {
	vars := mux.Vars(r)
	return vars["contract"], vars["wallet"]
}

func getTokenHandler(w http.ResponseWriter, r *http.Request) {
	contract, wallet := collectVars(r)
	log.Println("Fetching /token for Wallet:", wallet, "at Contract:", contract)
	query, err := tb.New(contract, wallet)
	if err != nil {
		m := errorResponse{
			Error:   true,
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(m)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(query.ToJSON()))
	}
}

func getBalanceHandler(w http.ResponseWriter, r *http.Request) {
	contract, wallet := collectVars(r)
	log.Println("Fetching /balance for Wallet:", wallet, "at Contract:", contract)
	query, err := tb.New(contract, wallet)
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
