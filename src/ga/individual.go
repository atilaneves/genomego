package individual


import "math/rand"
import "time"


type Individual struct {
	genome []bool
}


func New(size uint32) *Individual {
	ind := Individual{}
	ind.genome = make([]bool, size)
	generator := rand.New(rand.NewSource(time.Now().UnixNano()));
	for i := 0; i < len(ind.genome); i++ {
		ind.genome[i] = generator.Int() % 2 == 0;
	}
	return &ind
}


func (this *Individual) Crossover(other *Individual, pos uint32) [2]*Individual {
	if len(this.genome) == 0 {
		panic("Empty individual!")
	}
	if len(other.genome) != len(this.genome) {
		panic("Other individual has different length!");
	}

	child1 := New(uint32(len(this.genome)))
	child2 := New(uint32(len(this.genome)))
	for i := 0; uint32(i) < pos; i++ {
		child1.genome[i] = this.genome[i]
		child2.genome[i] = other.genome[i]
	}

	for i := pos; i < uint32(len(this.genome)); i++ {
		child2.genome[i] = this.genome[i]
		child1.genome[i] = other.genome[i]
	}

	return [...]*Individual{child1, child2}
}
