package types

import (
	"strings"
)

// QueryResVotedList is a list of title that address that has been voted to it
type QueryResVotedList []VoteOnFact

// implement fmt.Stringer
func (r QueryResVotedList) String() string {
	return (r[0].Title)
}

type QueryResFactList []string

// implement fmt.Stringer
func (r QueryResFactList) String() string {
	return strings.Join(r[:], "\n")
}
