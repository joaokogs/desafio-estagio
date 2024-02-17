package schemas

import "gorm.io/gorm"

type Missions struct {
	gorm.Model
	Name        string
	Description string
	Difficulty  string
	Reward      float32
}
