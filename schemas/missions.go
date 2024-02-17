package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Missions struct {
	gorm.Model
	Name        string
	Description string
	Difficulty  string
	Reward      float32
}

type MissionsOpening struct {
	ID          uint      `json: "id"`
	CreatedAt   time.Time `json: "createdAt"`
	UpdatedAt   time.Time `json: "updatedAt"`
	DeletedAt   time.Time `json: "deletedAt, omitempty"`
	Name        string    `json: "name"`
	Description string    `json: "description"`
	Difficulty  string    `json: "difficulty"`
	Reward      float32   `json: "reward"`
}
