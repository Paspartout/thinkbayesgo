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

func TestCookieProblem(t *testing.T) {
	// Given a prior distribution of 0.5 for Bowl 1 and 0.5 for Bowl 2
	pmf := make(PMF)
	pmf["Bowl 1"] = 0.5
	pmf["Bowl 2"] = 0.5

	// When updating with liklehood 0.75 for Bowl 1 and 0.5 for Bowl 2
	pmf.Mult("Bowl 1", 0.75)
	pmf.Mult("Bowl 2", 0.5)
	// and normalizing afterwards
	pmf.Normalize()

	// Then the anwser should be 0.6 for Bowl 1 and 0.4 for Bowl 2
	if posterior := pmf["Bowl 1"]; posterior != 0.6 {
		t.Error("Posterior distribution of Bowl 1 is wrong: ", posterior)
	}
	if posterior := pmf["Bowl 2"]; posterior != 0.4 {
		t.Error("Posterior distribution of Bowl 2 is wrong: ", posterior)
	}
}
