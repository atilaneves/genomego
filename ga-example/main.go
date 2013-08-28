package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/atilaneves/genomego"
)

func allOnes(genome []bool) float64 {
	ones := 0
	for _, b := range(genome) {
		if b {
			ones++
		}
	}

	return float64(ones)
}

func main() {
	const popSize = 20
	const genSize = 12
	generator := genomego.Generator(rand.New(rand.NewSource(time.Now().UnixNano())))
	ga := genomego.NewGeneticAlgorithm(popSize, genSize, allOnes, generator)
	const endFitness = genSize
	const mutationRate = 0.05
	fittest := ga.Run(endFitness, mutationRate)
	fmt.Println("fittest is ", fittest)
}
