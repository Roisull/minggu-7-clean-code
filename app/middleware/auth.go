package middleware

import (
	"belajar-go-echo/app/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

// AuthMiddleware adalah middleware untuk memeriksa keberlanjutan token JWT
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
	   tokenString := c.Request().Header.Get("Authorization")
	   if tokenString == "" {
		  return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Unauthorized"})
	   }
 
	   token, err := auth.VerifyToken(tokenString)
	   if err != nil || !token.Valid {
		  return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token"})
	   }
 
	   return next(c)
	}
 }