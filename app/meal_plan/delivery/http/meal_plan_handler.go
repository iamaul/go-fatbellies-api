package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	mealPlan "github.com/iamaul/fatbellies/app/meal_plan"
	"github.com/iamaul/fatbellies/app/models"
	"github.com/iamaul/fatbellies/utils"
	"github.com/labstack/echo/v4"
)

type MealPlanHandler struct {
	Mealplancase mealPlan.Usecase
}

func NewMealPlanHandler(e *echo.Echo, mpu mealPlan.Usecase) {
	handler := &MealPlanHandler{
		Mealplancase: mpu,
	}

	g := e.Group("/api")
	g.GET("/mealplans", handler.Fetch)
	g.GET("/mealplans/:id", handler.GetByID)
	g.GET("/mealplans/meal/:name", handler.GetByName)
	g.POST("/mealplans", handler.Store)
	g.DELETE("/delete/mealplans/:id", handler.Delete)
	g.PUT("/update/mealplans/:id", handler.Update)
	g.POST("/search/mealplans", handler.SearchPlans)
}

// @Summary List meal plan
// @Description Get a list of meal plans
// @Tags Meal Plans
// @Accept  json
// @Produce  json
// @Param limit query integer 5 "limit numbers"
// @Param page query integer 1 "pagination"
// @Param order query string false "created_at desc"
// @Success 200 {array} models.MealPlan
// @Router /api/mealplans [get]
func (mph *MealPlanHandler) Fetch(c echo.Context) error {
	queryLimit := c.QueryParam("limit")
	limit, _ := strconv.Atoi(queryLimit)
	queryPage := c.QueryParam("page")
	page, _ := strconv.Atoi(queryPage)
	queryOrder := c.QueryParam("order")

	res, err := mph.Mealplancase.Fetch(int64(limit), int64(page), queryOrder)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.ResponseJSON{
			Code:    http.StatusInternalServerError,
			Message: "An unexpected error has occurred",
			Error:   err.Error(),
			Success: false,
		})
	}

	return c.JSON(http.StatusOK, &utils.ResponseJSON{
		Code:    http.StatusOK,
		Result:  res,
		Message: "Fetched data successfully",
		Success: true,
	})
}

// @Summary Find one of all the meal plans
// @Description Get meal plan by ID
// @Tags Meal Plans
// @Accept  json
// @Produce  json
// @Param id path string uuid "Meal Plan ID"
// @Success 200 {array} models.MealPlan
// @Router /api/mealplans/{id} [get]
func (mph *MealPlanHandler) GetByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.ResponseJSON{
			Code:    http.StatusBadRequest,
			Message: "An unexpected error has occurred",
			Error:   err.Error(),
			Success: false,
		})
	}

	res, err := mph.Mealplancase.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusConflict, &utils.ResponseJSON{
			Code:    http.StatusConflict,
			Error:   err.Error(),
			Success: false,
		})
	}

	return c.JSON(http.StatusOK, &utils.ResponseJSON{
		Code:    http.StatusOK,
		Result:  res,
		Message: "Fetched data by id",
		Success: true,
	})
}

// @Summary Find one of all the meal plans
// @Description Get meal plan by Name
// @Tags Meal Plans
// @Accept  json
// @Produce  json
// @Param name path string false "Meal plan name"
// @Success 200 {array} models.MealPlan
// @Router /api/mealplans/meal/{name} [get]
func (mph *MealPlanHandler) GetByName(c echo.Context) error {
	planName := c.Param("name")

	res, err := mph.Mealplancase.GetByName(planName)
	if err != nil {
		return c.JSON(http.StatusConflict, &utils.ResponseJSON{
			Code:    http.StatusConflict,
			Error:   err.Error(),
			Success: false,
		})
	}

	return c.JSON(http.StatusOK, &utils.ResponseJSON{
		Code:    http.StatusOK,
		Result:  res,
		Message: "Fetched data by name",
		Success: true,
	})
}

func createMealPlanValidation(cb *models.MealPlan) (bool, error) {
	validate := validator.New()

	err := validate.Struct(cb)
	if err != nil {
		return false, err
	}
	return true, nil
}

