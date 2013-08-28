package genomego

func Tournament(population []*Individual, numParticipants int, generator Generator) []*Individual {
	index := 0
	newPopulation := make([]*Individual, len(population))
	for index < len(population) {
		father := getWinner(population, numParticipants, generator)
		mother := getWinner(population, numParticipants, generator)

		xoverPoint := generator.Intn(father.size)
		child1, child2 := father.Crossover(mother, xoverPoint)

		child1.Mutate(generator.Float32())
		child2.Mutate(generator.Float32())

		newPopulation[index] = child1
		index++ // sigh
		newPopulation[index] = child2
		index++ // sigh
	}

	return newPopulation
}

func getWinner(population []*Individual, numParticipants int, generator Generator) *Individual {
	participantIndices := make([]int, numParticipants)
	for i := 0; i < numParticipants; i++ {
		participantIndices[i] = generator.Intn(numParticipants)
	}
	var maxFitness float64 = 0
	var winner *Individual
	for _, individual := range(population) {
		if individual.fitness > maxFitness {
			maxFitness = individual.fitness
			winner = individual
		}
	}
	return winner
}
