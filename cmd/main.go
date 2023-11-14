package main

import (
	"belajar-go-echo/app/delivery"
	"belajar-go-echo/app/repository"
	"belajar-go-echo/app/usecase"
	"belajar-go-echo/config"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()

	// inisialisasi Repository, usecase dan delivery
	userRepositoryImpl := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepositoryImpl)
	userDelivery := delivery.NewUserDelivery(userUsecase)

	app := echo.New()

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
