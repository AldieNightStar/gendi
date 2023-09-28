package main

import (
	"fmt"
	"math/rand"

	"github.com/AldieNightStar/gendi"
)

type SimpleStringUnit struct {
	Commands   []string
	CommandSet []string
}

func (self *SimpleStringUnit) Mutate(random *rand.Rand) gendi.Unit {
	clone := &SimpleStringUnit{
		make([]string, len(self.Commands)),
		self.CommandSet,
	}

	// Copy elements
	copy(clone.Commands, self.Commands)

	// Mutation
	id := random.Int() % len(clone.Commands)
	idSet := random.Int() % len(clone.CommandSet)
	clone.Commands[id] = clone.CommandSet[idSet]

	return clone
}

func (self *SimpleStringUnit) Score() int {
	score := 0
	for id, cmd := range self.Commands {
		if cmd == "?" {
			score = 0
		} else if cmd == "~" {
			score *= -1
		}
		if id <= 1 {
			if cmd == "M" {
				score += 100
			}
		} else {
			if cmd == "M" {
				score -= 10
			}
		}

		if id > 5 {
			if cmd == "+" {
				score += 1
			} else if cmd == "-" {
				score -= 1
			} else if cmd == "/" {
				score *= 2
			}
		} else {
			if cmd == "+" {
				score -= 1
			} else if cmd == "-" {
				score += 1
			} else if cmd == "/" {
				score -= 4
			}
		}
	}
	return score
}

func main() {
	s := &SimpleStringUnit{
		[]string{"+", "?", "?", "?", "?", "?", "?", "-", "-", "-", "-", "-", "-", "-", "-", "?", "?", "?", "?", "?", "?", "?"},
		[]string{"+", "-", "?", "~", "/", "M"},
	}

	trained, score := gendi.Train(s, 5, 9100000)

	s = trained.(*SimpleStringUnit)

	fmt.Println("BEST", s.Commands, score)
}
