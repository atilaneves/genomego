package genomego

func Tournament(population []*Individual, numParticipants int, generator Generator) []*Individual {
	var newPopulation []*Individual;
	for len(newPopulation) < len(population) {
		father := getWinner(population, numParticipants, generator)
		mother := getWinner(population, numParticipants, generator)

		xoverPoint := generator.Intn(father.Size())
		child1, child2 := father.Crossover(mother, xoverPoint)

		child1.Mutate(generator.Float64())
		child2.Mutate(generator.Float64())

		newPopulation = append(newPopulation, child1)
		newPopulation = append(newPopulation, child2)
	}

	return newPopulation
}

func getWinner(population []*Individual, numParticipants int, generator Generator) (winner *Individual) {
	participantIndices := make([]int, numParticipants)
	for i := 0; i < numParticipants; i++ {
		participantIndices[i] = generator.Intn(numParticipants)
	}
	var maxFitness float64 = 0
	for _, individual := range population {
		if individual.fitness > maxFitness {
			maxFitness = individual.fitness
			winner = individual
		}
	}
	return
}
