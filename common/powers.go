package common

import (
	"github.com/storskegg/rifts-mgr/internal/common/rolls"
	"time"
)

// PowerPointType represents the points used to fuel an ability
type PowerPointType int

const (
	// ISP are Inner Strength Points, and fuel psionic powers
	ISP PowerPointType = 0
	// PPE are Potential Psychic Energy, and fuel magic powers
	PPE PowerPointType = 1
)

// Power represents the most generic common denominator between spells, psionics and effects
type Power struct {
	Name           string
	PowerPointType PowerPointType
	Cost           int
	SavingThrow    bool // TODO: Update this type
	Duration       *time.Duration
	Damage         rolls.Roll
}

type Spell struct {
	*Power
}

type Psionic struct {
	*Power
}

type Effect struct {
	*Power
}
