package main

import (
	"belajar-go-echo/app/delivery"
	"belajar-go-echo/app/repository"
	"belajar-go-echo/app/usecase"
	"belajar-go-echo/config"

	// "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
	// "github.com/labstack/echo-contrib/jwt"
	mid "github.com/labstack/echo-jwt"
)

func main() {
	db := config.InitDB()

	// inisialisasi Repository, usecase dan delivery
	userRepositoryImpl := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepositoryImpl)
	userDelivery := delivery.NewUserDelivery(*userUsecase)

	app := echo.New()

	// middleware untuk auth JWT

	app.Use(mid.JWT([]byte("rois-jwt-secret")))
	// config := middleware.JWTConfig{
	// 	Claims:     &jwt.StandardClaims{},
	// 	SigningKey: []byte("rois-jwt-secret"),
	// }

	// app.Use(middleware.JWTWithConfig(config))

	// routes
	userDelivery.RegisterRoutes(app)

	// start echo server
	err := app.Start(":8000")
	if err != nil {
		panic(err)
	}

	// err = config.MigrateDB(db)
	// if err != nil {
	// 	panic(err)
	// }

	// app := echo.New()
	// app.GET("/users", controller.GetAllUsers(db))
	// app.POST("/users", controller.CreateUser(db))
	// app.Start(":8080")
}
