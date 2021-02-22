package http

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/iamaul/fatbellies/app/branch"
	"github.com/iamaul/fatbellies/app/models"
	"github.com/iamaul/fatbellies/utils"
	"github.com/labstack/echo/v4"
)

type BranchHandler struct {
	Branchcase branch.Usecase
}

func NewBranchHandler(e *echo.Echo, bu branch.Usecase) {
	handler := &BranchHandler{
		Branchcase: bu,
	}

	g := e.Group("/api")
	g.GET("/branches", handler.Fetch)
	g.GET("/branches/:id", handler.GetByID)
	g.GET("/branches/branch/:name", handler.GetByName)
	g.POST("/branches", handler.Store)
	g.POST("/branches/mealplans", handler.StoreMealPlan)
	g.DELETE("/delete/branches/:id", handler.Delete)
	g.PUT("/update/branches/:id", handler.Update)
	g.POST("/search/branches", handler.SearchBranches)
	g.GET("/nearest/branches", handler.FindNearestLocation)
}

// @Summary List branches
// @Description Get a list of branches
// @Tags Branches
// @Accept  json
// @Produce  json
// @Param limit query integer 5 "limit numbers"
// @Param page query integer 1 "pagination"
// @Param order query string false "created_at desc"
// @Success 200 {array} models.Branch
// @Router /api/branches [get]
func (bh *BranchHandler) Fetch(c echo.Context) error {
	queryLimit := c.QueryParam("limit")
	limit, _ := strconv.Atoi(queryLimit)
	queryPage := c.QueryParam("page")
	page, _ := strconv.Atoi(queryPage)
	queryOrder := c.QueryParam("order")

	res, err := bh.Branchcase.Fetch(int64(limit), int64(page), queryOrder)
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

// @Summary Find one of all the branches
// @Description Get branch by ID
// @Tags Branches
// @Accept  json
// @Produce  json
// @Param id path string uuid "Branch ID"
// @Success 200 {array} models.Branch
// @Router /api/branches/{id} [get]
func (bh *BranchHandler) GetByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.ResponseJSON{
			Code:    http.StatusBadRequest,
			Message: "An unexpected error has occurred",
			Error:   err.Error(),
			Success: false,
		})
	}

	res, err := bh.Branchcase.GetByID(id)
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

// @Summary Find one of all the branches
// @Description Get branch by Name
// @Tags Branches
// @Accept  json
// @Produce  json
// @Param name path string false "Branch name"
// @Success 200 {array} models.Branch
// @Router /api/branches/branch/{name} [get]
func (bh *BranchHandler) GetByName(c echo.Context) error {
	branchName := c.Param("name")

	res, err := bh.Branchcase.GetByName(branchName)
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

// @Summary Find nearest location
// @Description Get nearest location between branch and user
// @Tags Branches
// @Accept  json
// @Produce  json
// @Param lat query string false "Latitude"
// @Param long query string false "Longitude"
// @Success 200 {array} models.BranchLocation
// @Router /api/nearest/branches [get]
func (bh *BranchHandler) FindNearestLocation(c echo.Context) error {
	queryLat := c.QueryParam("lat")
	latitude, _ := strconv.ParseFloat(queryLat, 64)
	queryLong := c.QueryParam("long")
	longitude, _ := strconv.ParseFloat(queryLong, 64)

	res, err := bh.Branchcase.FindNearestLocation(latitude, longitude)
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

func createBranchValidation(cb *models.Branch) (bool, error) {
	validate := validator.New()

	err := validate.Struct(cb)
	if err != nil {
		return false, err
	}
	return true, nil
}

// @Summary Add branch
// @Description Add new branch
// @Tags Branches
// @Accept  json
// @Produce  json
// @Param branch body models.SwagBranch true "Form JSON"
// @Success 200 {array} models.Branch
// @Router /api/branches [post]
func (bh *BranchHandler) Store(c echo.Context) error {
	var branch models.Branch

	err := c.Bind(&branch)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &utils.ResponseJSON{
			Code:    http.StatusUnprocessableEntity,
			Message: "An unexpected error has occurred",
			Error:   err.Error(),
			Success: false,
		})
	}

	if ok, err := createBranchValidation(&branch); !ok {
		return c.JSON(http.StatusBadRequest, &utils.ResponseJSON{
			Code:    http.StatusBadRequest,
			Message: "Validation invalid",
			Error:   err.Error(),
			Success: false,
		})
	}

	res, err := bh.Branchcase.Store(&branch)
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
		Message: "Branch created successfully",
		Success: true,
	})
}

