package gendi

import (
	"math/rand"
	"time"
)

var _random = rand.New(rand.NewSource(time.Now().UnixNano()))

// Unit data for train/work
type Unit interface {
	// Should return new clone with single change
	// Mutation is when single piece of data is changed.
	//
	// Accepts Random object
	Mutate(*rand.Rand) Unit

	// Returns score points according to data in the Unit
	// The more data we have, the more chance to survive
	Score() int
}

// Train models on testing Unit for taking best code algorithm.
// Your Unit data should have that code inside.
// After a training you could unwrap that interface into your type
// and take the data it results
func Train(lastBest Unit, countPerGen, minimumScore int) (Unit, int) {
	units := mutateMany(lastBest, countPerGen)
	lastBestScore := lastBest.Score()

	// Work with generations
	for {

		unitScorePairs := []*unitScorePair{}

		// Eval each runner and count the score for it
		for _, unit := range units {
			unitScorePairs = append(unitScorePairs, &unitScorePair{unit, unit.Score()})
		}

		// Then take the best one by score
		best := takeBestUnit(unitScorePairs)

		// if generation score is less than previous, then this is REGRESSION
		// That means it WILL NOT SURVIVE
		//
		// In this case we will reuse OLD one Best value
		if best == nil || best.Score < lastBestScore {
			// Use older Best Unit
			// This time we use doubleMutate (Mutate twice)
			units = doubleMutateMany(lastBest, countPerGen)
			continue
		}

		// If we found the best Runner
		if best.Score >= minimumScore {
			return best.Unit, best.Score
		}

		// Saving best score for the next generations
		lastBest = best.Unit
		lastBestScore = best.Score

		// Now create new generation from the best one
		units = mutateMany(lastBest, countPerGen)
	}
}

func mutateMany(original Unit, count int) []Unit {
	var units []Unit
	for i := 0; i < count; i++ {
		units = append(units, original.Mutate(_random))
	}
	return units
}

func doubleMutateMany(original Unit, count int) []Unit {
	var units []Unit
	for i := 0; i < count; i++ {
		units = append(units, original.Mutate(_random).Mutate(_random))
	}
	return units
}

func takeBestUnit(pairs []*unitScorePair) *unitScorePair {
	// Not able to take at least something
	if len(pairs) < 1 {
		return nil
	}

	// Take best scored Unit
	best := 0
	bestId := 0
	for id, pair := range pairs {
		if pair.Score > best {
			best = pair.Score
			bestId = id
		}
	}

	// Return best unit
	// If not it will 0 (Initial value)
	return pairs[bestId]
}
