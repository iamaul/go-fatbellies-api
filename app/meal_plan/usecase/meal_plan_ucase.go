package usecase

import (
	"time"

	"github.com/google/uuid"
	mealPlan "github.com/iamaul/fatbellies/app/meal_plan"
	"github.com/iamaul/fatbellies/app/models"
)

type mealPlanUsecase struct {
	mealPlanRepo   mealPlan.Repository
	contextTimeout time.Duration
}

func NewMealPlanUsecase(pr mealPlan.Repository) mealPlan.Usecase {
	return &mealPlanUsecase{
		mealPlanRepo: pr,
	}
}

func (mpu *mealPlanUsecase) Fetch(limit int64, offset int64, order string) (*[]models.MealPlan, error) {
	if limit == 0 {
		limit = 10
	}

	if order == "" {
		order = "created_at desc"
	}

	res, err := mpu.mealPlanRepo.Fetch(limit, offset, order)

	return res, err
}

func (mpu *mealPlanUsecase) GetByID(id uuid.UUID) (models.MealPlan, error) {
	res, err := mpu.mealPlanRepo.GetByID(id)

	return res, err
}

func (mpu *mealPlanUsecase) GetByName(name string) (models.MealPlan, error) {
	res, err := mpu.mealPlanRepo.GetByName(name)

	return res, err
}

func (mpu *mealPlanUsecase) Store(branch *models.MealPlan) (*models.MealPlan, error) {
	res, err := mpu.mealPlanRepo.Store(branch)

	return res, err
}

func (mpu *mealPlanUsecase) Update(id uuid.UUID, branch models.MealPlan) (models.MealPlan, error) {
	res, err := mpu.mealPlanRepo.Update(id, branch)

	return res, err
}

func (mpu *mealPlanUsecase) Delete(id uuid.UUID) (err error) {
	err = mpu.mealPlanRepo.Delete(id)

	return err
}

func (mpu *mealPlanUsecase) SearchPlans(column string, label string, order string) (*[]models.MealPlan, error) {
	if column == "price" {
		column = "price::text"
	}
	if order == "" {
		order = "created_at desc"
	}

	res, err := mpu.mealPlanRepo.SearchPlans(column, label, order)

	return res, err
}
