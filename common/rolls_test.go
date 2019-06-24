package common

import (
	"testing"
)

const maxRolls = 100

func TestXRoll(t *testing.T) {
	sides := 10

	for i := 0; i < maxRolls; i++ {
		got, err := XRoll(sides)
		if err != nil {
			t.Error(err)
		} else if got < 1 {
			t.Error("should never get a result less than 1")
		} else if uint64(got) > uint64(sides) {
			t.Error("should never get a result larger than the number of sides")
		}
		t.Logf("%d\n", got)
	}
}

// TODO: Rewrite this to be table-driven instead of single-case happy-path
func TestRollSimple_Roll(t *testing.T) {
	for i := 0; i < maxRolls; i++ {
		simple := &RollSimple{
			N:     1,
			Sides: 4,
		}
		if len(simple.Results) != 0 {
			t.Errorf("expecting %d Results to equal 0 before Roll()", len(simple.Results))
		}
		n, err := simple.Roll()
		if n != simple.N {
			t.Errorf("expected %d to equal N %d", n, simple.N)
		}
		if err != nil {
			t.Error(err)
		}
		if len(simple.Results) != simple.N {
			t.Errorf("expecting %d Results to equal %d after Roll()", len(simple.Results), simple.N)
		}
	}
}

// TODO: Rewrite this to be table-driven instead of single-case happy-path
func TestRollSimple_Sum(t *testing.T) {
	for i := 0; i < maxRolls; i++ {
		var want int
		simple := &RollSimple{
			N:     3,
			Sides: 20,
		}
		n, err := simple.Roll()
		if n != simple.N {
			t.Errorf("expected %d to equal N %d", n, simple.N)
		}
		if err != nil {
			t.Error(err)
		}
		for n := range simple.Results {
			want += n
		}
		got := simple.Sum()
		if got != want {
			t.Errorf("expected %d to equal %d", got, want)
		}
	}
}

// TODO: Rewrite this to be table-driven instead of single-case happy-path
func TestRollComplex_Roll(t *testing.T) {
	for i := 0; i < maxRolls; i++ {
		myRoll := &RollComplex{
			RollSimple: &RollSimple{
				N:     2,
				Sides: 4,
			},
			Multiplier: 1000,
		}
		if len(myRoll.RollSimple.Results) != 0 {
			t.Errorf("expected %d Results to equal 0 before Roll()", len(myRoll.RollSimple.Results))
		}
		n, err := myRoll.Roll()
		if err != nil {
			t.Error(err)
		}
		if n != myRoll.RollSimple.N {
			t.Errorf("expected %d to equal N %d", n, myRoll.RollSimple.N)
		}
	}
}

// TODO: Rewrite this to be table-driven instead of single-case happy-path
func TestRollComplex_Sum(t *testing.T) {
	for i := 0; i < maxRolls; i++ {
		var want int
		myRoll := &RollComplex{
			RollSimple: &RollSimple{
				N:     5,
				Sides: 8,
			},
			Multiplier: 100,
		}
		n, err := myRoll.Roll()
		if err != nil {
			t.Error(err)
		}
		if n != myRoll.RollSimple.N {
			t.Errorf("expected %d to equal N %d", n, myRoll.RollSimple.N)
		}
		for r := range myRoll.RollSimple.Results {
			want += r
		}
		want = want * myRoll.Multiplier
		got := myRoll.Sum()
		if got != want {
			t.Errorf("exptected %d to equal %d", got, want)
		}
	}
}
