package custom

import (
	u "peter/utils"
)

func LoadChars() []u.Entity {
	AgentZero := u.Entity{
		Name:    "AgentZero",
		Attack:  1,
		Defense: 1,
		Health:  100,
		Agility: 1,
	}
	var CustomChars []u.Entity = []u.Entity{AgentZero}
	return CustomChars

}
