package types

import "strings"

// QueryResResolve Queries Result Payload for a resolve query
// type QueryResFact struct {
// 	Title       string         `json:"title"`
// 	Bid         sdk.Coins      `json:"bid"`
// 	Creator     sdk.AccAddress `json:"creator"`
// 	Time        int64          `json:"time"`
// 	Place       string         `json:"place"`
// 	Description string         `json:"description"`
// }

// // implement fmt.Stringer
// func (r QueryResResolve) String() string {
// 	return r.Value
// }

// QueryResFactDelegationList is a list of title that address that has been delegated to it
type QueryResFactDelegationList []string

// implement fmt.Stringer
func (r QueryResFactDelegationList) String() string {
	return strings.Join(r[:], "\n")
}

type QueryResFactList []string

// implement fmt.Stringer
func (r QueryResFactList) String() string {
	return strings.Join(r[:], "\n")
}

// type QueryResFactDelegation struct {
// 	Title  string `json:"title"`
// 	Shares int64  `json:"shares"`
// }

// // implement fmt.Stringer
// func (n QueryResNames) String() string {
// 	return strings.Join(n[:], "\n")
// }
