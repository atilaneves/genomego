package genomego

import (
	"math/rand"
	"time"
)

var generator = rand.New(rand.NewSource(time.Now().UnixNano()))

type Individual struct {
	size int
	genome []bool
}

func NewIndividual(size int) *Individual {
	ind := Individual{ size: size }
	ind.genome = make([]bool, ind.size)
	for i := range ind.genome {
		ind.genome[i] = generator.Int31n(2) == 1
	}
	return &ind
}

func (me *Individual) Crossover(other *Individual, pos int) [2]*Individual {
	if len(other.genome) != len(me.genome) {
		panic("Other individual has different length!")
	}

	child1 := NewIndividual(me.size)
	child2 := NewIndividual(me.size)
	for i := 0; i < pos; i++ {
		child1.genome[i] = me.genome[i]
		child2.genome[i] = other.genome[i]
	}

	for i := pos; i < me.size; i++ {
		child2.genome[i] = me.genome[i]
		child1.genome[i] = other.genome[i]
	}

	return [2]*Individual{child1, child2}
}

func (me *Individual) Mutate(rate float64) {
	for i := 0; i < me.size; i++ {
		if generator.Float64() < rate {
			me.genome[i] = !me.genome[i]
		}
	}
}
