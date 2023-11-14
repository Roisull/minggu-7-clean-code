package delivery

import (
	"belajar-go-echo/app/model"
	"belajar-go-echo/app/usecase"

	"github.com/labstack/echo/v4"
	"net/http"
)

// handler untuk entitas User
type UserDelivery struct{
	userUsecase usecase.UserUsecase
}

// membuat instance UserDelivery
func NewUserDelivery(userUsecase *usecase.UserUsecase) *UserDelivery {
	return &UserDelivery{
		userUsecase: *userUsecase,
	}
}

// handler untuk request mendapatkan semua user
func (d *UserDelivery) GetAllUsers(c echo.Context) error{
	users, err := d.userUsecase.GetAllUsers()
	if err != nil{
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "ada salah kocak",
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get all data",
		"data": users,
	})
}

// handler untuk request membuat user baru
func (d *UserDelivery) CreateUser(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	createdUser, err := d.userUsecase.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "User Success Registered",
		"data": createdUser,
	})
}

// route user dengan echo
func (d *UserDelivery) RegisterRoutes(e *echo.Echo) {
	e.GET("/users", d.GetAllUsers)
	e.POST("/users", d.CreateUser)
}