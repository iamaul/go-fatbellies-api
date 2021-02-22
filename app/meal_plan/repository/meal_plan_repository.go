package repository

import (
	"github.com/google/uuid"
	mealPlan "github.com/iamaul/fatbellies/app/meal_plan"
	"github.com/iamaul/fatbellies/app/models"
	"github.com/jinzhu/gorm"
)

type mealPlanRepository struct {
	Db *gorm.DB
}

func NewMealPlanRepository(connection *gorm.DB) mealPlan.Repository {
	return &mealPlanRepository{connection}
}

func (mpr *mealPlanRepository) Fetch(limit int64, offset int64, order string) (res *[]models.MealPlan, err error) {
	plan := &[]models.MealPlan{}

	if err = mpr.Db.Model(&models.MealPlan{}).Limit(limit).Offset(limit * (offset - 1)).Order(order).Preload("Branches").Find(&plan).Error; err != nil {
		return
	}

	res = plan

	return
}

func (mpr *mealPlanRepository) GetByID(id uuid.UUID) (res models.MealPlan, err error) {
	plan := models.MealPlan{}

	// ToDo: Redis cache get

	if err = mpr.Db.Model(&models.MealPlan{}).Where("id = ?", id).Preload("Branches").First(&plan).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return
		}
		return
	}

	res = plan

	return
}

func (mpr *mealPlanRepository) GetByName(name string) (res models.MealPlan, err error) {
	plan := models.MealPlan{}

	// ToDo: Redis cache get

	if err = mpr.Db.Model(&models.MealPlan{}).Where("meal_plan_name = ?", name).Preload("Branches").First(&plan).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return
		}
		return
	}

	res = plan

	return
}

func (mpr *mealPlanRepository) Store(plan *models.MealPlan) (res *models.MealPlan, err error) {
	tx := mpr.Db.Begin()

	if err := tx.Create(&plan).Error; err != nil {
		tx.Rollback()
	}

	tx.Commit()

	res = plan

	return
}

func (mpr *mealPlanRepository) Update(id uuid.UUID, newPlan models.MealPlan) (res models.MealPlan, err error) {
	plan := models.MealPlan{}

	db := mpr.Db.Model(plan).Where("id = ?", id).UpdateColumns(newPlan)
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(err) {
			return
		}
		return
	}

	newPlan, err = mpr.GetByID(id)
	if err != nil {
		return
	}

	res = newPlan

	// ToDo: Redis cache update

	return
}

func (mpr *mealPlanRepository) Delete(id uuid.UUID) (err error) {
	if err = mpr.Db.Model(&models.MealPlan{}).Where("id = ?", id).Find(&models.MealPlan{}).Delete(&models.MealPlan{}).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return
		}
		return
	}

	// ToDo: Redis cache delete

	return
}

func (mpr *mealPlanRepository) SearchPlans(column string, query string, order string) (res *[]models.MealPlan, err error) {
	plan := &[]models.MealPlan{}

	if err = mpr.Db.Model(&models.MealPlan{}).Where(column+" ILIKE ?", "%"+query+"%").Order(order).Preload("Branches").Find(&plan).Error; err != nil {
		return
	}

	res = plan

	return
}
