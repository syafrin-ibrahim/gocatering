package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

const SECRET_JWT = "1234567"

func CreateToken(userId int, isadmin bool) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["is_admin"] = isadmin
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SECRET_JWT))
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		hToken := e.Request().Header.Get("Authorization")
		tokenString := ""
		arrayToken := strings.Split(hToken, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		//token := strings.Split(hToken, " ")[1]
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(tokenString, claims, func(*jwt.Token) (interface{}, error) {
			return []byte(SECRET_JWT), nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "opsss....galat anda belum terdaftar dalam aplikasi")
		}
		// if !claims["authorized"].(bool) {
		// 	return echo.NewHTTPError(http.StatusForbidden, "forbidden page not authorized, just for admin")
		// }
		return next(e)
	}
}

func isAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		isAdmin := claims["is_admin"].(bool)
		if isAdmin == false {
			return echo.NewHTTPError(http.StatusForbidden, "forbidden page not authorized, just for admin")
		}
		return next(c)
	}
}
