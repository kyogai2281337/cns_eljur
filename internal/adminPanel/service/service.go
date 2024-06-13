package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kyogai2281337/cns_eljur/pkg/server"
)

type AdminPAnelController struct {
	Server *server.Server
}

func NewAdminPanelController(s *server.Server) *AdminPAnelController {
	return &AdminPAnelController{
		Server: s,
	}
}

func (c *AdminPAnelController) Authentication() fiber.Handler {
	return func(req *fiber.Ctx) error {
		cookie := req.Cookies("auth")
		if token, err := GetUserDataJWT(cookie); err != nil || token.Role != "superuser" {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}
		return req.Next()
	}
}
