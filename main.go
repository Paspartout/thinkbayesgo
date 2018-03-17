package main

import (
	"fmt"

	. "github.com/paspartout/thinkbayesgo/bayes"
)

type cookie struct {
	PMF
	mixes map[string]map[Event]Probability
}

func newCookie(hypos []Event) cookie {
	c := cookie{PMF: make(PMF)}
	for _, hypo := range hypos {
		c.PMF[hypo] = 1
	}
	c.Normalize()

	c.mixes = map[string]map[Event]Probability{
		"Bowl 1": map[Event]Probability{"vanilla": 0.75, "chocolate": 0.25},
		"Bowl 2": map[Event]Probability{"vanilla": 0.5, "chocolate": 0.5},
	}

	return c
}

func (c cookie) likelihood(data Event, hypo string) Probability {
	mix, ok := c.mixes[hypo]
	if !ok {
		panic("")
	}
	like, ok := mix[data]
	if !ok {
		panic("")
	}
	return like
}

func (c cookie) update(data Event) {
	for hypo := range c.PMF {
		like := c.likelihood(data, string(hypo))
		fmt.Println("like", like)
		c.Mult(hypo, like)
	}
	c.Normalize()
}

func main() {
	hypos := []Event{"Bowl 1", "Bowl 2"}
	cookieProblem := newCookie(hypos)

	cookieProblem.update("vanilla")
	cookieProblem.update("chocolate")
	cookieProblem.update("chocolate")

	fmt.Println(cookieProblem)
}
