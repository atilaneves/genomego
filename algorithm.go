package genomego

import "fmt"

type GeneticAlgorithm struct {
	popSize int
	genSize int
	fitnessFunc func([] bool) float64
	generator Generator
}

func NewGeneticAlgorithm(popSize int, genSize int,
	fitnessFunc func([] bool) float64, generator Generator) *GeneticAlgorithm {
	return  &GeneticAlgorithm{popSize, genSize, fitnessFunc, generator}
}

func (ga *GeneticAlgorithm) Run(endFitness float64 , mutationRate float64) *Individual {
	population := make([]*Individual, ga.popSize)
	for i := 0; i < ga.popSize; i++ {
		population[i] = NewIndividual(ga.genSize, ga.generator)
		population[i].CalculateFitness(ga.fitnessFunc)
	}

	generation := 0

	for getHighestFitess(population) < endFitness {
		printGeneration(population, generation)
		const numParticipants = 2
		population = Tournament(population, numParticipants, ga.generator)
		for _, i := range(population) {
			i.CalculateFitness(ga.fitnessFunc)
		}
		generation++
	}

	printGeneration(population, generation)
	return getFittest(population)
}

func printGeneration(pop []*Individual, generation int) {
	fmt.Println("Generation ", generation);
	for _, i := range pop {
		fmt.Println(i.genome);
	}
	fmt.Println();
}

func getHighestFitess(pop []*Individual) float64 {
	var max float64
	for _, i := range(pop) {
		if i.fitness > max {
			max = i.fitness
		}
	}

	return max

}

func getFittest(pop []*Individual) *Individual {
	max := getHighestFitess(pop)
	var fittest *Individual
	for _, i := range(pop) {
		if i.fitness == max {
			fittest = i
		}
	}

	return fittest

}
