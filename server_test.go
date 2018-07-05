package main

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	http.Handle("/", Router())
}

func TestConnection(t *testing.T) {
	var err error
	url := "https://ropsten.infura.io/dZfBqHazp5fzdZVJ4Byc"
	conn, err = ethclient.Dial(url)
	assert.Nil(t, err)
}

func TestFormatDecimal(t *testing.T) {
	number := big.NewInt(0)
	number.SetString("72094368689712", 10)
	tokenCorrected := BigIntDecimal(number, 18)
	assert.Equal(t, "0.000072094368689712", tokenCorrected)
}

func TestFormatSmallDecimal(t *testing.T) {
	number := big.NewInt(0)
	number.SetString("123", 10)
	tokenCorrected := BigIntDecimal(number, 18)
	assert.Equal(t, "0.000000000000000123", tokenCorrected)
}

func TestFormatVerySmallDecimal(t *testing.T) {
	number := big.NewInt(0)
	number.SetString("1142400000000001", 10)
	tokenCorrected := BigIntDecimal(number, 18)
	assert.Equal(t, "0.001142400000000001", tokenCorrected)
}

func TestBalanceCheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/balance/0xcad9c6677f51b936408ca3631220c9e45a9af0f6/0x17a813df7322f8aac5cac75eb62c0d13b8aea29d", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, "10000.0", rr.Body.String(), "should be balance")
}

func TestFailingBalanceCheck(t *testing.T) {
	req, _ := http.NewRequest("GET", "/balance/0xBDe8f7820b5544a49D34F9dDeaCAbEDC7C0B5adc/0x17a813df7322f8aac5cac75eb62c0d13b8aea29d", nil)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 404, rr.Result().StatusCode)
}

func TestTokenJson(t *testing.T) {
	req, err := http.NewRequest("GET", "/token/0xcad9c6677f51b936408ca3631220c9e45a9af0f6/0x17a813df7322f8aac5cac75eb62c0d13b8aea29d", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	var d jsonResponse
	json.Unmarshal(rr.Body.Bytes(), &d)

	assert.Equal(t, "DreamTeam Token", d.Name, "should be token name")
	assert.Equal(t, "0x17A813dF7322F8AAC5cAc75eB62c0d13B8aea29D", d.Wallet, "should be wallet address")
	assert.Equal(t, int64(6), d.Decimals, "should be decimals")
	assert.Equal(t, "DTT", d.Symbol, "should be symbol")
	assert.Equal(t, "10000.0", d.Balance, "should be Token balance")
	assert.Equal(t, "49.999936999999995635", d.EthBalance, "should be ETH balance")
}

func TestFailingTokenJson(t *testing.T) {
	req, _ := http.NewRequest("GET", "/token/0xBDe8f7820b5544a49D34F9dDeaCAbEDC7C0B5adc/0x17a813df7322f8aac5cac75eb62c0d13b8aea29d", nil)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 404, rr.Result().StatusCode)
}

func TestMainnetConnection(t *testing.T) {
	var err error
	url := "https://mainnet.infura.io/dZfBqHazp5fzdZVJ4Byc"
	conn, err = ethclient.Dial(url)
	assert.Nil(t, err)
}

func TestMainnetTokenJson(t *testing.T) {
	req, err := http.NewRequest("GET", "/token/0xd26114cd6EE289AccF82350c8d8487fedB8A0C07/0x42d4722b804585cdf6406fa7739e794b0aa8b1ff", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	var d jsonResponse
	json.Unmarshal(rr.Body.Bytes(), &d)
	assert.Equal(t, "OMGToken", d.Name, "should be token name")
	assert.Equal(t, "0x42D4722B804585CDf6406fa7739e794b0Aa8b1FF", d.Wallet, "should be wallet address")
	assert.Equal(t, int64(18), d.Decimals, "should be decimals")
	assert.Equal(t, "OMG", d.Symbol, "should be symbol")
	assert.Equal(t, "600000.0", d.Balance, "should be Token balance")
}

func TestMainnetEOSTokenJson(t *testing.T) {
	req, err := http.NewRequest("GET", "/token/0x86fa049857e0209aa7d9e616f7eb3b3b78ecfdb0/0xbfaa1a1ea534d35199e84859975648b59880f639", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	var d jsonResponse
	json.Unmarshal(rr.Body.Bytes(), &d)
	assert.Equal(t, "", d.Name, "should be token name")
	assert.Equal(t, "0xbfaA1A1EA534d35199E84859975648B59880f639", d.Wallet, "should be wallet address")
	assert.Equal(t, int64(18), d.Decimals, "should be decimals")
	assert.Equal(t, "EOS", d.Symbol, "should be symbol")
	assert.Equal(t, "8750000.0", d.Balance, "should be Token balance")
}
