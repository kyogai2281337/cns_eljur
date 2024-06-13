package adminPanel_service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kyogai2281337/cns_eljur/internal/adminPanel/adminPanel_structures"
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

func (c *AdminPanelController) GetObj(req *fiber.Ctx) error {
	request := &adminPanel_structures.GetObjRequest{}
	if err := req.BodyParser(request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	switch request.TName {
	case "users":
		Resp, err := c.Server.Store.User().Find(request.Id)
		Response := &adminPanel_structures.GetObjResponse{
			ID: Resp.ID,
		}
	}

	return req.JSON(response)
}

func (c *AdminPanelController) GetList(ctx *fiber.Ctx) error {
	return nil
}
