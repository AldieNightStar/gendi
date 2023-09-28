package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/AldieNightStar/gendi"
)

var _random = rand.New(rand.NewSource(time.Now().UnixNano()))

type SimpleStringUnit struct {
	Commands   []string
	CommandSet []string
}

func (self *SimpleStringUnit) Mutate() gendi.Unit {
	clone := &SimpleStringUnit{
		make([]string, len(self.Commands)),
		self.CommandSet,
	}

	// Copy elements
	copy(clone.Commands, self.Commands)

	// Mutation
	id := _random.Int() % len(clone.Commands)
	idSet := _random.Int() % len(clone.CommandSet)
	clone.Commands[id] = clone.CommandSet[idSet]

	return clone
}

func (self *SimpleStringUnit) Score() int {
	score := 0
	for id, cmd := range self.Commands {
		if id > 5 {
			if cmd == "+" {
				score += 1
			} else if cmd == "-" {
				score -= 1
			}
		} else {
			if cmd == "+" {
				score -= 1
			} else if cmd == "-" {
				score += 1
			}
		}
	}
	return score
}

func main() {
	s := &SimpleStringUnit{
		[]string{"+", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"},
		[]string{"+", "-"},
	}

	trained := gendi.Train(s, 10, 14)

	s = trained.(*SimpleStringUnit)

	fmt.Println(s.Commands)
}
