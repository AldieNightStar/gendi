package gendi

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var DO_NOTHING = RunnerFunc(func(r *Runner) error { return nil })

type RunnerFunc func(*Runner) error

type Runner struct {
	Ptr      int
	Code     []rune
	Commands map[rune]RunnerFunc
	Data     []int
	Random   *rand.Rand
	Score    int
	Charset  []rune
}

func NewRunner(code []rune, datalen int) *Runner {
	return &Runner{
		Ptr:      0,
		Code:     []rune(code),
		Commands: make(map[rune]RunnerFunc),
		Data:     make([]int, datalen),
		Random:   random(),
		Score:    0,
	}
}

func (r *Runner) Done() {
	// Generate a charset for the future run
	var runes []rune
	for k, _ := range r.Commands {
		runes = append(runes, k)
	}
	r.Charset = runes
}

func (r *Runner) SetCommand(chr rune, c RunnerFunc) {
	r.Commands[chr] = c
}

func (r *Runner) IsBetterThan(r2 *Runner) bool {
	// Always better than nil :)
	if r2 == nil {
		return true
	}

	// Return comparing
	return r.Score > r2.Score
}

func (r *Runner) Clone() *Runner {
	// Get len of data
	datalen := len(r.Data)

	// Create new Runner
	r2 := NewRunner(r.Code, datalen)

	// It will reuse the commands
	r2.Commands = r.Commands

	// Clone the data
	r2.Data = make([]int, datalen)
	for id, d := range r.Data {
		r2.Data[id] = d
	}

	// Random cannot be cloned
	// Rather it could be recreated new one

	// Clone the score
	r2.Score = r.Score

	// Return the clone
	return r2
}

func (r *Runner) CloneMutated(times int) *Runner {
	r2 := r.Clone()
	r2.Mutate(times)
	return r2
}

func (r *Runner) CloneMutatedMany(times int, count int) []*Runner {
	var runners []*Runner
	for i := 0; i < count; i++ {
		runners = append(runners, r.CloneMutated(times))
	}
	return runners
}

func (r *Runner) Mutate(times int) {
	// Get len's predefined
	charsetlen := len(r.Charset)
	codelen := len(r.Code)

	// Operation itself
	for i := 0; i < times; i++ {
		id := r.Random.Int() % codelen
		chr := r.Charset[r.Random.Int()%charsetlen]

		// Mutate single character according to random choose
		r.Code[id] = chr
	}
}

func (r *Runner) StepAll() error {
	// Reset pointer to start
	r.Ptr = 0

	// Operation itself
	for {
		// Execute all
		err := r.Step()
		if err != nil {
			// Take the error message
			message := err.Error()
			// Break only if the error is: Out of bounds
			if strings.HasPrefix(message, "Out of bounds") {
				break
			}

			// If other errors then return it
			return err
		}
	}
	// If no errors
	return nil
}

func (r *Runner) Step() error {
	// Turn code into runes to access the runes
	coderunes := []rune(r.Code)

	// Get overall length of the code
	LEN := len(coderunes)

	// Reset PTR back to zero if it too far
	if r.Ptr >= LEN || r.Ptr < 0 {
		return fmt.Errorf("Out of bounds PTR: %d", r.Ptr)
	}

	// Get character from the code
	chr := coderunes[r.Ptr]

	// Find the command according to a character
	command, commandOk := r.Commands[chr]
	if !commandOk {
		return fmt.Errorf("Character Command '%s' is not defined", string(chr))
	}

	// Run the command
	err := command(r)
	if err != nil {
		return err
	}

	// Increase the pointer
	r.Ptr += 1

	return nil
}

func (r *Runner) Train(countPerGen, generations, mutationLevel int) *Runner {
	runners := r.CloneMutatedMany(mutationLevel, countPerGen)

	// Work with generations
	for gid := 0; gid < generations; gid++ {

		// Eval each runner
		for _, r := range runners {
			r.StepAll()
		}

		// Then take the best one
		best := TakeBestRunner(runners)

		// Now create new generation from best one
		runners = best.CloneMutatedMany(mutationLevel, countPerGen)
	}

	// When generations is done then take best one again and return
	return TakeBestRunner(runners)
}

// Get randomizer from UnixNano
func random() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func TakeBestRunner(runners []*Runner) *Runner {
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

func randomchars(r *rand.Rand, count int, charset string) string {
	// Prepare
	runes := make([]rune, count)
	charsetlen := len(charset)

	// Populate string with random chars
	for i := 0; i < count; i++ {
		runes[i] = rune(charset[r.Int()%charsetlen])
	}

	// Return new string
	return string(runes)
}
