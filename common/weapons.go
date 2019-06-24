package common

type Ammo struct {
	Name string
	Qty  int
}

type Weapon struct {
	*Equipment
	Damage         Roll
	StrikeAimBurst bool     // TODO: Update the type here
	Special        []string // TODO: Update the type here
}

type WeaponModern struct {
	*Weapon
	Ammo       Ammo
	ParryRange bool // TODO: Update the type here
}

type WeaponAncient struct {
	*Weapon
	PPEISP int
	Parry  bool // TODO: Update the type here
}
