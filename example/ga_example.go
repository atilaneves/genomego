package main

import "fmt"
import "ga"

func calcFitness() float64 {
	return 0;
}

func main() {
	fmt.Println("GA!")
	father := individual.New(7)
	mother := individual.New(7)
	children := father.Crossover(mother, 2)
	fmt.Print("Children: ")
	for i := 0; i < 2; i++ {
		fmt.Print(*children[i], " ");
	}
	fmt.Println();
}
