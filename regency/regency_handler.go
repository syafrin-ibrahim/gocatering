package regency

import (
	"gocatering/helper"
	"gocatering/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type RegencyHandler struct {
	RegencyService Service
}

func NewRegencyHandler(r Service) *RegencyHandler {
	return &RegencyHandler{RegencyService: r}
}

func (h *RegencyHandler) Createregency(e echo.Context) error {
	var regency model.Regency
	if err := e.Bind(&regency); err != nil {
		return e.JSON(http.StatusUnprocessableEntity,
			helper.Apiresponse(err.Error(), http.StatusUnprocessableEntity, "failed", nil))
	}
	err := h.RegencyService.CreateRegency(&regency)
	if err != nil {
		return e.JSON(http.StatusInternalServerError,
			helper.Apiresponse(err.Error(), http.StatusInternalServerError, "failed", nil))
	}
	return e.JSON(http.StatusCreated,
		helper.Apiresponse("Regency Created", http.StatusOK, "success", regency))

}

func (h *RegencyHandler) GetRegencyByID(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	res, err := h.RegencyService.GetRegencyByID(id)

	if err != nil {
		return e.JSON(http.StatusInternalServerError,
			helper.Apiresponse(err.Error(), http.StatusInternalServerError, "failed", nil))
	}
	if res == nil {
		return e.JSON(http.StatusNotFound,
			helper.Apiresponse(err.Error(), http.StatusNotFound, "Regency Not Found", nil))
	}

	return e.JSON(http.StatusCreated,
		helper.Apiresponse("Regency Found", http.StatusOK, "success", res))
}

func (h *RegencyHandler) GetAllRegency(e echo.Context) error {

	res, err := h.RegencyService.GetAllRegency()

	if err != nil {
		return e.JSON(http.StatusInternalServerError,
			helper.Apiresponse(err.Error(), http.StatusInternalServerError, "failed", nil))
	}

	return e.JSON(http.StatusCreated,
		helper.Apiresponse("List of Regency", http.StatusOK, "success", res))

}

func (h *RegencyHandler) UpdateRegency(e echo.Context) error {
	var regency model.Regency
	id, _ := strconv.Atoi(e.Param("id"))

	if err := e.Bind(&regency); err != nil {
		return e.JSON(http.StatusUnprocessableEntity,
			helper.Apiresponse(err.Error(), http.StatusUnprocessableEntity, "failed", nil))
	}

	errors := h.RegencyService.UpdateRegency(id, &regency)

	if errors != nil {
		return e.JSON(http.StatusInternalServerError,
			helper.Apiresponse(errors.Error(), http.StatusInternalServerError, "failed", nil))
	}

	return e.JSON(http.StatusCreated,
		helper.Apiresponse("Regency Updated", http.StatusOK, "success", regency))
}

func (h *RegencyHandler) DeleteRegency(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	deletedReg, errs := h.RegencyService.DeleteRegency(id)
	if errs != nil {
		return e.JSON(http.StatusInternalServerError,
			helper.Apiresponse(errs.Error(), http.StatusInternalServerError, "failed", nil))
	}

	return e.JSON(http.StatusCreated,
		helper.Apiresponse("Regency Deleted", http.StatusOK, "success", deletedReg))
}
