package main

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/ethclient"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	http.Handle("/", Router())
}

func TestConnection(t *testing.T) {
	var err error
	conn, err = ethclient.Dial("https://main.cjx.io")
	if err != nil {
		t.Fail()
	}
}

func TestBalanceCheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/balance/0xa74476443119A942dE498590Fe1f2454d7D4aC0d/0xda0aed568d9a2dbdcbafc1576fedc633d28eee9a", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	if rr.Body.String() != "5401731.0867782926098" {
		t.Fail()
	}
}

func TestTokenJson(t *testing.T) {
	req, err := http.NewRequest("GET", "/token/0xa74476443119A942dE498590Fe1f2454d7D4aC0d/0xda0aed568d9a2dbdcbafc1576fedc633d28eee9a", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	var d BalanceResponse

	json.Unmarshal(rr.Body.Bytes(), &d)

	if d.Name != "Golem Network Token" {
		t.Fail()
	}

	if d.Wallet != "0xda0aed568d9a2dbdcbafc1576fedc633d28eee9a" {
		t.Fail()
	}

	if d.Symbol != "GNT" {
		t.Fail()
	}
}
