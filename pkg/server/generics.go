package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store"
)

type Server struct {
	App   *fiber.App
	Store store.Store
}

func NewServer(store store.Store) *Server {
	s := &Server{
		App:   fiber.New(),
		Store: store,
	}

	s.App.Use("/private", s.Authentication())
	return s
}

func (s *Server) ServeHTTP(addr string) error {
	return s.App.Listen(addr)
}

func (s *Server) Authentication() fiber.Handler {
	return func(req *fiber.Ctx) error {
		cookie := req.Cookies("auth")
		jwt, err := GetUserDataJWT(cookie)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}

		user, err := s.Store.User().Find(jwt.ID)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}
		req.Locals("user", user)
		return req.Next()
	}
}
