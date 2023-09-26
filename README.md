# Gendi

* Genetic algorithm to solve the problem
* Program will find solution and write small BF-like program for their solution
* It could guess the number, but also could solve more complex stuff

## Runner Type
```go
type Runner struct {
	Ptr      int
	Code     []rune
	Commands map[rune]RunnerFunc
	Data     []float64
	Score    float64
	Charset  []rune
	Judge    JudgeFunc
}
```

## Usage
```go
// Create new runner
r := gendi.NewRunner([]rune("+-++++"), 1, judge)

// Set training functions
r.SetCommand('+', plus)
r.SetCommand('-', minus)

// Done with commands
r.Done()


// Get trained Runner
// It will return best trained value
//
// countPerGen   - Number of species per a generation
// mutationLevel - How many mutations should single unit have
// minScore      - Minimum score to pass the training successfully
//
// Return BEST trained runner
trained := r.Train(100, 2, 1000)
```






# Train to guess number
```go
// Let's
guessNumber := 12.0

// Prepare some functions
plus := gendi.RunnerFunc(func(r *gendi.Runner) error {
    r.Data[0] += 1
    return nil
})

minus := gendi.RunnerFunc(func(r *gendi.Runner) error {
    r.Data[0] -= 1
    return nil
})


// Judgement function. This function will tell how much score each runner should take
// It takes data from the runner and counts score
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


// Create new runner with basic code, data length and by giving judgment function
r := gendi.NewRunner([]rune("+++++"), 1, judge)

// Give it a functions to work with
// Last should be do nothing function
r.SetCommand('+', plus)
r.SetCommand('-', minus)
r.SetCommand(' ', gendi.DO_NOTHING)

// Done with setup
r.Done()


// Now train it
// It will return best trained value
//
// countPerGen   - Number of species per a generation
// mutationLevel - How many mutations should single unit have
// minScore      - Minimum score to pass the training successfully
//
// Return BEST trained runner
trained := r.Train(100, 2, 9000)


```