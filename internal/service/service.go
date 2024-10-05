package service

import (
	"appGin/internal/models"
	"errors"
	"github.com/golang-jwt/jwt"
	"os"
	"time"

	"appGin/internal/repo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// RegisterUser handles registration logic
func RegisterUser(user *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := models.User{
		Username: user.Username,
		Password: string(hashedPassword),
	}

	return repo.CreateUser(&newUser)
}

// LoginUser handles login logic
func LoginUser(user *User) (string, error) {
	dbUser, err := repo.GetUserByUsername(user.Username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := generateToken(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func generateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// VerifyToken checks the validity of the JWT token
func VerifyToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid token signature")
		}
		return nil, errors.New("could not parse token")
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
