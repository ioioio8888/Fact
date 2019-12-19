package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

// MinPrice is 1 factcoin
var MinPrice = sdk.Coins{sdk.NewInt64Coin("factcoin", 1)}

//Fact struct
type Fact struct {
	Title       string           `json:"title"`
	Time        int64            `json:"time"`
	Place       string           `json:"place"`
	Description string           `json:"description"`
	Creator     sdk.AccAddress   `json:"creator"`
	Price       sdk.Coins        `json:"price"`
	Delegators  []sdk.AccAddress `json:"delegators"`
	Tags        []string         `json:"tags"`
}

// implement fmt.Stringer
func (w Fact) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Title: %s
Time: %d
Place: %s
Description: %s
Creator: %s
Price: %s
Delegators: %s
`, w.Title, w.Time, w.Place, w.Description, w.Creator, w.Price, w.Delegators))
}

// NewFact returns a new Fact
func NewFact() Fact {
	return Fact{
		Price: MinPrice,
	}
}

//Fact Delegation struct
type FactDelegation struct {
	Delegator sdk.AccAddress `json:"delegator_address"`
	Title     string         `json:"title"`
	Shares    sdk.Coins      `json:"shares"`
}

// NewDelegation creates a new delegation object
func NewDelegation() FactDelegation {
	return FactDelegation{}
}

//Fact Delegation struct
type FactDelegationList struct {
	Delegations []FactDelegation `json:"delegations"`
}

// NewDelegation creates a new delegation object
func NewDelegationList() FactDelegationList {
	return FactDelegationList{}
}
