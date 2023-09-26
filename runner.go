package gendi

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type RunnerFunc func(*Runner) error

type Runner struct {
	Ptr      int
	Code     string
	Commands map[byte]RunnerFunc
	Data     []byte
	Random   *rand.Rand
	Score    int
	Charset  string
}

func NewRunner(code string, datalen int) *Runner {
	return &Runner{
		Ptr:      0,
		Code:     code,
		Commands: make(map[byte]RunnerFunc),
		Data:     make([]byte, datalen),
		Random:   random(),
		Score:    0,
	}
}

func SpawnNewRunners(count int, datalen int, codelen int, charset string) []*Runner {
	var arr []*Runner
	r := random()
	for i := 0; i < count; i++ {
		arr = append(
			arr,
			NewRunner(randomchars(r, codelen, charset), datalen),
		)
	}
	return arr
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
	r2.Data = make([]byte, datalen)
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

func (r *Runner) CloneMutatedArray(times int, charset string, count int) []*Runner {
	var runners []*Runner
	for i := 0; i < count; i++ {
		runners = append(runners, r.CloneMutated(times))
	}
	return runners
}

func (r *Runner) Mutate(times int) {
	// Create runes from string to be changed
	newchars := []rune(r.Code)

	// Get len's predefined
	charsetrunes := []rune(r.Charset)
	charsetlen := len(r.Charset)
	codelen := len(newchars)

	// Operation itself
	for i := 0; i < times; i++ {
		id := r.Random.Int() % codelen
		chr := charsetrunes[r.Random.Int()%charsetlen]

		// Mutate single character according to random choose
		newchars[id] = chr
	}

	// Then update the code into new chars
	r.Code = string(newchars)
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
	// Get overall length of the code
	LEN := len(r.Code)

	// Reset PTR back to zero if it too far
	if r.Ptr >= LEN || r.Ptr < 0 {
		return fmt.Errorf("Out of bounds PTR: %d", r.Ptr)
	}

	// Get character from the code
	chr := r.Code[r.Ptr]

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
	// randomchars()
	return nil // TODO
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
