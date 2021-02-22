package models

import (
	"time"

	"github.com/google/uuid"
)

type MealPlan struct {
	ID           uuid.UUID `gorm:"primary_key; type:uuid; default:uuid_generate_v4()" json:"id"`
	MealPlanName string    `gorm:"type:varchar(150); null;" json:"meal_plan_name" validate:"required,min=3"`
	MaxCapacity  uint8     `gorm:"type:integer; default:10" json:"max_capacity" validate:"required,numeric"`
	Price        uint64    `gorm:"type:integer; default:5" json:"price" validate:"required,numeric"`
	Day          string    `gorm:"type:varchar(40); null;" json:"day" validate:"required"`
	StartTime    time.Time `gorm:"type:timestamp without time zone; null;" json:"start_time"`
	EndTime      time.Time `gorm:"type:timestamp without time zone; null;" json:"end_time"`
	Branches     []Branch  `gorm:"many2many:branch_meal_plans;" json:"branch_meal_plans"`
	CreatedAt    time.Time `gorm:"type:timestamp without time zone; default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"type:timestamp without time zone; default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    time.Time `gorm:"type:timestamp without time zone; null; default:null" json:"deleted_at"`
}

type SwagMealPlan struct {
	MealPlanName    string `json:"meal_plan_name"`
	BranchLocations uint8  `json:"max_capacity"`
	Price           uint64 `json:"price"`
	Day             string `json:"day"`
}
