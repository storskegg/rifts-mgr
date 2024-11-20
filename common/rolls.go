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
	for i := 0; i < rs.N; i++ {
		r, err := XRoll(rs.Sides)
		if err != nil {
			return 0, err
		}
		rs.Results = append(rs.Results, r)
	}
	return rs.Sum(), nil
}
func (rs *RollSimple) Sum() int {
	var sum int
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
	_, err := rc.RollSimple.Roll()
	if err != nil {
		return 0, err
	}
	return rc.Sum(), nil
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

// StatRoll represents the special type of roll used for a single stat at character creation
func StatRoll() (*RollSimple, error) {
	core := &RollSimple{
		N:     3,
		Sides: 6,
	}
	got, err := core.Roll()
	if err != nil {
		return nil, err
	}
	if got == 16 || got == 17 || got == 18 {
		exceptional := &RollSimple{
			N:     2,
			Sides: 6,
		}
		_, err = exceptional.Roll()
		if err != nil {
			return nil, err
		}
		core.N++
		core.Results = append(core.Results, exceptional.Results[0])
		if exceptional.Results[0] == 6 {
			core.N++
			core.Results = append(core.Results, exceptional.Results[1])
		}
	}
	return core, nil
}
