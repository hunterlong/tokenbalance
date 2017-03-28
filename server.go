package main


import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"fmt"
)


type BalanceResponse struct {
	Name	string	`json:"name"`
	Wallet	string	`json:"wallet"`
	Balance	float64	`json:"balance"`
	Symbol	string	`json:"symbol"`
	Decimals	uint8	`json:"decimals"`
}

func getMembersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	contract := vars["contract"]
	wallet := vars["wallet"]

	fmt.Println("Fetching Wallet ",wallet," at contract ",contract)

	name, balance, token, decimals := GetAccount(contract,wallet)

	new := BalanceResponse{
		Name: name,
		Wallet: wallet,
		Balance: balance,
		Symbol: token,
		Decimals: decimals,
	}

	j, _ := json.Marshal(new)
	w.Write(j)
}


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/balance/{contract}/{wallet}", getMembersHandler).Methods("GET")

	fmt.Println("Server Running: 0.0.0.0:19705")

	http.Handle("/", r)
	http.ListenAndServe("0.0.0.0:19705", nil)
}

