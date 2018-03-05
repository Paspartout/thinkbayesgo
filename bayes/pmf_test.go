package bayes_test

import (
	"math"
	"testing"

	. "github.com/paspartout/thinkbayesgo/bayes"
)

func TestNormalize(t *testing.T) {
	// Given a pmf
	pmf := make(PMF)
	for i := 0; i <= 1337; i++ {
		pmf[Event(i)] = Probability(i) * Probability(i)
	}

	// When normalizing
	pmf.Normalize()

	// Then expect the sum of all events probabilities to be (around) 1.0
	sum := Probability(0.0)
	for _, prob := range pmf {
		sum += prob
	}
	maxDelta := 0.001
	delta := math.Abs(1.0 - float64(sum))
	if delta > maxDelta {
		t.Error("Sum of probabilities is not 1.0, it is", sum)
	}
}
