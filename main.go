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
	route.POST("regency", regencyHandler.Createregency)
	route.GET("regency/:id", regencyHandler.GetRegencyByID)
	route.GET("regencies", regencyHandler.GetAllRegency)
	route.PUT("regency/:id", regencyHandler.UpdateRegency)
	route.DELETE("regency/:id", regencyHandler.DeleteRegency)

	route.POST("category", categoryHandler.CreateCategory)
	route.GET("category/:id", categoryHandler.GetCategoryByID)
	route.GET("categories", categoryHandler.GetAllCategory)
	route.PUT("category/:id", categoryHandler.UpdateCategory)
	route.DELETE("category/:id", categoryHandler.DeleteCategory)

	route.POST("paket", paketHandler.CreatePaket)
	route.GET("paket/:id", paketHandler.GetPaketByID)
	route.GET("pakets", paketHandler.GetAllPaket)
	route.PUT("paket/:id", paketHandler.UpdatePaket)
	route.DELETE("paket/:id", paketHandler.DeletePaket)
	route.POST("image", paketHandler.CreateImage)
	route.DELETE("image/:id", paketHandler.DeleteImage)

	e.POST("registrasi", userHandler.RegisterUser)
	e.POST("login", userHandler.LoginUser)

	e.POST("transaction", transHandler.CreateTransaction)
	e.GET("transactions", transHandler.GetAllTransaction, middle.AuthMiddleware)
	e.GET("transactions/user", transHandler.FindTransactionByUserId)
	e.POST("transaction/notif", transHandler.GetNotif)
	e.Start(":8000")

}