func createBranchMealPlanValidation(cb *models.BranchMealPlan) (bool, error) {
	validate := validator.New()

	err := validate.Struct(cb)
	if err != nil {
		return false, err
	}
	return true, nil
}

// @Summary Add branch meal plan
// @Description Add new branch meal plan
// @Tags Branches
// @Accept  json
// @Produce  json
// @Param branch_meal_plan body models.BranchMealPlan true "Form JSON"
// @Success 200 {array} models.BranchMealPlan
// @Router /api/branches/mealplans [post]
func (bh *BranchHandler) StoreMealPlan(c echo.Context) error {
	var mealPlan models.BranchMealPlan

	err := c.Bind(&mealPlan)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &utils.ResponseJSON{
			Code:    http.StatusUnprocessableEntity,
			Message: "An unexpected error has occurred",
			Error:   err.Error(),
			Success: false,
		})
	}

	if ok, err := createBranchMealPlanValidation(&mealPlan); !ok {
		return c.JSON(http.StatusBadRequest, &utils.ResponseJSON{
			Code:    http.StatusBadRequest,
			Message: "Validation invalid",
			Error:   err.Error(),
			Success: false,
		})
	}

	err = bh.Branchcase.StoreMealPlan(&mealPlan)
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
		Message: "Branch meal plan added successfully",
		Success: true,
	})
}

// @Summary Update branch
// @Description Update branch by ID
// @Tags Branches
// @Accept  json
// @Produce  json
// @Param id path string uuid "Branch ID"
// @Param branch body models.SwagBranch true "Form JSON"
// @Success 200 {array} models.Branch
// @Router /api/update/branches/{id} [put]
func (bh *BranchHandler) Update(c echo.Context) error {
	var branch models.Branch

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.ResponseJSON{
			Code:    http.StatusBadRequest,
			Message: "An unexpected error has occurred",
			Error:   err.Error(),
			Success: false,
		})
	}

	errBind := c.Bind(&branch)
	if errBind != nil {
		return c.JSON(http.StatusUnprocessableEntity, &utils.ResponseJSON{
			Code:    http.StatusUnprocessableEntity,
			Message: "An unexpected error has occurred",
			Error:   errBind.Error(),
			Success: false,
		})
	}

	if ok, errValidation := createBranchValidation(&branch); !ok {
		return c.JSON(http.StatusBadRequest, &utils.ResponseJSON{
			Code:    http.StatusBadRequest,
			Message: "Validation invalid",
			Error:   errValidation.Error(),
			Success: false,
		})
	}

	res, errBranch := bh.Branchcase.Update(id, branch)
	if errBranch != nil {
		return c.JSON(http.StatusNotFound, &utils.ResponseJSON{
			Code:    http.StatusNotFound,
			Error:   errBranch.Error(),
			Success: false,
		})
	}

	return c.JSON(http.StatusOK, &utils.ResponseJSON{
		Code:    http.StatusOK,
		Result:  res,
		Message: "Branch updated successfully",
		Success: true,
	})
}

// @Summary Delete one of all the branches
// @Description Delete branch by ID
// @Tags Branches
// @Accept  json
// @Produce  json
// @Param id path string false "Branch ID"
// @Success 200
// @Router /api/delete/branches/{id} [delete]
func (bh *BranchHandler) Delete(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &utils.ResponseJSON{
			Code:    http.StatusBadRequest,
			Message: "An unexpected error has occurred",
			Error:   err.Error(),
			Success: false,
		})
	}

	err = bh.Branchcase.Delete(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &utils.ResponseJSON{
			Code:    http.StatusNotFound,
			Error:   err.Error(),
			Success: false,
		})
	}

	return c.JSON(http.StatusOK, &utils.ResponseJSON{
		Code:    http.StatusOK,
		Message: "Branch deleted successfully",
		Success: true,
	})
}

// @Summary Search branches
// @Description Search branch by specific queries column, q, and ID
// @Tags Branches
// @Accept  json
// @Produce  json
// @Param column query string false "branch_name"
// @Param q query string false "restaurant"
// @Param order query string false "created_at desc"
// @Success 200 {array} models.Branch
// @Router /api/search/branches [post]
func (bh *BranchHandler) SearchBranches(c echo.Context) error {
	column := c.QueryParam("column")
	label := c.QueryParam("q")
	order := c.QueryParam("order")

	res, err := bh.Branchcase.SearchBranches(column, label, order)
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
