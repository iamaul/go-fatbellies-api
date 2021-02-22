package models

import "github.com/google/uuid"

type BranchMealPlan struct {
	BranchID   uuid.UUID `json:"branch_id" validate:"required"`
	MealPlanID uuid.UUID `json:"meal_plan_id" validate:"required"`
}
