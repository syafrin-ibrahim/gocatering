package user

import (
	"gocatering/helper"
	"gocatering/model"
	"net/http"

	"github.com/labstack/echo"
)

type UserHandler struct {
	service Service
}

func NewUserHandler(s Service) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) RegisterUser(e echo.Context) error {
	var user model.User

	if err := e.Bind(&user); err != nil {
		return e.JSON(http.StatusUnprocessableEntity,
			helper.Apiresponse(err.Error(), http.StatusUnprocessableEntity, "failed", nil))
	}

	currentUser, err := h.service.FindUserByEmail(user.Email)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed check email",
			"error":   err.Error(),
		})
	}

	if currentUser.ID != 0 {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "email already in use",
			"error":   "opsss..... get error",
		})
	}

	hashPass, errors := helper.HashPassword(user.Password)
	if errors != nil {

		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "eror generate password",
			"error":   errors.Error,
		})
	}

	newPass := string(hashPass)
	newUser := model.User{
		FullName: user.FullName,
		Mobile:   user.FullName,
		Address:  user.Address,
		Email:    user.Email,
		Password: newPass,
	}

	getErrs := h.service.RegisterUser(&newUser)

	if getErrs != nil {
		return e.JSON(http.StatusInternalServerError,
			helper.Apiresponse(err.Error(), http.StatusInternalServerError, "failed", nil))
	}
	return e.JSON(http.StatusCreated,
		helper.Apiresponse("User Registered", http.StatusOK, "success", newUser))

}

func (h *UserHandler) LoginUser(e echo.Context) error {
	var user model.User
	if err := e.Bind(&user); err != nil {
		return e.JSON(http.StatusUnprocessableEntity,
			helper.Apiresponse(err.Error(), http.StatusUnprocessableEntity, "failed", nil))
	}

	currentUser, err := h.service.FindUserByEmail(user.Email)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed check email",
			"error":   err.Error(),
		})
	}

	if currentUser.ID == 0 {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "email not registered",
			"error":   "opsss..... get error",
		})
	}

	match, error := helper.CheckPasswordHash(user.Password, currentUser.Password)
	if !match {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "your password not match",
			"error":   error.Error(),
		})
	}

	return e.JSON(http.StatusCreated,
		helper.Apiresponse("Login Success", http.StatusOK, "success", currentUser))
}
