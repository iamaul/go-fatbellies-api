package models

import (
	"time"

	"github.com/google/uuid"
)

type Branch struct {
	ID              uuid.UUID      `gorm:"primary_key; type:uuid; default:uuid_generate_v4()" json:"id"`
	BranchName      string         `gorm:"type:varchar(125); null;" json:"branch_name" validate:"required,min=3"`
	BranchLocations BranchLocation `json:"locations"`
	OpeningHours    uint8          `gorm:"type:integer; default:0" json:"opening_hours" validate:"required,numeric"`
	MealPlans       []MealPlan     `gorm:"many2many:branch_meal_plans;" json:"branch_meal_plans"`
	CreatedAt       time.Time      `gorm:"type:timestamp without time zone; default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"type:timestamp without time zone; default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt       time.Time      `gorm:"type:timestamp without time zone; null; default:null" json:"deleted_at"`
}

type SwagBranchLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type SwagBranch struct {
	BranchName      string             `json:"branch_name"`
	BranchLocations SwagBranchLocation `json:"locations"`
	OpeningHours    uint8              `json:"opening_hours"`
}
