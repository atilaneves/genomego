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

	kids := mom.Crossover(pop, howLong/2)

	for i := range kids[0].genome {
		if kids[0].genome[i] == kids[1].genome[i] {
			t.Error("position ", i, " should not be equal")
		}
	}
}
