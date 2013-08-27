package main

import (
	"fmt"
	"github.com/jeffallen/genomego"
)

func calcFitness() float64 {
	return 0
}

func main() {
	fmt.Println("GA!")
	pop := genomego.NewPopulation()
	father := pop.NewIndividual(7)
	fmt.Println("father:", father)

	mother := pop.NewIndividual(7)
	fmt.Println("mother:", mother)

	child1, child2 := father.Crossover(mother, 2)
	fmt.Println("Children:", child1, child2)

	father.Mutate(0.0) //0% mutation
	fmt.Println("father after 0% mutation:", father)

	father.Mutate(1.0) //100% mutation
	fmt.Println("father after 100% mutation:", father)
}
