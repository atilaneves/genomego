package genomego

import (
	"testing"
)

type altIntGenerator struct {
	floatValue float64
	intValue   int
}

func (d *altIntGenerator) Float64() float64 {
	return d.floatValue
}

func (d *altIntGenerator) Intn(int) int {
	d.intValue++
	return (d.intValue - 1) % 2 //sigh
}

func TestTournament(t *testing.T) {
	const numIndividuals = 2
	pop := make([]*Individual, numIndividuals)

	generator := altIntGenerator{0.0, 0}
	for i := 0; i < numIndividuals; i++ {
		pop[i] = NewIndividual(howLong, Generator(&generator))
		pop[i].CalculateFitness(numTrues)
	}

	const numParticipants = 2
	newPop := Tournament(pop, numParticipants, Generator(&generator))
	if len(newPop) != numIndividuals {
		t.Error("new population has different number of individuals")
	}
	for _, i := range newPop {
		for _, b := range i.genome {
			if !b {
				t.Error("All bits should be true")
			}
		}
	}
}
