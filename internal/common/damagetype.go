package common

// DamageType refers to SDC vs MDC
type DamageType int

const (
	UnknownDamageType DamageType = iota
	// SDC is Standard Damage Capacity
	SDC
	// MDC is Mega Damage Capacity
	MDC
)

const (
	damageTypeUnknown = "Unknown Damage Type"
	damageTypeSDC     = "SDC"
	damageTypeMDC     = "MDC"
)

var damageTypeNames = map[DamageType]string{
	UnknownDamageType: damageTypeUnknown,
	SDC:               damageTypeSDC,
	MDC:               damageTypeMDC,
}

func (dt DamageType) String() string {
	if s, ok := damageTypeNames[dt]; ok {
		return s
	}
	return damageTypeUnknown
}
