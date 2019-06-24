package common

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

// ArmorType refers to whether a piece of armor is body armor or power armor
type ArmorType int

const (
	BodyArmor  ArmorType = 0
	PowerArmor ArmorType = 1
)

type Ammo struct {
	Name string
	Qty  int
}

// Armor represents armor that is worn
type Armor struct {
	*Equipment
	ArmorRating int
	ArmorType   ArmorType
}

type Weapon struct {
	*Equipment
	Damage         Roll
	StrikeAimBurst bool     // TODO: Update the type here
	Special        []string // TODO: Update the type here
}

type WeaponAncient struct {
	*Weapon
	PPEISP int
	Parry  bool // TODO: Update the type here
}

type WeaponModern struct {
	*Weapon
	Ammo       Ammo
	ParryRange bool // TODO: Update the type here
}
