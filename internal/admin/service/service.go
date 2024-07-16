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
		dbRequest, err := c.Server.Store.User().Find(request.Id)
		if err != nil {
			return err
		}
		response := &structures.GetUserResponse{
			ID:        dbRequest.ID,
			Email:     dbRequest.Email,
			FirstName: dbRequest.FirstName,
			LastName:  dbRequest.LastName,
			Role:      dbRequest.Role,
			IsActive:  dbRequest.IsActive,
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
		users, err := c.Server.Store.User().GetList(request.Page, request.Limit)
		if err != nil { // Todo: эти структуры юзеров не отдавать(отдавать структуру которая в тудухе)
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		var response structures.GetListResponse
		for _, n := range users {
			user := &structures.GetUserListResponse{ID: n.ID, Email: n.Email}
			response.Table = append(response.Table, user)
		}
		return req.JSON(response)
	case "roles":
		roles, err := c.Server.Store.Role().GetList(request.Page, request.Limit)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return req.JSON(roles)

	case "cabinets":
		cabinets, err := c.Server.Store.Cabinet().GetList(request.Page, request.Limit)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return req.JSON(cabinets)

	case "groups":
		groups, err := c.Server.Store.Group().GetList(request.Page, request.Limit)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return req.JSON(groups)

	case "specializations":
		specializations, err := c.Server.Store.Specialization().GetList(request.Page, request.Limit)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return req.JSON(specializations)

	case "subjects":

		//TODO Доделать учителя
	//case "teachers":

	default:
		return fiber.NewError(fiber.StatusBadRequest, "Invalid table name")
	}
}

func (c *AdminPanelController) GetTables(req *fiber.Ctx) error {
	response := structures.GetTablesResponse{Tables: c.Server.Store.GetTables()}
	return req.JSON(response)
}
func (c *AdminPanelController) SetObj(req *fiber.Ctx) error {
	request := &structures.SetObj{}
	if err := req.BodyParser(request); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	switch request.TableName {
	case "users":
		if err := c.Server.Store.User().Update(request.Table); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		dbRequest, err := c.Server.Store.User().Find(request.Table.ID)
		if err != nil {
			return err
		}
		response := &structures.GetUserResponse{
			ID:        dbRequest.ID,
			Email:     dbRequest.Email,
			FirstName: dbRequest.FirstName,
			LastName:  dbRequest.LastName,
			Role:      dbRequest.Role,
			IsActive:  dbRequest.IsActive,
		}
		return req.JSON(response)
	default:
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}
}
