package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kyogai2281337/cns_eljur/internal/admin/structures"
	"github.com/kyogai2281337/cns_eljur/pkg/server"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
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
	return func(ctx *fiber.Ctx) error {
		user := ctx.Locals("user").(*model.User)
		if user.Role.Name != "superuser" {
			return fiber.NewError(fiber.StatusForbidden, "forbidden")
		}
		return ctx.Next()
	}
}

func (c *AdminPanelController) GetObj(req *fiber.Ctx) error {
	request := &structures.GetObjRequest{}
	if err := req.BodyParser(request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	switch request.TableName {
	case "user":
		Resp, err := c.Server.Store.User().Find(request.Id)
		if err != nil {
			return err
		}
		Response := &structures.GetUserResponse{
			ID:        Resp.ID,
			Email:     Resp.Email,
			FirstName: Resp.FirstName,
			LastName:  Resp.LastName,
			Role:      Resp.Role,
			IsActive:  Resp.IsActive,
			PermsSet:  Resp.PermsSet,
		}
		return req.JSON(Response)
	}

	return nil
}

func (c *AdminPanelController) GetList(req *fiber.Ctx) error {
	request := &structures.GetListRequest{}
	if err := req.BodyParser(request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	Table := request.TableName
	switch Table {
	case "users":
		resp, err := c.Server.Store.User().GetUserList(request.Page, request.Limit)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return req.JSON(resp)
	}
}
