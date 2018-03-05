package main

import (
	"fmt"
	"strconv"

	. "github.com/paspartout/thinkbayesgo/bayes"
)

func main() {
	// Distributions: Set of values and correspning probabilities
	die := make(PMF)
	events := []int{1, 2, 3, 4, 5, 6}
	for _, ev := range events {
		evStr := Event(strconv.Itoa(ev))
		die[evStr] = Probability(1) / Probability(len(events))
	}

	fmt.Println(die)
}
