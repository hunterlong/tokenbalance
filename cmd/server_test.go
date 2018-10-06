package main

import (
	"encoding/json"
	tb "github.com/hunterlong/tokenbalance"
	"github.com/mkideal/cli"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	http.Handle("/", router())
}

func TestVersionCLICommand(t *testing.T) {
	app := &cli.Command{}
	app.Register(versionCommand)
	err := app.Run([]string{"version"})
	assert.Nil(t, err)
}

func TestStartCLICommand(t *testing.T) {
	app := &cli.Command{}
	app.Register(startCommand)
	err := app.Run([]string{"start", "--geth", "https://google.com"})
	assert.Error(t, err)
}

func TestHelpCLICommand(t *testing.T) {
	app := &cli.Command{}
	app.Register(help)
	err := app.Run([]string{"help"})
	assert.Nil(t, err)
}

func TestRootCLICommand(t *testing.T) {
	app := &cli.Command{}
	app.Register(rootCommand)
	err := app.Run([]string{"tokenbalance"})
	assert.Nil(t, err)
}

func TestFailedHTTPServer(t *testing.T) {
	err := startServer("google.com", 8080)
	assert.Error(t, err)
}

func TestConfig(t *testing.T) {
	configs = &tb.Config{
		GethLocation: "https://ropsten.coinapp.io",
		Logs:         true,
	}
	err := configs.Connect()
	assert.Nil(t, err)
}

func TestBalanceCheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/balance/0xcad9c6677f51b936408ca3631220c9e45a9af0f6/0xbfd04af48c978cc0d9bc5e06d9593cb4fb7f6f98", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	router().ServeHTTP(rr, req)
	assert.Equal(t, "929955618.999999", rr.Body.String(), "should be balance")
}

func TestFailingBalanceCheck(t *testing.T) {
	req, _ := http.NewRequest("GET", "/balance/0xBDe8f7820b5544a49D34F9dDeaCAbEDC7C0B5adc/0x17a813df7322f8aac5cac75eb62c0d13b8aea29d", nil)
	rr := httptest.NewRecorder()
	router().ServeHTTP(rr, req)
	assert.Equal(t, 404, rr.Result().StatusCode)
}

func TestTokenJson(t *testing.T) {
	req, err := http.NewRequest("GET", "/token/0xcad9c6677f51b936408ca3631220c9e45a9af0f6/0x17a813df7322f8aac5cac75eb62c0d13b8aea29d", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	router().ServeHTTP(rr, req)
	var d tokenBalanceJson
	json.Unmarshal(rr.Body.Bytes(), &d)

	t.Log(d)
	assert.Equal(t, "DreamTeam Token", d.Name, "should be token name")
	assert.Equal(t, "0x17A813dF7322F8AAC5cAc75eB62c0d13B8aea29D", d.Wallet, "should be wallet address")
	assert.Equal(t, int64(6), d.Decimals, "should be decimals")
	assert.Equal(t, "DTT", d.Symbol, "should be symbol")
	assert.Equal(t, "10000.0", d.Balance, "should be token balance")
	assert.Equal(t, "49.997629392334026477", d.ETH, "should be ETH balance")
}

func TestFailingTokenJson(t *testing.T) {
	req, _ := http.NewRequest("GET", "/token/0xBDe8f7820b5544a49D34F9dDeaCAbEDC7C0B5adc/0x17a813df7322f8aac5cac75eb62c0d13b8aea29d", nil)
	rr := httptest.NewRecorder()
	router().ServeHTTP(rr, req)
	assert.Equal(t, 404, rr.Result().StatusCode)
}

func TestMainnetConnection(t *testing.T) {
	configs = &tb.Config{
		GethLocation: "https://eth.coinapp.io",
		Logs:         true,
	}
	err := configs.Connect()
	assert.Nil(t, err)
}

func TestMainnetTokenJson(t *testing.T) {
	req, err := http.NewRequest("GET", "/token/0xd26114cd6EE289AccF82350c8d8487fedB8A0C07/0x42d4722b804585cdf6406fa7739e794b0aa8b1ff", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	router().ServeHTTP(rr, req)
	var d tokenBalanceJson
	json.Unmarshal(rr.Body.Bytes(), &d)
	assert.Equal(t, "OMGToken", d.Name, "should be token name")
	assert.Equal(t, "0x42D4722B804585CDf6406fa7739e794b0Aa8b1FF", d.Wallet, "should be wallet address")
	assert.Equal(t, int64(18), d.Decimals, "should be decimals")
	assert.Equal(t, "OMG", d.Symbol, "should be symbol")
	assert.Equal(t, "600000.0", d.Balance, "should be token balance")
}

func TestMainnetEOSTokenJson(t *testing.T) {
	req, err := http.NewRequest("GET", "/token/0x86fa049857e0209aa7d9e616f7eb3b3b78ecfdb0/0xbfaa1a1ea534d35199e84859975648b59880f639", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	router().ServeHTTP(rr, req)
	var d tokenBalanceJson
	json.Unmarshal(rr.Body.Bytes(), &d)
	assert.Equal(t, "", d.Name, "should be token name")
	assert.Equal(t, "0xbfaA1A1EA534d35199E84859975648B59880f639", d.Wallet, "should be wallet address")
	assert.Equal(t, int64(18), d.Decimals, "should be decimals")
	assert.Equal(t, "EOS", d.Symbol, "should be symbol")
	assert.Equal(t, "8750000.0", d.Balance, "should be token balance")
}

type tokenBalanceJson struct {
	Contract string `json:"token,omitempty"`
	Wallet   string `json:"wallet,omitempty"`
	Name     string `json:"name,omitempty"`
	Symbol   string `json:"symbol,omitempty"`
	Balance  string `json:"balance"`
	ETH      string `json:"eth_balance,omitempty"`
	Decimals int64  `json:"decimals,omitempty"`
	Block    int64  `json:"block,omitempty"`
}
