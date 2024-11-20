package skills

type Category int

func (c Category) String() string {
	if s, ok := categoryNames[c]; ok {
		return s
	}
	return categoryNameUnknownCategory
}

const (
	UnknownCategory Category = iota
	CommunicationSkills
	CowboySkills
	DomesticSkills
	ElectricalSkills
	EspionageSkills
	HorsemanshipSkills
	MechanicalSkills
	MedicalSkills
	MilitarySkills
	PhysicalSkills
	PilotSkills
	PilotRelatedSkills
	RogueSkills
	ScienceSkills
	TechnicalSkills
	WeaponProficienciesAncient
	WeaponProficienciesModern
	WildernessSkills
)

const (
	categoryNameUnknownCategory            = "Unknown Category"
	categoryNameCommunicationSkills        = "Communication Skills"
	categoryNameCowboySkills               = "Cowboy Skills"
	categoryNameDomesticSkills             = "Domestic Skills"
	categoryNameElectricalSkills           = "Electrical Skills"
	categoryNameEspionageSkills            = "Espionage Skills"
	categoryNameHorsemanshipSkills         = "Horsemanship Skills"
	categoryNameMechanicalSkills           = "Mechanical Skills"
	categoryNameMedicalSkills              = "Medical Skills"
	categoryNameMilitarySkills             = "Military Skills"
	categoryNamePhysicalSkills             = "Physical Skills"
	categoryNamePilotSkills                = "Pilot Skills"
	categoryNamePilotRelatedSkills         = "Pilot-Related Skills"
	categoryNameRogueSkills                = "Rogue Skills"
	categoryNameScienceSkills              = "Science Skills"
	categoryNameTechnicalSkills            = "Technical Skills"
	categoryNameWeaponProficienciesAncient = "Weapon Proficiencies (Ancient)"
	categoryNameWeaponProficienciesModern  = "Weapon Proficiencies (Modern)"
	categoryNameWildernessSkills           = "Wilderness Skills"
)

var categoryNames = map[Category]string{
	UnknownCategory:            categoryNameUnknownCategory,
	CommunicationSkills:        categoryNameCommunicationSkills,
	CowboySkills:               categoryNameCowboySkills,
	DomesticSkills:             categoryNameDomesticSkills,
	ElectricalSkills:           categoryNameElectricalSkills,
	EspionageSkills:            categoryNameEspionageSkills,
	HorsemanshipSkills:         categoryNameHorsemanshipSkills,
	MechanicalSkills:           categoryNameMechanicalSkills,
	MedicalSkills:              categoryNameMedicalSkills,
	MilitarySkills:             categoryNameMilitarySkills,
	PhysicalSkills:             categoryNamePhysicalSkills,
	PilotSkills:                categoryNamePilotSkills,
	PilotRelatedSkills:         categoryNamePilotRelatedSkills,
	RogueSkills:                categoryNameRogueSkills,
	ScienceSkills:              categoryNameScienceSkills,
	TechnicalSkills:            categoryNameTechnicalSkills,
	WeaponProficienciesAncient: categoryNameWeaponProficienciesAncient,
	WeaponProficienciesModern:  categoryNameWeaponProficienciesModern,
	WildernessSkills:           categoryNameWildernessSkills,
}
