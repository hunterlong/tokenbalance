package tokenbalance

import (
	"fmt"
)

var example *TokenBalance

// Create a New Token Balance request with the ERC20 contract address and a wallet address
func ExampleNew() {
	token := "0xb64ef51c888972c908cfacf59b47c1afbc0ab8ac"
	wallet := "0x9ea0c535b3eb166454c8ccbaba86850c8df3ee57"
	example, _ = New(token, wallet)
	fmt.Printf("This wallet has %v %v tokens", example.BalanceString(), example.Name)
	// Output: This wallet has 7.282 StorjToken tokens
}

// View all token information as JSON
func ExampleTokenBalance_ToJSON() {
	data := example.ToJSON()
	fmt.Println(data)
}

// Return the Token Balance as a string rather than a *big.Int
func ExampleTokenBalance_BalanceString() {
	tokens := example.BalanceString()
	fmt.Printf("This wallet has %v %v tokens", tokens, example.Name)
	// Output: This wallet has 7.282 StorjToken tokens
}

// Return the ETH Balance as a string rather than a *big.Int
func ExampleTokenBalance_ETHString() {
	eth := example.ETHString()
	fmt.Printf("This wallet has %v ETH", eth)
	// Output: This wallet has 0.277525175999999985 ETH
}
