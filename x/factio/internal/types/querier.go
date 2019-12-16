package types

import "strings"

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