// @Summary Add meal plan
// @Description Add new meal plan
// @Tags Meal Plans
// @Accept  json
// @Produce  json
// @Param meal_plan body models.SwagMealPlan true "Form JSON"
// @Success 200 {array} models.MealPlan
// @Router /api/mealplans [post]
func (mph *MealPlanHandler) Store(c echo.Context) error {
	var mealPlan models.MealPlan

	err := c.Bind(&mealPlan)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &utils.ResponseJSON{
			Code:    http.StatusUnprocessableEntity,
			Message: "An unexpected error has occurred",
			Error:   err.Error(),
			Success: false,
		})
	}

	if ok, err := createMealPlanValidation(&mealPlan); !ok {
		return c.JSON(http.StatusBadRequest, &utils.ResponseJSON{
			Code:    http.StatusBadRequest,
			Message: "Validation invalid",
			Error:   err.Error(),
			Success: false,
		})
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")

	// set timezone,
	now := time.Now().In(loc)
	end := time.Now().In(loc).Add(3)

	mealPlan.StartTime = now
	mealPlan.EndTime = end

	res, err := mph.Mealplancase.Store(&mealPlan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.ResponseJSON{
			Code:    http.StatusInternalServerError,
			Message: "An unexpected error has occurred",
			Error:   err.Error(),
			Success: false,
		})
	}

	return c.JSON(http.StatusCreated, &utils.ResponseJSON{
		Code:    http.StatusCreated,
		Result:  res,
		Message: "Plan created successfully",
		Success: true,
	})
}

// @Summary Update meal plan
// @Description Update meal plan by ID
// @Tags Meal Plans
// @Accept  json
// @Produce  json
// @Param id path string uuid "Meal plan ID"
// @Param meal_plan body models.SwagMealPlan true "Form JSON"
// @Success 200 {array} models.SwagMealPlan
// @Router /api/update/mealplans/{id} [put]
func (mph *MealPlanHandler) Update(c echo.Context) error {
	var mealPlan models.MealPlan

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.ResponseJSON{
			Code:    http.StatusBadRequest,
			Message: "An unexpected error has occurred",
			Error:   err.Error(),
			Success: false,
		})
	}

	errBind := c.Bind(&mealPlan)
	if errBind != nil {
		return c.JSON(http.StatusUnprocessableEntity, &utils.ResponseJSON{
			Code:    http.StatusUnprocessableEntity,
			Message: "An unexpected error has occurred",
			Error:   errBind.Error(),
			Success: false,
		})
	}

	if ok, errValidation := createMealPlanValidation(&mealPlan); !ok {
		return c.JSON(http.StatusBadRequest, &utils.ResponseJSON{
			Code:    http.StatusBadRequest,
			Message: "Validation invalid",
			Error:   errValidation.Error(),
			Success: false,
		})
	}

	res, errPlan := mph.Mealplancase.Update(id, mealPlan)
	if errPlan != nil {
		return c.JSON(http.StatusNotFound, &utils.ResponseJSON{
			Code:    http.StatusNotFound,
			Error:   errPlan.Error(),
			Success: false,
		})
	}

	return c.JSON(http.StatusOK, &utils.ResponseJSON{
		Code:    http.StatusOK,
		Result:  res,
		Message: "Plan updated successfully",
		Success: true,
	})
}

// @Summary Delete one of all the meal plans
// @Description Delete meal plan by ID
// @Tags Meal Plans
// @Accept  json
// @Produce  json
// @Param id path string false "Meal plan ID"
// @Success 200
// @Router /api/delete/mealplans/{id} [delete]
func (mph *MealPlanHandler) Delete(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.ResponseJSON{
			Code:    http.StatusBadRequest,
			Message: "An unexpected error has occurred",
			Error:   err.Error(),
			Success: false,
		})
	}

	err = mph.Mealplancase.Delete(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &utils.ResponseJSON{
			Code:    http.StatusNotFound,
			Error:   err.Error(),
			Success: false,
		})
	}

	return c.JSON(http.StatusOK, &utils.ResponseJSON{
		Code:    http.StatusOK,
		Message: "Plan deleted successfully",
		Success: true,
	})
}

// @Summary Search meal plans
// @Description Search meal plan by specific queries column, q, and ID
// @Tags Meal Plans
// @Accept  json
// @Produce  json
// @Param column query string false "meal_plan_name"
// @Param q query string false "Buffet A"
// @Param order query string false "created_at desc"
// @Success 200 {array} models.MealPlan
// @Router /api/search/mealplans [post]
func (mph *MealPlanHandler) SearchPlans(c echo.Context) error {
	column := c.QueryParam("column")
	label := c.QueryParam("q")
	order := c.QueryParam("order")

	res, err := mph.Mealplancase.SearchPlans(column, label, order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &utils.ResponseJSON{
			Code:    http.StatusInternalServerError,
			Message: "An unexpected error has occurred",
			Error:   err.Error(),
			Success: false,
		})
	}

	return c.JSON(http.StatusOK, &utils.ResponseJSON{
		Code:    http.StatusOK,
		Result:  res,
		Message: "Filtered data successfully",
		Success: true,
	})
}
