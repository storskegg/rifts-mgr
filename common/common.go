package common

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
)

// Roll is an interface generally referring to any form of roll, simple or complex
type Roll interface {
	Roll() (int, error)
	Sum() int
}

// RollSimple represents a simple roll (e.g. 1d8 or 3d6) with no multipliers
type RollSimple struct {
	N       int
	Sides   int
	Results []int
}

func (rs *RollSimple) Roll() (int, error) {
	var n int
	for i := 0; i < rs.N; i++ {
		n++
		r, err := XRoll(rs.Sides)
		if err != nil {
			return n, err
		}
		rs.Results = append(rs.Results, r)
	}
	return n, nil
}
func (rs *RollSimple) Sum() (sum int) {
	for r := range rs.Results {
		sum += r
	}
	return sum
}

// RollComplex represents a simple roll with a multiplier (e.g. 2d4 x 1000)
type RollComplex struct {
	RollSimple *RollSimple
	Multiplier int
}

func (rc *RollComplex) Roll() (int, error) {
	n, err := rc.RollSimple.Roll()
	return n, err
}
func (rc *RollComplex) Sum() int {
	return rc.RollSimple.Sum() * rc.Multiplier
}

// XRoll is a method that accepts a number of sides, and returns a value between 1 and `sides` inclusively.
// Because so many people "don't trust" the machine rolls, I've opted for cryptographically random numbers.
// This is overkill, but you know...gotta satisfy the masses.
func XRoll(sides int) (int, error) {
	if sides < 1 {
		return 0, fmt.Errorf("dice roll must have 1 or more sides")
	}

	buf := make([]byte, 2)
	var result int

	for result < 1 || result > sides {
		_, err := rand.Read(buf)
		if err != nil {
			return 0, err
		}

		result = int(binary.BigEndian.Uint16(buf))
	}

	return result, nil
}

// DamageType is refers to SDC vs MDC
type DamageType int

const (
	// SDC is Standard Damage Capacity
	SDC DamageType = 0
	// MDC is Mega Damage Capacity
	MDC DamageType = 1
)

// Artifact represents any _thing_ of value, be it armor, a weapon, or a crocheted doily made by the character's
// grandmother. This does not include faction credits
type Artifact struct {
	Name           string
	Notes          []string
	Cost           int
	Weight         float64
	DamageCapacity int
	DamageType     DamageType
}

// Equipment represents an Artifact that may be equipped or worn (including power armor), but does not include things
// that can be piloted like mounts or vehicles.
type Equipment struct {
	*Artifact
	Equipped bool
}
