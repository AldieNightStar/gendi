package main

import (
	"fmt"
	"math"

	"github.com/AldieNightStar/gendi"
)

func main() {
	guessNumber := 12.0

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
		if math.IsInf(r.Data[0], 1) || math.IsInf(r.Data[0], -1) {
			r.Data[0] = 0
		}
		return nil
	})

	judge := gendi.JudgeFunc(func(r *gendi.Runner) float64 {
		g := r.Data[0]

		diff := math.Abs(guessNumber - g)
		if diff > 1000 {
			diff = 1000
		} else if diff < 0 {
			diff = 0
		}

		result := 1000 - diff

		return result
	})

	r := gendi.NewRunner([]rune("++++++++++++"), 1, judge)
	r.SetCommand('+', plus)
	r.SetCommand('-', minus)
	r.SetCommand('*', multiply)
	r.SetCommand(' ', gendi.DO_NOTHING)

	r.Done()

	trained := r.Train(250, 1, 996)

	fmt.Println("CODE:", string(trained.Code), "NUMBER:", trained.Data[0], "SCORE:", trained.Score)

}
