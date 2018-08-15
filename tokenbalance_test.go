package tokenbalance

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestConnection(t *testing.T) {
	c := &Config{
		GethLocation: "https://eth.coinapp.io",
		Logs:         true,
	}
	err := c.Connect()
	assert.Nil(t, err)
}

func TestFormatDecimal(t *testing.T) {
	number := big.NewInt(0)
	number.SetString("72094368689712", 10)
	tokenCorrected := BigIntString(number, 18)
	assert.Equal(t, "0.000072094368689712", tokenCorrected)
}

func TestFormatSmallDecimal(t *testing.T) {
	number := big.NewInt(0)
	number.SetString("123", 10)
	tokenCorrected := BigIntString(number, 18)
	assert.Equal(t, "0.000000000000000123", tokenCorrected)
}

func TestFormatVerySmallDecimal(t *testing.T) {
	number := big.NewInt(0)
	number.SetString("1142400000000001", 10)
	tokenCorrected := BigIntString(number, 18)
	assert.Equal(t, "0.001142400000000001", tokenCorrected)
}

func TestNewTokenBalance(t *testing.T) {
	tb, err := NewTokenBalance("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07", "0x42d4722b804585cdf6406fa7739e794b0aa8b1ff")
	assert.Nil(t, err)
	assert.Equal(t, "0x42D4722B804585CDf6406fa7739e794b0Aa8b1FF", tb.Wallet.String())
	assert.Equal(t, "0xd26114cd6EE289AccF82350c8d8487fedB8A0C07", tb.Contract.String())
	assert.Equal(t, "600000.0", tb.BalanceString())
	assert.Equal(t, "1.020095885777777767", tb.ETHString())
	assert.Equal(t, int64(18), tb.Decimals)
	assert.Equal(t, "OMG", tb.Symbol)
}
