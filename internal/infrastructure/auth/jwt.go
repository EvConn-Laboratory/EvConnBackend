package auth

import (
	"errors"
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type jwtAuth struct {
	secretKey string
	expired   time.Duration
}

type JWTClaim struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func NewJWTAuth(secretKey string, expired time.Duration) ports.JWTAuth {
	return &jwtAuth{
		secretKey: secretKey,
		expired:   expired,
	}
}

func (j *jwtAuth) GenerateToken(userID uint, role string) (string, error) {
	claims := &JWTClaim{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.expired)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtAuth) ValidateToken(tokenString string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return &models.User{
		Base: models.Base{ID: claims.UserID},
		Role: claims.Role,
	}, nil
}
