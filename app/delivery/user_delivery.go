package delivery

import (
	"belajar-go-echo/app/model"
	"belajar-go-echo/app/usecase"

	"net/http"

	"github.com/labstack/echo/v4"
	m "belajar-go-echo/app/middleware"
)

// user delivery interface
type UserDelivery interface {
	RegisterRoutes(e *echo.Echo)
	GetAllUsers(c echo.Context) error
   	CreateUser(c echo.Context) error
   	GetAuthenticatedUsers(c echo.Context) error
}

// handler untuk entitas User
type UserDeliveryImpl struct{
	userUsecase usecase.UserUsecase
}

// membuat instance UserDelivery
func NewUserDelivery(userUsecase usecase.UserUsecase) *UserDeliveryImpl {
	return &UserDeliveryImpl{
		userUsecase: userUsecase,
	}
}

// handler untuk request mendapatkan semua user dengan auth
func (d *UserDeliveryImpl) GetAuthenticatedUsers(c echo.Context) error {
	// Panggil metode dari use case atau interaksi aplikasi
	authUsers, err := d.userUsecase.GetAuthenticatedUsers()
	if err != nil {
	   return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "gakpunya jwt lu",
			"error": "Internal Server Error",
		})
	}
 
	return c.JSON(http.StatusOK, authUsers)
 }

// handler untuk request mendapatkan semua user
func (d *UserDeliveryImpl) GetAllUsers(c echo.Context) error{
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
func (d *UserDeliveryImpl) CreateUser(c echo.Context) error {
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

// RegisterRoutes route user dengan echo
func (d *UserDeliveryImpl) RegisterRoutes(e *echo.Echo) {
	// route perlu auth
	// authenticated := e.Group("/authent", middleware.AuthMiddleware)
	e.GET("/users/authent", d.GetAuthenticatedUsers)

	e.GET("/users", d.GetAllUsers)
	e.POST("/users", d.CreateUser)

	// log middleware
	m.LogMiddleware(e)
}