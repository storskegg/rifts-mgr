package common

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
)

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
