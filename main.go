package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/iamaul/fatbellies/docs"
	echoSwagger "github.com/swaggo/echo-swagger"

	bh "github.com/iamaul/fatbellies/app/branch/delivery/http"
	br "github.com/iamaul/fatbellies/app/branch/repository"
	bu "github.com/iamaul/fatbellies/app/branch/usecase"

	mph "github.com/iamaul/fatbellies/app/meal_plan/delivery/http"
	mpr "github.com/iamaul/fatbellies/app/meal_plan/repository"
	mpu "github.com/iamaul/fatbellies/app/meal_plan/usecase"

	"github.com/iamaul/fatbellies/config"
	"github.com/iamaul/fatbellies/config/database"
	"github.com/iamaul/fatbellies/config/migrations"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title Fatbellies API
// @version 1.0
// @description Project Assignment Xcidic.

// @contact.name iamaul
// @contact.url https://iamaul.me
// @contact.email iamaul@hotmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:5000

func main() {
	config := config.NewConfig()

	dbConnection, err := database.ConnectDatabase(config)
	if err != nil {
		log.Fatal(err)
	}

	// Migrate tables
	migrations.Migrate(dbConnection)

	// ToDo: Redis connection

	e := echo.New()

	// appMiddl := middleware.InitAppMiddleware(config.AppName)
	// e.Use(appMiddl.CORS)
	corsMiddl := middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXRequestedWith},
	}
	e.Use(middleware.CORSWithConfig(corsMiddl))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": http.StatusAccepted,
			"app":    "fatbellies-backend-api-services",
		})
	})

	// Branch
	branchRepo := br.NewBranchRepository(dbConnection)
	branchCase := bu.NewBranchUsecase(branchRepo)
	// Plan
	mealPlanRepo := mpr.NewMealPlanRepository(dbConnection)
	mealPlanCase := mpu.NewMealPlanUsecase(mealPlanRepo)

	// Branch
	bh.NewBranchHandler(e, branchCase)
	// Plan
	mph.NewMealPlanHandler(e, mealPlanCase)

	// Swagger docs
	e.GET("/api/docs/*any", echoSwagger.WrapHandler)

	log.Fatal(e.Start(fmt.Sprintf(`%s`, config.AppPort)))
}
