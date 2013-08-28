package genomego

import (
	"testing"
)

const howLong = 7

func TestCrossover(t *testing.T) {
	factory := NewIndividualFactory()

	factory.Rand = func() float32 { return 0.0 }
	mom := factory.NewIndividual(howLong)
	factory.Rand = func() float32 { return 1.0 }
	pop := factory.NewIndividual(howLong)

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
	factory := NewIndividualFactory()

	factory.Rand = func() float32 { return 1.0 }
	mom := factory.NewIndividual(1)

	// mom is now all trues
	mom.Mutate(0.5)
	// mom should now be all falses
	if mom.genome[0] {
		t.Error("mutate didn't mutate")
	}
}


func numTrues(genome []bool) float64 {
	var value uint
	for _, b := range(genome) {
		if(b) {
			value++
		}
	}
	return float64(value)
}

func TestCalcFitness(t *testing.T) {
	factory := NewIndividualFactory()
	factory.Rand = func() float32 { return 0.0 }
	size := 10
	ind1 := factory.NewIndividual(size)
	ind1.CalculateFitness(numTrues)
	if int(ind1.fitness) != size {
		t.Error("fitness not equal to size")
	}
	ind1.genome[0] = false
	ind1.CalculateFitness(numTrues)
	if int(ind1.fitness) != size - 1 {
		t.Error("fitness did not decrease")
	}
}
