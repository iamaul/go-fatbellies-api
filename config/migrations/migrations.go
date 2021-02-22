package migrations

import (
	"github.com/iamaul/fatbellies/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Migrate(db *gorm.DB) {
	Branch := &models.Branch{}
	BranchLocation := &models.BranchLocation{}
	MealPlan := &models.MealPlan{}
	db.AutoMigrate(&Branch, &BranchLocation, &MealPlan)
}
