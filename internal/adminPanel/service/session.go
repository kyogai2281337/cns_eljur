package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
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

func (u *UserJWT) GenJWT(t string) (string, error) {
	var exptime int64
	switch t {
	case "a":
		exptime = time.Now().Add(time.Hour * 24).Unix()
	case "r":
		exptime = time.Now().Add(time.Hour * (24 * 7)).Unix()
	default:
		exptime = time.Now().Add(time.Hour * 24).Unix()
	}

	prompt := jwt.New(jwt.SigningMethodHS256)
	claims := prompt.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["email"] = u.Email
	claims["role"] = u.Role
	claims["exp"] = exptime
	token, err := prompt.SignedString([]byte(hashKey))
	if err != nil {
		return "", err
	}
	log.Printf("[200] => CREATED %sJWT FOR %s", t, u.Email)
	return token, nil
}
