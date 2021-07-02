package main

import (
	///"gocatering/handler"

	"gocatering/category"
	middle "gocatering/middleware"
	"gocatering/model"
	"gocatering/paket"
	"gocatering/transaction"
	"gocatering/user"

	"gocatering/regency"

	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(e echo.Context) error {
// 		hToken := e.Request().Header.Get("Authorization")
// 		token := strings.Split(hToken, " ")[1]
// 		claims := jwt.MapClaims{}
// 		_, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
// 			return []byte(middle.SECRET_JWT), nil
// 		})
// 		if err != nil {
// 			return echo.NewHTTPError(http.StatusInternalServerError, "gagal brooo")
// 		}
// 		// if !claims["authorized"].(bool) {
// 		// 	return echo.NewHTTPError(http.StatusForbidden, "forbidden page not authorized, just for admin")
// 		// }
// 		return next(e)
// 	}
// }

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/gocatering?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&model.User{}, &model.Regency{}, &model.Category{}, &model.Paket{}, &model.Image{}, &model.Transaction{})
	//	db.AutoMigrate(&user.User{})

	if err != nil {
		log.Fatal(err.Error())
	}

	regencyRepository := regency.NewRegencyRepository(db)
	regencyService := regency.NewRegencyService(regencyRepository)
	regencyHandler := regency.NewRegencyHandler(regencyService)
	categoryRepository := category.NewCategoryRepository(db)
	categoryService := category.NewCategoryService(categoryRepository)
	categoryHandler := category.NewCategoryHandler(categoryService)

	paketRepository := paket.NewPaketRepository(db)
	paketService := paket.NewPaketService(paketRepository)
	paketHandler := paket.NewPaketHandler(paketService)

	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)

	transRepository := transaction.NewTransactionRepository(db)
	transService := transaction.NewTransactionService(transRepository)
	transHandler := transaction.NewTransactionHandler(transService)

	e := echo.New()
	route := e.Group("/")
	route.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(middle.SECRET_JWT),
		TokenLookup: "header:authorization",
	}))
	route.POST("regency", regencyHandler.Createregency, middle.AuthMiddleware, middle.IsAdmin)
	route.GET("regency/:id", regencyHandler.GetRegencyByID, middle.AuthMiddleware, middle.IsAdmin)
	route.GET("regencies", regencyHandler.GetAllRegency, middle.AuthMiddleware, middle.IsAdmin)
	route.PUT("regency/:id", regencyHandler.UpdateRegency, middle.AuthMiddleware, middle.IsAdmin)
	route.DELETE("regency/:id", regencyHandler.DeleteRegency, middle.AuthMiddleware, middle.IsAdmin)
	route.POST("category", categoryHandler.CreateCategory, middle.AuthMiddleware, middle.IsAdmin)
	route.PUT("category/:id", categoryHandler.UpdateCategory, middle.AuthMiddleware, middle.IsAdmin)
	route.DELETE("category/:id", categoryHandler.DeleteCategory, middle.AuthMiddleware, middle.IsAdmin)
	route.GET("category/:id", categoryHandler.GetCategoryByID, middle.AuthMiddleware, middle.IsAdmin)
	route.POST("paket", paketHandler.CreatePaket, middle.AuthMiddleware, middle.IsAdmin)
	route.GET("paket/:id", paketHandler.GetPaketByID, middle.AuthMiddleware, middle.IsAdmin)
	route.PUT("paket/:id", paketHandler.UpdatePaket, middle.AuthMiddleware, middle.IsAdmin)
	route.DELETE("paket/:id", paketHandler.DeletePaket, middle.AuthMiddleware, middle.IsAdmin)
	route.POST("image", paketHandler.CreateImage, middle.AuthMiddleware, middle.IsAdmin)
	route.DELETE("image/:id", paketHandler.DeleteImage, middle.AuthMiddleware, middle.IsAdmin)
	route.PUT("updateprofile", userHandler.UpdateProfile, middle.AuthMiddleware)
	route.GET("transactions", transHandler.GetAllTransaction, middle.AuthMiddleware, middle.IsAdmin)

	route.POST("transaction", transHandler.CreateTransaction, middle.AuthMiddleware)
	route.POST("upload", userHandler.UploadImage, middle.AuthMiddleware)
	route.GET("transactions/user", transHandler.FindTransactionByUserId, middle.AuthMiddleware)

	e.POST("transaction/notif", transHandler.GetNotif)

	e.POST("registrasi", userHandler.RegisterUser)
	e.POST("login", userHandler.LoginUser)
	e.GET("pakets", paketHandler.GetAllPaket)
	e.GET("detailpaket/:id", paketHandler.DetailPaket)

	e.GET("categories", categoryHandler.GetAllCategory)

	e.Start(":8000")

}
