package genomego

import (
	"testing"
)

const howLong = 7

func TestCrossover(t *testing.T) {
	population := NewPopulation()

	population.Rand = func() float32 { return 0.0 }
	mom := population.NewIndividual(howLong)
	population.Rand = func() float32 { return 1.0 }
	pop := population.NewIndividual(howLong)

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

func TestMutate(t *testing.T) {
	population := NewPopulation()

	population.Rand = func() float32 { return 1.0 }
	mom := population.NewIndividual(1)

	// mom is now all trues
	mom.Mutate(0.5)
	// mom should now be all falses
	if mom.genome[0] {
		t.Error("mutate didn't mutate")
	}
}
