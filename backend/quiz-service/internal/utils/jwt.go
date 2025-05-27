package utils

import (
	"errors"
	models "internal/internal/model"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your_secret_key")

type Claims struct {
	Email  string `json:"email"`
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userID string, role string, email string) (string, error) {
	claims := Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ExtractUserFromRequest(r *http.Request) (models.User, error) {
	ctx := r.Context().Value("user")
	if ctx == nil {
		return models.User{}, errors.New("unauthenticated")
	}

	user, ok := ctx.(models.User)
	if !ok {
		return models.User{}, errors.New("invalid user type")
	}

	return user, nil
}
