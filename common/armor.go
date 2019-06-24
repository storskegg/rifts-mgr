package common

// ArmorType refers to whether a piece of armor is body armor or power armor
type ArmorType int

const (
	BodyArmor  ArmorType = 0
	PowerArmor ArmorType = 1
)

// Armor represents armor that is worn
type Armor struct {
	*Equipment
	ArmorRating int
	ArmorType   ArmorType
}
