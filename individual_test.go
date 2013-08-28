package genomego

import (
	"testing"
)

const howLong = 7

type ZeroGenerator struct { }
func (*ZeroGenerator) Float64() float64 {
	return 0.0;
}
func (*ZeroGenerator) Intn(int) int {
	return 0;
}

type OneGenerator struct { }
func (*OneGenerator) Float64() float64 {
	return 1.0;
}
func (*OneGenerator) Intn(int) int {
	return 0;
}


func TestCrossover(t *testing.T) {
	zeroGenerator := Generator(&ZeroGenerator{})
	mom := NewIndividual(howLong, zeroGenerator)

	oneGenerator := Generator(&OneGenerator{})
	pop := NewIndividual(howLong, oneGenerator)

	kid1, kid2 := mom.Crossover(pop, howLong/2)

	if mom.Size() != kid1.Size() || kid1.Size() != kid2.Size() {
		t.Error("unexpected sizes")
	}

	for i := range kid1.genome {
		if kid1.genome[i] == kid2.genome[i] {
			t.Error("position ", i, " should not be equal")
		}
	}
}

func TestMutate(t *testing.T) {
	oneGenerator := Generator(&OneGenerator{})
	mom := NewIndividual(1, oneGenerator)

	// mom is now all trues
	mom.Mutate(0.5)
	// mom should now be all falses
	if mom.genome[0] {
		t.Error("mutate didn't mutate")
	}
}


func numTrues(genome []bool) float64 {
	var value uint
	for _, b := range genome {
		if(b) {
			value++
		}
	}
	return float64(value)
}

func TestCalcFitness(t *testing.T) {
	zeroGenerator := Generator(&ZeroGenerator{})
	size := 10
	ind1 := NewIndividual(size, zeroGenerator)
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
