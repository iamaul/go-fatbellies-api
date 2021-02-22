package meal_plan

import (
	"github.com/google/uuid"
	"github.com/iamaul/fatbellies/app/models"
)

// Usecase represent the Meal Plan's usecases
type Usecase interface {
	Fetch(limit int64, offset int64, order string) (*[]models.MealPlan, error)
	GetByID(id uuid.UUID) (models.MealPlan, error)
	GetByName(name string) (models.MealPlan, error)
	Store(*models.MealPlan) (*models.MealPlan, error)
	Update(id uuid.UUID, plan models.MealPlan) (models.MealPlan, error)
	Delete(id uuid.UUID) error
	SearchPlans(column string, label string, order string) (*[]models.MealPlan, error)
}
