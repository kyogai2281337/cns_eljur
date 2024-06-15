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
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	switch request.TableName {
	case "users":
		resp, err := c.Server.Store.User().Find(request.Id)
		if err != nil {
			return err
		}
		response := &structures.GetUserResponse{
			ID:        resp.ID,
			Email:     resp.Email,
			FirstName: resp.FirstName,
			LastName:  resp.LastName,
			Role:      resp.Role,
			IsActive:  resp.IsActive,
			PermsSet:  resp.PermsSet,
		}
		return req.JSON(response)
	}

	return nil
}

func (c *AdminPanelController) GetList(req *fiber.Ctx) error {
	request := &structures.GetListRequest{}
	if err := req.BodyParser(request); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	Table := request.TableName
	switch Table {
	case "users":
		users, err := c.Server.Store.User().GetUserList(request.Page, request.Limit)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return req.JSON(users)
	case "roles":
		roles, err := c.Server.Store.Role().GetRoleList(request.Page, request.Limit)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return req.JSON(roles)
	default:
		return fiber.NewError(fiber.StatusBadRequest, "Invalid table name")
	}
}

func (c *AdminPanelController) GetTables(req *fiber.Ctx) error {
	return nil
}
func (c *AdminPanelController) SetObj(req *fiber.Ctx) error {
	return nil
}
