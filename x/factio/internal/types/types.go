package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MinPrice is 1 facttoken
var MinPrice = sdk.Coins{sdk.NewInt64Coin("facttoken", 1)}

//Fact struct
type Fact struct {
	Title       string           `json:"title"`
	Time        int64            `json:"time"`
	Place       string           `json:"place"`
	Description string           `json:"description"`
	Creator     sdk.AccAddress   `json:"creator"`
	Price       sdk.Coins        `json:"price"`
	Stakers     []sdk.AccAddress `json:"stakers"`
}

// NewFact returns a new Fact
func NewFact() Fact {
	return Fact{
		Price: MinPrice,
	}
}

//Fact Check Delegation struct
type FactCheckDelegation struct {
	DelegatorAddress sdk.AccAddress `json:"delegator_address"`
	Title            string         `json:"title"`
	Shares           sdk.Dec        `json:"shares"`
}

// NewDelegation creates a new delegation object
func NewDelegation(delegatorAddr sdk.AccAddress, title string, shares sdk.Dec) FactCheckDelegation {
	return FactCheckDelegation{
		DelegatorAddress: delegatorAddr,
		Title:            title,
		Shares:           shares,
	}
}
