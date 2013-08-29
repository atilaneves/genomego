package genomego

import "fmt"
import "sync"

type GeneticAlgorithm struct {
	popSize     int
	genSize     int
	fitnessFunc func([]bool) float64
	generator   Generator
}

func NewGeneticAlgorithm(popSize int, genSize int,
	fitnessFunc func([]bool) float64, generator Generator) *GeneticAlgorithm {
	return &GeneticAlgorithm{popSize, genSize, fitnessFunc, generator}
}

func (ga *GeneticAlgorithm) Run(endFitness float64, mutationRate float64) *Individual {
	var population []*Individual
	for i := 0; i < ga.popSize; i++ {
		individual := NewIndividual(ga.genSize, ga.generator)
		individual.CalculateFitness(ga.fitnessFunc)
		population = append(population, individual)
	}

	generation := 0

	for getHighestFitess(population) < endFitness {
		printGeneration(population, generation)
		const numParticipants = 2
		population = Tournament(population, numParticipants, ga.generator)
		var waitGroup sync.WaitGroup
		for _, i := range population {
			waitGroup.Add(1)
			go func(i *Individual) {
				i.CalculateFitness(ga.fitnessFunc)
				waitGroup.Done()
			}(i)
		}
		waitGroup.Wait()
		generation++
	}

	printGeneration(population, generation)
	return getFittest(population)
}

func printGeneration(pop []*Individual, generation int) {
	fmt.Println("Generation ", generation)
	for _, i := range pop {
		fmt.Println(i.genome)
	}
	fmt.Println()
}

func getHighestFitess(pop []*Individual) (max float64) {
	for _, i := range pop {
		if i.fitness > max {
			max = i.fitness
		}
	}

	return
}

func getFittest(pop []*Individual) (fittest *Individual) {
	max := getHighestFitess(pop)
	for _, i := range pop {
		if i.fitness == max {
			fittest = i
		}
	}

	return
}
