package genomego

import (
	"testing"
)

const howLong = 7

func TestCrossover(t *testing.T) {
	mom := NewIndividual(howLong)
	pop := NewIndividual(howLong)

	for i := range mom.genome {
		mom.genome[i] = true
		pop.genome[i] = false
	}

	kid1, kid2 := mom.Crossover(pop, howLong/2)

	if mom.size != kid1.size || kid1.size != kid2.size {
		t.Error("unexpected sizes")
	}

	for i := range kid1.genome {
		if kid1.genome[i] == kid2.genome[i] {
			t.Error("position ", i, " should not be equal")
		}
	}
}
