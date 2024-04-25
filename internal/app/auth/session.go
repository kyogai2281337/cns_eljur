package apiserver

import (
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kyogai2281337/cns_eljur/internal/app/model"
)

type UserJWT struct {
	ID      int64
	Email   string
	Role    string
	ExpTime time.Time
}

var (
	errInvalidToken = errors.New("invalid token")
	errExpiredToken = errors.New("token is expired")
)

const (
	hashKey string = "max_verstrappen"
)

func GenAuthJWT(u *model.User) (string, string, error) {
	access := jwt.New(jwt.SigningMethodHS256)
	claims := access.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["email"] = u.Email
	claims["role"] = u.Role.Name
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()    // Устанавливаем срок действия токена
	tokenAccess, err := access.SignedString([]byte(hashKey)) // здесь блять ключ поменяй сука
	if err != nil {
		return "", "", err
	}

	refresh := jwt.New(jwt.SigningMethodHS256)
	claims = refresh.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["email"] = u.Email
	claims["role"] = u.Role.Name
	claims["exp"] = time.Now().Add(time.Hour * (24 * 30)).Unix() // Устанавливаем срок действия токена
	tokenRefresh, err := refresh.SignedString([]byte(hashKey))   // здесь блять ключ поменяй сука
	if err != nil {
		return "", "", err
	}
	log.Printf("[200] => CREATED AJWT AND RJWT FOR %s", u.Email)
	return tokenAccess, tokenRefresh, nil
}

func GetUserDataJWT(prompt string) (*UserJWT, error) {
	token, err := jwt.Parse(prompt, func(token *jwt.Token) (interface{}, error) {
		return []byte(hashKey), nil
	})
	if err != nil {
		return nil, err
	}
	res := &UserJWT{}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//log.Println(claims)
		// Извлечение данных из токена
		res.ExpTime = time.Unix(int64(claims["exp"].(float64)), 0)
		res.ID = int64(claims["id"].(float64))
		res.Email = claims["email"].(string)
		res.Role = claims["role"].(string)

	} else {
		return nil, errInvalidToken
	}
	return res, nil
}

func ValidateViaJWT(access string, refresh string) (*UserJWT, error) {
	// получаем 2 JWT, пишем ебаное условие обновления на Access
	a, err := GetUserDataJWT(access)
	if err != nil {
		return nil, err
	}
	r, err := GetUserDataJWT(refresh)
	if err != nil {
		return nil, err
	}
	curTime := time.Now().Unix()
	if curTime > a.ExpTime.Unix() {
		if curTime > r.ExpTime.Unix() {
			return nil, errExpiredToken
		} else {
			a.ExpTime = time.Now().Add(time.Hour * 24)

		}
	}
	return a, nil
}

func (u *UserJWT) GenJWT() (string, error) {
	prompt := jwt.New(jwt.SigningMethodHS256)
	claims := prompt.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["email"] = u.Email
	claims["role"] = u.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Устанавливаем срок действия токена
	token, err := prompt.SignedString([]byte(hashKey))    // здесь блять ключ поменяй сука
	if err != nil {
		return "", err
	}
	log.Printf("[200] => CREATED AJWT FOR %s", u.Email)
	return token, nil
}
