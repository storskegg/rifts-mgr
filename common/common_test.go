package common

import (
	"testing"
)

func TestDoRoll(t *testing.T) {
	maxRolls := 100
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
