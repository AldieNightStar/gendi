package main

import (
	"fmt"
	"time"

	"github.com/AldieNightStar/gendi"
)

func main() {
	r := gendi.NewRunner(
		[]string{"..", "+1", "-1", "*2", "rv", "^2", "/2"},
		12,
		func(r *gendi.Runner[string]) int {
			n := 0
			for _, c := range r.Code {
				if c == "+1" {
					n += 1
				} else if c == "-1" {
					n -= 1
				} else if c == "*2" {
					n *= 2
				} else if c == "rv" {
					n *= -1
				} else if c == "^2" {
					n *= n
				} else if c == "/2" {
					n /= 2
				}
			}
			fmt.Println(r.Code, "RES:", n)
			time.Sleep(time.Millisecond * 100)
			return n
		},
	)

	trained := r.Train(2, 12, 100000)

	fmt.Println("CODE:", trained.Code, "RESULT:", trained.Score)

}
