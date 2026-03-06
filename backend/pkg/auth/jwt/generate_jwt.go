package jwtEnc

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var jwtKey = []byte(os.Getenv("JWT_SECURE_KEY")) // Use a strong, secure key

type Claims struct {
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	Role     string    `json:"role"`
	Exp      time.Time `json:"expired"`
	jwt.RegisteredClaims
}

type User struct {
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	Role     string    `json:"role"`
	Exp      time.Time `json:"expired"`
}

func GenerateJWT(user User) (string, time.Time, error) {
	exp, _ := strconv.Atoi(os.Getenv("JWT_EXPIRED_HOUR"))
	expirationTime := time.Now().Add(time.Duration(exp) * time.Hour)

	claims := &Claims{
		UserID:   user.UserID,
		Username: user.Username,
		Role:     user.Role,
		Exp:      expirationTime,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    os.Getenv("APP_NAME"),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expirationTime, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		fmt.Println(err)
		return nil, fmt.Errorf("unauthorization")
	}

	return claims, nil
}
