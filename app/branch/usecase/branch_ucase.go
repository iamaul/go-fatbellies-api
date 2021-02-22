package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/iamaul/fatbellies/app/branch"
	"github.com/iamaul/fatbellies/app/models"
)

type branchUsecase struct {
	branchRepo     branch.Repository
	contextTimeout time.Duration
}

func NewBranchUsecase(br branch.Repository) branch.Usecase {
	return &branchUsecase{
		branchRepo: br,
	}
}

func (bu *branchUsecase) Fetch(limit int64, offset int64, order string) (*[]models.Branch, error) {
	if limit == 0 {
		limit = 10
	}

	if order == "" {
		order = "created_at desc"
	}

	res, err := bu.branchRepo.Fetch(limit, offset, order)

	return res, err
}

func (bu *branchUsecase) GetByID(id uuid.UUID) (models.Branch, error) {
	res, err := bu.branchRepo.GetByID(id)

	return res, err
}

func (bu *branchUsecase) GetByName(name string) (models.Branch, error) {
	res, err := bu.branchRepo.GetByName(name)

	return res, err
}

func (bu *branchUsecase) FindNearestLocation(lat float64, long float64) (*[]models.BranchLocation, error) {
	res, err := bu.branchRepo.FindNearestLocation(lat, long)

	return res, err
}

func (bu *branchUsecase) Store(branch *models.Branch) (*models.Branch, error) {
	res, err := bu.branchRepo.Store(branch)

	return res, err
}

func (bu *branchUsecase) StoreMealPlan(mealPlan *models.BranchMealPlan) (err error) {
	err = bu.branchRepo.StoreMealPlan(mealPlan)

	return err
}

func (bu *branchUsecase) Update(id uuid.UUID, branch models.Branch) (models.Branch, error) {
	res, err := bu.branchRepo.Update(id, branch)

	return res, err
}

func (bu *branchUsecase) Delete(id uuid.UUID) (err error) {
	err = bu.branchRepo.Delete(id)

	return err
}

func (bu *branchUsecase) SearchBranches(column string, label string, order string) (*[]models.Branch, error) {
	if order == "" {
		order = "created_at desc"
	}

	res, err := bu.branchRepo.SearchBranches(column, label, order)

	return res, err
}
