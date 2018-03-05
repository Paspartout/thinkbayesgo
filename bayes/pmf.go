// bayes contains types and methods for Bayesian statistics that
// follow the example of the Think Bayes[1] book.
//
// [1] - http://greenteapress.com/wp/think-bayes/
package bayes

import (
	"fmt"
	"math"
)

const (
	// maxDelta determins how much floating point values are allowed to
	// deviate from a specific value, see for example IsNormalized()
	maxDelta float64 = 0.001
)

// Event is the representation of an outcome of an experiment
// Mathematically an event is a subset all possible outcomes of an experiment.
type Event string

// Probability is usually a value between 0 and 1 that measures the liklehood
// that a specific Event occurs
type Probability float64

// PMF stands for Probability Mass Function and maps an event to a certain
// probability.
type PMF map[Event]Probability

// Incr increases the "probability" or frequency for event by 1.
func (pmf PMF) Incr(event Event, num Probability) {
	pmf[event] += num
}

// Sum sums up all probabilties
func (pmf PMF) Sum() Probability {
	var sum Probability
	for _, freq := range pmf {
		sum += freq
	}

	return sum
}

// Normalize assigns each event a probability, so all of them add up to 1
func (pmf PMF) Normalize() {
	sum := pmf.Sum()
	// Divide each event probability/frequency by sum to normalize
	for ev, freq := range pmf {
		pmf[ev] = freq / sum
	}
}

// IsNormalized returns if the probabilites roughly add up to 1.0
func (pmf PMF) IsNormalized() bool {
	delta := math.Abs(1.0 - float64(pmf.Sum()))
	return delta <= maxDelta
}

// String returns a better string representation for the map
func (pmf PMF) String() string {
	str := "PMF ["
	for ev, prob := range pmf {
		str += fmt.Sprintf("%s: %.3f\n", ev, prob)
	}
	str += "]"
	return str
}
