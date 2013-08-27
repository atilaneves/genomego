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
	father := genomego.NewIndividual(7)
	fmt.Println("father: ", *father)

	mother := genomego.NewIndividual(7)
	fmt.Println("mother: ", *mother)

	children := father.Crossover(mother, 2)
	fmt.Print("Children: ")
	for i := 0; i < 2; i++ {
		fmt.Print(*children[i], " ")
	}
	fmt.Println()

	father.Mutate(0.0) //0% mutation
	fmt.Println("father after 0% mutation: ", father)

	father.Mutate(1.0) //100% mutation
	fmt.Println("father after 100% mutation: ", father)
}
