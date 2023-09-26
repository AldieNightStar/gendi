package gendi

import "testing"

func testPrepareRunner() *Runner {
	r := NewRunner([]rune("00000000"), 1, func(r *Runner) float64 {
		return r.Data[0]
	})

	r.SetCommand('0', func(r *Runner) error {
		r.Data[0] += 1
		return nil
	})
	r.SetCommand('1', func(r *Runner) error {
		r.Data[0] -= 1
		return nil
	})

	r.Done()

	return r
}

func TestMutate(t *testing.T) {
	r := testPrepareRunner()
	runners := r.CloneMutatedMany(1, 30)

	prev := ""
	similarities := 0
	for id, runner := range runners {
		strRunnerCode := string(runner.Code)
		if strRunnerCode == prev {
			similarities += 1
		}
		prev = strRunnerCode
		t.Logf("%d -> %s", id, string(runner.Code))
	}

	if similarities > 15 {
		t.Fatalf("Too much similarities: %d", similarities)
	}
}
