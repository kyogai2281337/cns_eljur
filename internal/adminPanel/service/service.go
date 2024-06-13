package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kyogai2281337/cns_eljur/internal/adminPanel/structures"
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

func (c *AdminPAnelController) GetObj(req *fiber.Ctx) error {
	request := &structures.GetObjRequest{}
	if err := req.BodyParser(request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	switch request.TName {
	case "users":
		Resp, err := c.Server.Store.User().Find(request.Id)
		Response := &structures.GetObjResponse{
			ID: Resp.ID,
		}
	}

	return req.JSON(response)
}
