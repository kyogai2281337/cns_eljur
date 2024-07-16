package service

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/kyogai2281337/cns_eljur/internal/admin/structures"
	"github.com/kyogai2281337/cns_eljur/pkg/server"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
	"github.com/sirupsen/logrus"
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
		logrus.WithError(err).Error("Failed to parse req Body")
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	logrus.WithFields(logrus.Fields{
		"tableName": request.TableName,
		"id":        request.Id,
	}).Info("GetObj request received")

	switch request.TableName {
	case "users":
		dbRequest, err := c.Server.Store.User().Find(request.Id)
		if err != nil {
			logrus.WithError(err).Error("Failed to find user")
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
		return fiber.NewError(fiber.StatusBadRequest, "Invalid table name")
	}
}

func (c *AdminPanelController) GetList(req *fiber.Ctx) error {
	request := &structures.GetListRequest{}
	if err := req.BodyParser(request); err != nil {
		logrus.WithError(err).Error("Failed to parse request body")
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	logrus.WithFields(logrus.Fields{
		"tableName": request.TableName,
		"page":      request.Page,
		"limit":     request.Limit,
	}).Info("GetList request received")

	Table := request.TableName
	switch Table {
	case "users":
		users, err := c.Server.Store.User().GetUserList(request.Page, request.Limit)
		if err != nil { // Todo: эти структуры юзеров не отдавать(отдавать структуру которая в тудухе)
			logrus.WithError(err).Error("Failed to get user list")
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		var response structures.GetListResponse
		for _, n := range users {
			user := &structures.GetUserListResponse{ID: n.ID, Email: n.Email}
			response.Table = append(response.Table, user)
		}
		return req.JSON(response)
	case "roles":
		roles, err := c.Server.Store.Role().GetRoleList(request.Page, request.Limit)
		if err != nil {
			logrus.WithError(err).Error("Failed to get role list")
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return req.JSON(roles)
	default:
		return fiber.NewError(fiber.StatusBadRequest, "Invalid table name")
	}
}

func (c *AdminPanelController) GetTables(req *fiber.Ctx) error {
	logrus.Info("GetTables request received")
	response := structures.GetTablesResponse{Tables: c.Server.Store.GetTables()}
	return req.JSON(response)
}
func (c *AdminPanelController) SetObj(req *fiber.Ctx) error {
	request := &structures.SetObj{}
	if err := req.BodyParser(request); err != nil {
		logrus.WithError(err).Error("Failed to parse request body")
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	logrus.WithFields(logrus.Fields{
		"tableName": request.TableName,
		"table":     request.Table,
	}).Info("SetObj request received")

	switch request.TableName {
	case "users":
		var user model.User
		tabledata, err := json.Marshal(request.Table)
		if err != nil {
			logrus.WithError(err).Error("Failed to marshal table data")
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		if err := json.Unmarshal(tabledata, &user); err != nil {
			logrus.WithError(err).Error("Failed to unmarshal table data into user object")
			return fiber.NewError(fiber.StatusBadRequest, "Invalid user object")
		}
		if err := c.Server.Store.User().UpdateUser(&user); err != nil {
			logrus.WithError(err).Error("Failed to update user")
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		dbRequest, err := c.Server.Store.User().Find(user.ID)
		if err != nil {
			logrus.WithError(err).Error("Failed to find updated user")
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
		logrus.WithFields(logrus.Fields{
			"userID":    response.ID,
			"userEmail": response.Email,
		}).Info("User updated successfully")
		return req.JSON(response)
	default:
		logrus.Error("Invalid table name")
		return fiber.NewError(fiber.StatusBadRequest, "Invalid table name")
	}
}
