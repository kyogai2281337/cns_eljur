package adminPanel_service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kyogai2281337/cns_eljur/pkg/server"
)

type AdminPanelController struct {
	Server *server.Server
}

func NewAdminPanelController(s *server.Server) *AdminPanelController {
	return &AdminPanelController{
		Server: s,
	}
}

func (c *AdminPanelController) Authentication() fiber.Handler {
	return func(req *fiber.Ctx) error {
		cookie := req.Cookies("auth")
		if token, err := GetUserDataJWT(cookie); err != nil || token.Role != "superuser" {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}
		return req.Next()
	}
}

func (c *AdminPanelController) User(ctx *fiber.Ctx) error {

	return nil
}

func (c *AdminPanelController) Logout(ctx *fiber.Ctx) error {
	return nil
}

func (c *AdminPanelController) Delete(ctx *fiber.Ctx) error {
	return nil
}
