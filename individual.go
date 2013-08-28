package genomego

import (
	"fmt"
)

type Generator interface {
	Float64() float64
	Intn(int) int
}

type Individual struct {
	genome    []bool
	generator Generator
	fitness   float64
}

func NewIndividual(size int, generator Generator) *Individual {
	ind := Individual{genome: make([]bool, size), generator: generator}
	ind.genome = make([]bool, size)
	for i := range ind.genome {
		ind.genome[i] = (generator.Float64() < 0.5)
	}
	return &ind
}

// Clone creates a new Individual with the same genome.
func (me *Individual) Clone() *Individual {
	them := me.cloneEmpty()
	copy(them.genome, me.genome)
	return them
}

func (me *Individual) Size() int {
	return len(me.genome)
}

// cloneEmpty clones an Individual, but leaves the genome initialized
// to the zero values.
func (me *Individual) cloneEmpty() *Individual {
	return &Individual{make([]bool, len(me.genome)), me.generator, me.fitness}
}

func (me *Individual) Crossover(other *Individual, pos int) (child1, child2 *Individual) {
	if len(other.genome) != len(me.genome) {
		panic("Other individual has different length!")
	}

	child1 = me.cloneEmpty()
	child2 = me.cloneEmpty()
	for i := 0; i < pos; i++ {
		child1.genome[i] = me.genome[i]
		child2.genome[i] = other.genome[i]
	}

	for i := pos; i < len(me.genome); i++ {
		child2.genome[i] = me.genome[i]
		child1.genome[i] = other.genome[i]
	}

	return
}

// String makes Individual implement interface Stringer (see fmt)
// so that when printed Println or printf %v, we only show useful stuff.
func (me *Individual) String() string {
	return fmt.Sprint(me.genome)
}

// Mutate flips the alelles in the Individual's genome with
// a probability of rate (for 0.0 <= rate < 1.0)
func (me *Individual) Mutate(rate float64) {
	for i := 0; i < len(me.genome); i++ {
		if me.generator.Float64() < rate {
			me.genome[i] = !me.genome[i]
		}
	}
}

func (me *Individual) CalculateFitness(fitnessFunc func([]bool) float64) {
	me.fitness = fitnessFunc(me.genome)
}
