package skills

import "github.com/shopspring/decimal"

type SkillValue decimal.Decimal

type BaseValue interface {
	Base() SkillValue
	PerLevel() SkillValue
	Total() SkillValue
}

type baseValue struct {
	base     decimal.Decimal
	perLevel decimal.Decimal
}

func (bs *baseValue) Base() SkillValue {
	return SkillValue(bs.base)
}

func (bs *baseValue) PerLevel() SkillValue {
	return SkillValue(bs.perLevel)
}

func (bs *baseValue) Total(characterLevel int) SkillValue {
	if characterLevel < 2 {
		return bs.Base()
	}
	levelsOverBase := decimal.NewFromInt(int64(characterLevel - 1))
	return SkillValue(bs.base.Add(bs.perLevel.Mul(levelsOverBase)))
}

type Skill interface {
	Category() Category
	Name() string
	Description() string
	BaseValue()
	Requires() []Skill
	Notes() []string
}

type skill struct {
	category    Category
	name        string
	description string
	baseValue   BaseValue
	requires    []Skill
	notes       []string
}

func (s *skill) Category() Category {
	return s.category
}

func (s *skill) Name() string {
	return s.name
}

func (s *skill) Description() string {
	return s.description
}

func (s *skill) BaseValue() BaseValue {
	return s.baseValue
}

func (s *skill) Requires() []Skill {
	return s.requires
}

func (s *skill) Notes() []string {
	return s.notes
}
