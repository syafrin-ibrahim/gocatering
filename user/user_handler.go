package user

import (
	"gocatering/helper"
	"gocatering/middleware"
	"gocatering/model"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
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

	token, errorToken := middleware.CreateToken(newUser.ID, user.IsAdmin)
	if errorToken != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "eror generate token",
			"error":   errors.Error,
		})
	}
	userResponse := model.UserResponse{
		UserName: newUser.FullName,
		Mobile:   newUser.Mobile,
		Email:    newUser.Email,
		Token:    token,
	}

	e.Response().Header().Set("Authorization", "Bearer "+token)

	return e.JSON(http.StatusCreated,
		helper.Apiresponse("User Registered", http.StatusOK, "success", userResponse))

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

	token, errors := middleware.CreateToken(currentUser.ID, currentUser.IsAdmin)

	if errors != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error generate token",
			"error":   err.Error(),
		})
	}

	userResponse := model.UserResponse{
		UserName: currentUser.FullName,
		Mobile:   currentUser.Mobile,
		Email:    currentUser.Email,
		Token:    token,
	}
	//set header
	//e.Set("user_id", currentUser.ID)
	e.Response().Header().Set("Authorization", "Bearer "+token)
	// e.Response().Header().Set("user_id", currentUser.Email)

	return e.JSON(http.StatusCreated,
		helper.Apiresponse("Login Success", http.StatusOK, "success", userResponse))
}

func (h *UserHandler) UpdateProfile(e echo.Context) error {
	user := model.User{}
	sessionUser := e.Get("user").(*jwt.Token)
	claims := sessionUser.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)

	getUser, errors := h.service.FindUserById(int(userId))

	e.Bind(&user)

	hashPass, errors := helper.HashPassword(user.Password)

	if errors != nil {

		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "eror generate password",
			"error":   errors.Error,
		})
	}

	newPass := string(hashPass)
	getUser.Password = newPass
	getUser.Email = user.Email
	getUser.Mobile = user.Mobile
	getUser.FullName = user.FullName
	getUser.Address = user.Address

	err := h.service.UpdateUser(int(userId), getUser)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "eror update user",
			"error":   errors.Error,
		})
	}
	userResponse := model.UserResponse{
		UserName: getUser.FullName,
		Mobile:   getUser.Mobile,
		Email:    getUser.Email,
	}

	response := helper.Apiresponse("Updated Profile ok", http.StatusOK, "success", userResponse)
	return e.JSON(http.StatusOK, response)

}

func (h *UserHandler) UploadImage(e echo.Context) error {
	sessionUser := e.Get("user").(*jwt.Token)
	claims := sessionUser.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)

	getUser, errors := h.service.FindUserById(int(userId))
	if errors != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error sql",
			"error":   errors.Error(),
		})
	}

	userName := strings.Replace(getUser.FullName, " ", "_", -1)
	file, err := e.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	filePath := "./avatar/" + userName + "_" + file.Filename
	fileSrc := "avatar/" + userName + "_" + file.Filename

	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	getUser.ImageUrl = fileSrc

	getErrors := h.service.UpdateUser(int(userId), getUser)

	if getErrors != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "eror upload image",
			"error":   errors.Error,
		})
	}
	userResponse := model.UserResponse{
		UserName: getUser.FullName,
		Mobile:   getUser.Mobile,
		Email:    getUser.Email,
	}

	response := helper.Apiresponse("Updated Profile ok", http.StatusOK, "success", userResponse)
	return e.JSON(http.StatusOK, response)

}
