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
	Upvoters    []sdk.AccAddress `json:"upvoters"`
	Downvoters  []sdk.AccAddress `json:"downvoters"`
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

type VotePower struct {
	Voter sdk.AccAddress `json:"voter"`
	Power sdk.Dec        `json:"power"`
}

// implement fmt.Stringer
func (r VotePower) String() string {
	return r.Power.String()
}

// NewDelegation creates a new delegation object
func NewVotePower() VotePower {
	return VotePower{}
}

//Votepower list struct
type VotePowerList struct {
	VotePowers []VotePower `json:"votePower"`
}

// NewVotePowerList creates a new votePowerList object
func NewVotePowerList() VotePowerList {
	return VotePowerList{}
}

//VoteOnFact struct
type VoteOnFact struct {
	Voter sdk.AccAddress `json:"voter"`
	Title string         `json:"title"`
	//it is true when it is a upvote, false when it is a downvote
	Upvote bool `json:"upvote"`
}

// NewVoteOnFact creates a new VoteOnFact object
func NewVoteOnFact() VoteOnFact {
	return VoteOnFact{}
}
