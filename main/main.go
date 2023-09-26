package main

import (
	"fmt"

	"github.com/AldieNightStar/gendi"
)

func main() {
	// guessNumber := 32

	plus := gendi.RunnerFunc(func(r *gendi.Runner) error {
		r.Data[0] += 1
		return nil
	})

	minus := gendi.RunnerFunc(func(r *gendi.Runner) error {
		r.Data[0] -= 1
		return nil
	})

	multiply := gendi.RunnerFunc(func(r *gendi.Runner) error {
		r.Data[0] *= 2
		return nil
	})

	r := gendi.NewRunner([]rune("+++++++++++++"), 1)
	r.SetCommand('+', plus)
	r.SetCommand('-', minus)
	r.SetCommand('*', multiply)
	r.SetCommand(' ', gendi.DO_NOTHING)

	r.Done()
	r.StepAll()

	fmt.Println(r.Data[0])
}
