package matcher

import (
	"fmt"
	"github.com/wesovilabs/goa/advice"
	"github.com/wesovilabs/goa/joinpoint"
)

// Matches struct
type Matches []*Match

// Match struct
type Match struct {
	Function *joinpoint.JoinPoint
	Advices  map[string]*advice.Advice
}

// FindMatches return the list of existing matches
func FindMatches(functions *joinpoint.JoinPoints, definitions *advice.Advices) Matches {
	matches := Matches{}

	for _, f := range functions.List() {
		aspects := make(map[string]*advice.Advice)

		for index, d := range definitions.List() {
			if d.Match(f.Path()) {
				aspects[fmt.Sprintf("aspect%v", index)] = d
			}
		}

		if len(aspects) > 0 {
			matches = append(matches, &Match{
				Function: f,
				Advices:  aspects,
			})
		}
	}

	return matches
}
