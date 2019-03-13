package nameservice

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// whois is a struct that contains all the metadata of a name
type Whois struct {
	Value string			`json:"value"`
	Owner sdk.AccAddress	`json:"owner"`
	Price sdk.Coins			`json:"price"`
}

var MinNamePrice = sdk.Coins{sdk.NewInt64Coin("nametoken", 1)}

// returns a new whois with the minprice as the price
func NewWhois() Whois {
	return Whois {
		Price: MinNamePrice
	}
}
