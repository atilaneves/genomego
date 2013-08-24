package main

import "fmt"
import "ga/individual"

func calcFitness() float64 {
	return 0;
}

func main() {
	fmt.Println("GA!")
	father := individual.New(5)
	mother := Individual.New(5)
	children := father.crossover(mother)
	fmt.Println("Children: %v", children)
}
