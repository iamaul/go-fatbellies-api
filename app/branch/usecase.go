package branch

import (
	"github.com/google/uuid"
	"github.com/iamaul/fatbellies/app/models"
)

// Usecase represent the Branch's usecases
type Usecase interface {
	Fetch(limit int64, offset int64, order string) (*[]models.Branch, error)
	GetByID(id uuid.UUID) (models.Branch, error)
	GetByName(name string) (models.Branch, error)
	FindNearestLocation(lat float64, long float64) (*[]models.BranchLocation, error)
	Store(*models.Branch) (*models.Branch, error)
	StoreMealPlan(*models.BranchMealPlan) error
	Update(id uuid.UUID, branch models.Branch) (models.Branch, error)
	Delete(id uuid.UUID) error
	SearchBranches(column string, label string, order string) (*[]models.Branch, error)
}
