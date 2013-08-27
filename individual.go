package genomego

import (
	"fmt"
	"math/rand"
	"time"
)

type Population struct {
	Rand func() float32
}

func NewPopulation() *Population {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &Population{Rand: func() float32 { return generator.Float32() }}
}

type Individual struct {
	size   int
	genome []bool
	p      *Population
}

func (p *Population) NewIndividual(size int) *Individual {
	ind := Individual{size: size, p: p}
	ind.genome = make([]bool, ind.size)
	for i := range ind.genome {
		ind.genome[i] = (p.Rand() < 0.5)
	}
	return &ind
}

// Clone creates a new Individual with the same genome.
func (me *Individual) Clone() *Individual {
	them := me.cloneEmpty()
	copy(them.genome, me.genome)
	return them
}

// cloneEmpty clones an Individual, but leaves the genome initialized
// to the zero values.
func (me *Individual) cloneEmpty() *Individual {
	them := &Individual{
		size:   me.size,
		p:      me.p,
		genome: make([]bool, me.size),
	}
	return them
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

	for i := pos; i < me.size; i++ {
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
func (me *Individual) Mutate(rate float32) {
	for i := 0; i < me.size; i++ {
		if me.p.Rand() < rate {
			me.genome[i] = !me.genome[i]
		}
	}
}
