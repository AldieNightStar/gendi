package gendi

import (
	"math/rand"
	"time"
)

type JudgeFunc[T any] func(*Runner[T]) int

var _random = rand.New(rand.NewSource(time.Now().UnixNano()))

type Runner[T any] struct {
	Code  []T
	Score int
	Set   []T
	Judge JudgeFunc[T]
}

func NewRunner[T any](set []T, codelen int, judger JudgeFunc[T]) *Runner[T] {
	code := make([]T, codelen)
	for id, _ := range code {
		code[id] = set[0]
	}
	return &Runner[T]{
		Code:  code,
		Score: 0,
		Set:   set,
		Judge: judger,
	}
}

func (r *Runner[T]) Clone() *Runner[T] {

	// Create new Runner
	r2 := NewRunner[T](r.Set, len(r.Code), r.Judge)

	// Clone the code
	r2.Code = make([]T, len(r.Code))
	for id, c := range r.Code {
		r2.Code[id] = c
	}

	// Return the clone
	return r2
}

func (r *Runner[T]) CloneMutated(times int) *Runner[T] {
	r2 := r.Clone()
	r2.Mutate(times)
	return r2
}

func (r *Runner[T]) CloneMutatedMany(times int, count int) []*Runner[T] {
	var runners []*Runner[T]
	for i := 0; i < count; i++ {
		clone := r.CloneMutated(times)
		runners = append(runners, clone)
	}
	return runners
}

func (r *Runner[T]) Mutate(times int) {
	// Get len's predefined
	codelen := len(r.Code)

	// Operation itself
	for i := 0; i < times; i++ {
		id := _random.Int() % codelen
		cmd := r.Set[_random.Int()%len(r.Set)]

		// Mutate single character according to random choose
		r.Code[id] = cmd
	}
}

func (r *Runner[T]) Train(countPerGen, mutationLevel int, minimumScore int) *Runner[T] {
	runners := r.CloneMutatedMany(mutationLevel, countPerGen)

	lastBest := r.Clone()
	lastBest.Score = 0

	// Work with generations
	for {

		// Eval each runner and judge the score for it
		for _, runner := range runners {

			// Judge how many it scored based on data itself
			runner.Score = runner.Judge(runner)

		}

		// Then take the best one by score
		best := TakeBestRunner(runners)

		// if generation score is less than previous, then this is REGRESSION
		// That means it WILL NOT SURVIVE
		//
		// In this case we will reuse OLD one Best value
		if best.Score < lastBest.Score {
			// Use older Best Runner
			runners = lastBest.CloneMutatedMany(mutationLevel*2, countPerGen)
			continue
		}

		// If we found the best Runner
		if best.Score >= minimumScore {
			return best
		}

		// Saving best score for the next generations
		lastBest = best

		// Now create new generation from best one
		runners = best.CloneMutatedMany(mutationLevel, countPerGen)
	}
}

func TakeBestRunner[T any](runners []*Runner[T]) *Runner[T] {
	// Not able to take at least something
	if len(runners) < 1 {
		return nil
	}

	// Take best scored runner
	best := 0
	bestId := 0
	for id, runner := range runners {
		if runner.Score > best {
			best = runner.Score
			bestId = id
		}
	}

	// Return best runner
	// If no best found it will be first one (bestId is 0 initialy)
	return runners[bestId]
}
