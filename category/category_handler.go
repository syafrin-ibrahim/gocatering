package category

import (
	"gocatering/helper"
	"gocatering/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type CategoryHandler struct {
	service Service
}

func NewCategoryHandler(s Service) *CategoryHandler {
	return &CategoryHandler{service: s}
}

func (h *CategoryHandler) CreateCategory(e echo.Context) error {
	var category model.Category
	if err := e.Bind(&category); err != nil {
		return e.JSON(http.StatusUnprocessableEntity,
			helper.Apiresponse(err.Error(), http.StatusUnprocessableEntity, "failed", nil))
	}
	err := h.service.CreateCategory(&category)
	if err != nil {
		return e.JSON(http.StatusInternalServerError,
			helper.Apiresponse(err.Error(), http.StatusInternalServerError, "failed", nil))
	}
	return e.JSON(http.StatusCreated,
		helper.Apiresponse("category Created", http.StatusOK, "success", category))

}

func (h *CategoryHandler) GetCategoryByID(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	res, err := h.service.GetCategoryByID(id)

	if err != nil {
		return e.JSON(http.StatusInternalServerError,
			helper.Apiresponse(err.Error(), http.StatusInternalServerError, "failed", nil))
	}
	if res == nil {
		return e.JSON(http.StatusNotFound,
			helper.Apiresponse(err.Error(), http.StatusNotFound, "Category Not Found", nil))
	}

	return e.JSON(http.StatusCreated,
		helper.Apiresponse("Category Found", http.StatusOK, "success", res))
}

func (h *CategoryHandler) GetAllCategory(e echo.Context) error {

	res, err := h.service.GetAllCategory()

	if err != nil {
		return e.JSON(http.StatusInternalServerError,
			helper.Apiresponse(err.Error(), http.StatusInternalServerError, "failed", nil))
	}

	return e.JSON(http.StatusCreated,
		helper.Apiresponse("List of Category", http.StatusOK, "success", res))

}

func (h *CategoryHandler) UpdateCategory(e echo.Context) error {
	var category model.Category
	id, _ := strconv.Atoi(e.Param("id"))

	if err := e.Bind(&category); err != nil {
		return e.JSON(http.StatusUnprocessableEntity,
			helper.Apiresponse(err.Error(), http.StatusUnprocessableEntity, "failed", nil))
	}

	errors := h.service.UpdateCategory(id, &category)

	if errors != nil {
		return e.JSON(http.StatusInternalServerError,
			helper.Apiresponse(errors.Error(), http.StatusInternalServerError, "failed", nil))
	}

	return e.JSON(http.StatusCreated,
		helper.Apiresponse("Category Updated", http.StatusOK, "success", category))
}

func (h *CategoryHandler) DeleteCategory(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	deletedCategory, errs := h.service.DeleteCategory(id)
	if errs != nil {
		return e.JSON(http.StatusInternalServerError,
			helper.Apiresponse(errs.Error(), http.StatusInternalServerError, "failed", nil))
	}

	return e.JSON(http.StatusCreated,
		helper.Apiresponse("Category Deleted", http.StatusOK, "success", deletedCategory))
}
