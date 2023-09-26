package main

import (
	"fmt"
	"math"

	"github.com/AldieNightStar/gendi"
)

func main() {
	guessNumber := 32.00

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

	judge := gendi.JudgeFunc(func(r *gendi.Runner) float64 {
		// Difference between guessed number and data
		diff := math.Abs(guessNumber - r.Data[0])

		// The less difference the more score it takes
		return 100.00 - diff
	})

	r := gendi.NewRunner([]rune("++++++++++++"), 1, judge)
	r.SetCommand('+', plus)
	r.SetCommand('-', minus)
	r.SetCommand('*', multiply)
	r.SetCommand(' ', gendi.DO_NOTHING)

	r.Done()

	trained := r.Train(10, 5, 1)
	trained.StepAll()

	fmt.Println(string(trained.Code), trained.Data[0], trained.Score)
}
