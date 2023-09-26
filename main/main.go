package main

import (
	"fmt"
	"math"

	"github.com/AldieNightStar/gendi"
)

func main() {
	guessNumber := -9999999.0

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

	divide := gendi.RunnerFunc(func(r *gendi.Runner) error {
		r.Data[0] /= 2
		return nil
	})

	power := gendi.RunnerFunc(func(r *gendi.Runner) error {
		r.Data[0] *= r.Data[0]
		return nil
	})

	add5000 := gendi.RunnerFunc(func(r *gendi.Runner) error {
		r.Data[0] += 5000
		return nil
	})

	add10000 := gendi.RunnerFunc(func(r *gendi.Runner) error {
		r.Data[0] += 10000
		return nil
	})

	minusOne := gendi.RunnerFunc(func(r *gendi.Runner) error {
		r.Data[0] *= -1
		return nil
	})

	judge := gendi.JudgeFunc(func(r *gendi.Runner) float64 {
		g := r.Data[0]

		diff := math.Abs(guessNumber - g)
		if diff > 10000 {
			diff = 10000
		} else if diff < 0 {
			diff = 0
		}

		result := 10000 - diff

		// fmt.Println(string(r.Code), r.Data, result)

		return result
	})

	r := gendi.NewRunner([]rune("                                "), 1, judge)
	r.SetCommand('+', plus)
	r.SetCommand('-', minus)
	r.SetCommand('*', multiply)
	r.SetCommand('/', divide)
	r.SetCommand('^', power)
	r.SetCommand(':', add5000)
	r.SetCommand(';', add10000)
	r.SetCommand('[', minusOne)
	r.SetCommand(' ', gendi.DO_NOTHING)

	r.Done()

	trained := r.Train(100, 2, 9000)

	fmt.Println("CODE:", string(trained.Code), "NUMBER:", trained.Data[0], "SCORE:", trained.Score)

}
