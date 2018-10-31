package tokenbalance

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"os"
	"testing"
)

func TestFailedConnection(t *testing.T) {
	c := &Config{
		GethLocation: "https://google.com",
		Logs:         true,
	}
	err := c.Connect()
	assert.Error(t, err)
}

func TestFailingNoConfig(t *testing.T) {
	_, err := New("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07", "0x42d4722b804585cdf6406fa7739e794b0aa8b1ff")
	assert.Error(t, err)
}

func TestConnection(t *testing.T) {
	c := &Config{
		GethLocation: os.Getenv("ETH"),
		Logs:         true,
	}
	err := c.Connect()
	assert.Nil(t, err)
}

func TestZeroDecimal(t *testing.T) {
	number := big.NewInt(123456789)
	tokenCorrected := bigIntString(number, 0)
	assert.Equal(t, "123456789", tokenCorrected)
}

func TestZeroBalance(t *testing.T) {
	number := big.NewInt(0)
	tokenCorrected := bigIntString(number, 18)
	assert.Equal(t, "0.0", tokenCorrected)
}

func TestFormatDecimal(t *testing.T) {
	number := big.NewInt(0)
	number.SetString("72094368689712", 10)
	tokenCorrected := bigIntString(number, 18)
	assert.Equal(t, "0.000072094368689712", tokenCorrected)
}

func TestFormatSmallDecimal(t *testing.T) {
	number := big.NewInt(0)
	number.SetString("123", 10)
	tokenCorrected := bigIntString(number, 18)
	assert.Equal(t, "0.000000000000000123", tokenCorrected)
}

func TestFormatVerySmallDecimal(t *testing.T) {
	number := big.NewInt(0)
	number.SetString("1142400000000001", 10)
	tokenCorrected := bigIntString(number, 18)
	assert.Equal(t, "0.001142400000000001", tokenCorrected)
}

func TestFailedNewTokenBalance(t *testing.T) {
	_, err := New("0x42D4722B804585CDf6406fa7739e794b0Aa8b1FF", "0x42d4722b804585cdf6406fa7739e794b0aa8b1ff")
	assert.Error(t, err)
}

func TestSymbolFix(t *testing.T) {
	symbol := symbolFix("0x86Fa049857E0209aa7D9e616F7eb3b3B78ECfdb0")
	assert.Equal(t, "EOS", symbol)
}

func TestTokenBalance_ToJSON(t *testing.T) {
	symbol := symbolFix("0x86Fa049857E0209aa7D9e616F7eb3b3B78ECfdb0")
	assert.Equal(t, "EOS", symbol)
}

func TestNewTokenBalance(t *testing.T) {
	c := &Config{
		GethLocation: os.Getenv("ETH"),
		Logs:         true,
	}
	err := c.Connect()
	assert.Nil(t, err)
	tb, err := New("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07", "0x42d4722b804585cdf6406fa7739e794b0aa8b1ff")
	assert.Nil(t, err)
	assert.Equal(t, "0x42D4722B804585CDf6406fa7739e794b0Aa8b1FF", tb.Wallet.String())
	assert.Equal(t, "0xd26114cd6EE289AccF82350c8d8487fedB8A0C07", tb.Contract.String())
	assert.Equal(t, "600000.0", tb.BalanceString())
	assert.Equal(t, "1.020095885777777767", tb.ETHString())
	assert.Equal(t, int64(18), tb.Decimals)
	assert.Equal(t, "OMG", tb.Symbol)
}
