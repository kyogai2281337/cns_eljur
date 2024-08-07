package server

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserJWT struct {
	ID      int64
	Email   string
	Role    string
	ExpTime time.Time
}

var (
	errInvalidToken = errors.New("invalid token")
)

const (
	hashKey string = "max_verstrappen"
)

func GetUserDataJWT(prompt string) (*UserJWT, error) {
	token, err := jwt.Parse(prompt, func(token *jwt.Token) (interface{}, error) {
		return []byte(hashKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errInvalidToken
	}
	res := &UserJWT{}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}
	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, errors.New("invalid 'exp' claim")
	}
	res.ExpTime = time.Unix(int64(exp), 0)

	id, ok := claims["id"].(float64)
	if !ok {
		return nil, errors.New("invalid 'id' claim")
	}
	res.ID = int64(id)

	email, ok := claims["email"].(string)
	if !ok {
		return nil, errors.New("invalid 'email' claim")
	}
	res.Email = email

	role, ok := claims["role"].(string)
	if !ok {
		return nil, errors.New("invalid 'role' claim")
	}
	res.Role = role

	return res, nil
}
