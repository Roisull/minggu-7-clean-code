package auth

import (
	"belajar-go-echo/app/model"
	"time"

	"github.com/golang-jwt/jwt"
)

const jwtSecret = "rois-jwt-secret"

// GenerateToken JWT
func GenerateToken(user *model.User) (string, error) {
	claims := jwt.MapClaims{
	   "user_id": user.ID,
	   "exp":     time.Now().Add(time.Hour * 24).Unix(), // Token berlaku selama 1 hari
	}
 
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
	   return "", err
	}
 
	return tokenString, nil
 }
 
 // VerifyToken
 func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	   return jwtSecret, nil
	})
 
	if err != nil {
	   return nil, err
	}
 
	return token, nil
 }