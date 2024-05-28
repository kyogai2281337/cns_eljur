package apiserver

import (
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kyogai2281337/cns_eljur/internal/auth/model"
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

func toUserJWT(u *model.User) *UserJWT {
	return &UserJWT{
		u.ID,
		u.Email,
		u.Role.Name,
		time.Now(),
	}
}

func GenAuthJWT(u *model.User) (string, string, error) {
	ujwt := toUserJWT(u)
	access, err := ujwt.GenJWT("a")
	if err != nil {
		return "", "", err
	}
	refresh, err := ujwt.GenJWT("r")
	if err != nil {
		return "", "", err
	}
	log.Printf("[200] => CREATED AJWT AND RJWT FOR %s", u.Email)
	return access, refresh, nil
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
	if time.Now().Unix() > a.ExpTime.Unix() {
		if time.Now().Unix() > r.ExpTime.Unix() {
			return nil, errExpiredToken
		} else {
			a.ExpTime = time.Now().Add(time.Hour * 24)
		}
	}
	// upd: добавить проверку на наличие в базе + upd2 отладить код на перепрошив access через refresh
	return a, nil
}

func (u *UserJWT) GenJWT(t string) (string, error) {
	var exptime int64
	switch t {
	case "a":
		exptime = time.Now().Add(time.Hour * 24).Unix()
	case "r":
		exptime = time.Now().Add(time.Hour * (24 * 30)).Unix()
	default:
		exptime = time.Now().Add(time.Hour * 24).Unix()
	}

	prompt := jwt.New(jwt.SigningMethodHS256)
	claims := prompt.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["email"] = u.Email
	claims["role"] = u.Role
	claims["exp"] = exptime                            // Устанавливаем срок действия токена
	token, err := prompt.SignedString([]byte(hashKey)) // здесь блять ключ поменяй сука
	if err != nil {
		return "", err
	}
	log.Printf("[200] => CREATED %sJWT FOR %s", t, u.Email)
	return token, nil
}
