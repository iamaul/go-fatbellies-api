package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/iamaul/fatbellies/app/branch"
	"github.com/iamaul/fatbellies/app/models"
	"github.com/iamaul/fatbellies/utils"
	"github.com/jinzhu/gorm"
)

type branchRepository struct {
	Db *gorm.DB
}

func NewBranchRepository(connection *gorm.DB) branch.Repository {
	return &branchRepository{connection}
}

func (br *branchRepository) Fetch(limit int64, offset int64, order string) (res *[]models.Branch, err error) {
	branch := &[]models.Branch{}

	if err = br.Db.Model(&models.Branch{}).Limit(limit).Offset(limit * (offset - 1)).Order(order).Preload("MealPlans").Preload("BranchLocations").Find(&branch).Error; err != nil {
		return
	}

	res = branch

	return
}

func (br *branchRepository) GetByID(id uuid.UUID) (res models.Branch, err error) {
	branch := models.Branch{}

	// ToDo: Redis cache get

	if err = br.Db.Model(&models.Branch{}).Where("id = ?", id).Preload("MealPlans").Preload("BranchLocations").First(&branch).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return
		}
		return
	}

	res = branch

	return
}

func (br *branchRepository) GetByName(name string) (res models.Branch, err error) {
	branch := models.Branch{}

	// ToDo: Redis cache get

	if err = br.Db.Model(&models.Branch{}).Where("branch_name = ?", name).Preload("MealPlans").Preload("BranchLocations").First(&branch).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return
		}
		return
	}

	res = branch

	return
}

func (br *branchRepository) FindNearestLocation(lat float64, long float64) (res *[]models.BranchLocation, err error) {
	branchLocation := &[]models.BranchLocation{}

	if err = br.Db.Model(&models.BranchLocation{}).Raw("SELECT branch_locations.id, branch_locations.branch_id, branch_locations.latitude, branch_locations.longitude, (3959 * acos(cos(radians(?)) * cos(radians(branch_locations.latitude)) * cos(radians(branch_locations.longitude) - radians(?)) + sin(radians(?)) * sin(radians(branch_locations.latitude)))) AS distance FROM branch_locations",
		lat, long, lat).Order("distance").Find(&branchLocation).Error; err != nil {
		return
	}

	res = branchLocation

	return
}

func (br *branchRepository) Store(branch *models.Branch) (res *models.Branch, err error) {
	if err = br.Db.Model(&models.Branch{}).Where("branch_name = ?", branch.BranchName).Find(&models.Branch{}).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			tx := br.Db.Begin()

			if err := tx.Create(&branch).Error; err != nil {
				tx.Rollback()
			}

			tx.Commit()

			return branch, nil
		}
	}

	err = errors.New(utils.BranchExists)

	return
}

func (br *branchRepository) StoreMealPlan(mealPlan *models.BranchMealPlan) (err error) {
	tx := br.Db.Begin()

	if err := tx.Create(&mealPlan).Error; err != nil {
		tx.Rollback()
	}

	tx.Commit()

	return
}

func (br *branchRepository) Update(id uuid.UUID, newBranch models.Branch) (res models.Branch, err error) {
	branch := models.Branch{}

	db := br.Db.Model(branch).Where("id = ?", id).UpdateColumns(newBranch)
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(err) {
			return
		}
		return
	}

	newBranch, err = br.GetByID(id)
	if err != nil {
		return
	}

	res = newBranch

	// ToDo: Redis cache update

	return
}

func (br *branchRepository) Delete(id uuid.UUID) (err error) {
	if err = br.Db.Model(&models.Branch{}).Where("id = ?", id).Find(&models.Branch{}).Delete(&models.Branch{}).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return
		}
		return
	}

	// ToDo: Redis cache delete

	return
}

func (br *branchRepository) SearchBranches(column string, query string, order string) (res *[]models.Branch, err error) {
	branch := &[]models.Branch{}

	if err = br.Db.Model(&models.Branch{}).Where(column+" ILIKE ?", "%"+query+"%").Order(order).Preload("MealPlans").Preload("BranchLocations").Find(&branch).Error; err != nil {
		return
	}

	res = branch

	return
}
