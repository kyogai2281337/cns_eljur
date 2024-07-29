package service

import (
	"encoding/json"
	"fmt"

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
			EncPass:   dbRequest.EncPass,
		}
		return req.Status(fiber.StatusOK).JSON(response)

	case "roles":
		dbRequest, err := c.Server.Store.Role().Find(request.Id)
		if err != nil {
			return err
		}
		response := &structures.GetRoleResponse{
			ID:   dbRequest.ID,
			Name: dbRequest.Name,
		}
		return req.Status(fiber.StatusOK).JSON(response)

	case "cabinets":
		dbRequest, err := c.Server.Store.Cabinet().Find(request.Id)
		if err != nil {
			return err
		}
		response := &structures.GetCabinetResponse{
			ID:   dbRequest.ID,
			Name: dbRequest.Name,
			Type: dbRequest.Type,
		}
		return req.Status(fiber.StatusOK).JSON(response)

	case "groups":
		dbRequest, err := c.Server.Store.Group().Find(request.Id)
		if err != nil {
			return err
		}
		response := &structures.GetGroupResponse{
			ID:             dbRequest.ID,
			Name:           dbRequest.Name,
			Specialization: dbRequest.Specialization,
			MaxPairs:       dbRequest.MaxPairs,
		}
		return req.Status(fiber.StatusOK).JSON(response)

	case "specializations":
		dbRequest, err := c.Server.Store.Specialization().Find(request.Id)
		if err != nil {
			return err
		}
		response := &structures.GetSpecializationResponse{
			ID:     dbRequest.ID,
			Name:   dbRequest.Name,
			Course: dbRequest.Course,
			//EduPlan: dbRequest.EduPlan,
			PlanId:    dbRequest.PlanId,
			ShortPlan: dbRequest.ShortPlan,
		}
		return req.Status(fiber.StatusOK).JSON(response)

	case "subjects":
		dbRequest, err := c.Server.Store.Subject().Find(request.Id)
		if err != nil {
			return err
		}
		response := &structures.GetSubjectResponse{
			ID:               dbRequest.ID,
			Name:             dbRequest.Name,
			RecommendCabType: dbRequest.RecommendCabType,
		}
		return req.Status(fiber.StatusOK).JSON(response)

	case "teachers":
		dbRequest, err := c.Server.Store.Teacher().Find(request.Id)
		if err != nil {
			return fmt.Errorf("Failed to find person %s", err.Error())
		}
		response := &structures.GetTeacherResponse{
			ID:   dbRequest.ID,
			Name: dbRequest.Name,
			//Links:            dbRequest.Links,
			LinksID:          dbRequest.LinksID,
			RecommendSchCap_: dbRequest.RecommendSchCap_,
			Sl:               dbRequest.SL,
		}
		return req.Status(fiber.StatusOK).JSON(response)

	default:
		return fiber.NewError(fiber.StatusBadRequest, "Invalid object name")
	}
}

func (c *AdminPanelController) GetList(req *fiber.Ctx) error {
	request := &structures.GetListRequest{}
	if err := req.BodyParser(request); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	tableName := request.TableName
	switch tableName {
	case "users":
		users, err := c.Server.Store.User().GetList(request.Page, request.Limit)
		if err != nil {
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		var response structures.GetListResponse
		for _, n := range users {
			user := &structures.GetUserListResponse{
				ID:    n.ID,
				Email: n.Email,
			}
			response.Table = append(response.Table, user)
		}
		return req.Status(fiber.StatusOK).JSON(response)

	case "roles":
		roles, err := c.Server.Store.Role().GetList(request.Page, request.Limit)
		if err != nil {
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		var response structures.GetListResponse
		for _, n := range roles {
			roleResponse := &structures.GetRoleResponse{
				ID:   n.ID,
				Name: n.Name,
			}
			response.Table = append(response.Table, roleResponse)
		}
		return req.Status(fiber.StatusOK).JSON(response)

	case "cabinets":
		cabinets, err := c.Server.Store.Cabinet().GetList(request.Page, request.Limit)
		if err != nil {
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		var response structures.GetListResponse
		for _, n := range cabinets {
			cabinetResponse := &structures.GetCabinetResponse{
				ID:   n.ID,
				Name: n.Name,
				Type: n.Type,
			}
			response.Table = append(response.Table, cabinetResponse)
		}
		return req.Status(fiber.StatusOK).JSON(response)

	case "groups":
		groups, err := c.Server.Store.Group().GetList(request.Page, request.Limit)
		if err != nil {
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		var response structures.GetListResponse
		for _, n := range groups {
			groupResponse := &structures.GetGroupResponse{
				ID:             n.ID,
				Specialization: n.Specialization,
				Name:           n.Name,
				MaxPairs:       n.MaxPairs,
			}
			response.Table = append(response.Table, groupResponse)
		}
		return req.Status(fiber.StatusOK).JSON(response)

	case "specializations":
		specializations, err := c.Server.Store.Specialization().GetList(request.Page, request.Limit)
		if err != nil {
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		var response structures.GetListResponse
		for _, n := range specializations {
			specializationsResponse := &structures.GetSpecializationResponse{
				ID:        n.ID,
				Name:      n.Name,
				Course:    n.Course,
				PlanId:    n.PlanId,
				ShortPlan: n.ShortPlan,
			}
			response.Table = append(response.Table, specializationsResponse)
		}
		return req.Status(fiber.StatusOK).JSON(response)

	case "subjects":
		subject, err := c.Server.Store.Subject().GetList(request.Page, request.Limit)
		if err != nil {
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		var response structures.GetListResponse
		for _, n := range subject {
			subjectsResponse := &structures.GetSubjectResponse{
				ID:               n.ID,
				Name:             n.Name,
				RecommendCabType: n.RecommendCabType,
			}
			response.Table = append(response.Table, subjectsResponse)
		}
		return req.Status(fiber.StatusOK).JSON(response)

	case "teachers":
		teachers, err := c.Server.Store.Teacher().GetList(request.Page, request.Limit)
		if err != nil {
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		var response structures.GetListResponse
		for _, n := range teachers {
			teachersResponse := &structures.GetTeacherResponse{
				ID:               n.ID,
				Name:             n.Name,
				RecommendSchCap_: n.RecommendSchCap_,
				LinksID:          n.LinksID,
				Sl:               n.SL,
			}
			response.Table = append(response.Table, teachersResponse)
		}
		return req.Status(fiber.StatusOK).JSON(response)

	default:
		return fiber.NewError(fiber.StatusBadRequest, "Invalid table name")
	}
}

func (c *AdminPanelController) GetTables(req *fiber.Ctx) error {
	response := structures.GetTablesResponse{Tables: c.Server.Store.GetTables()}
	return req.Status(fiber.StatusOK).JSON(response)
}
func (c *AdminPanelController) SetObj(req *fiber.Ctx) error {
	request := &structures.SetObj{}
	if err := req.BodyParser(request); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	switch request.TableName {
	case "users":

		// i love byte <3333
		// сериализуем
		data, err := json.Marshal(request.Table)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		var userData model.User

		// десериализуем
		if err := json.Unmarshal(data, &userData); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		dbCtx, err := c.Server.Store.BeginTx(req.Context())
		if err != nil {
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}
		defer c.Server.Store.RollbackTx(dbCtx)

		if err := c.Server.Store.User().Update(dbCtx, &userData); err != nil {
			c.Server.Store.RollbackTx(dbCtx)
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := c.Server.Store.CommitTx(dbCtx); err != nil {
			c.Server.Store.RollbackTx(dbCtx)
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{

				"error": err,
			})
		}
		response := &structures.GetUserResponse{
			ID:        userData.ID,
			Email:     userData.Email,
			FirstName: userData.FirstName,
			LastName:  userData.LastName,
			Role:      userData.Role,
			IsActive:  userData.IsActive,
			EncPass:   userData.EncPass,
		}
		return req.Status(fiber.StatusOK).JSON(response)

	case "cabinets":
		data, err := json.Marshal(request.Table)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		var cabinetsData model.Cabinet
		if err := json.Unmarshal(data, &cabinetsData); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		dbCtx, err := c.Server.Store.BeginTx(req.Context())
		if err != nil {

			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}
		defer c.Server.Store.RollbackTx(dbCtx)

		if err := c.Server.Store.Cabinet().Update(dbCtx, &cabinetsData); err != nil {
			c.Server.Store.RollbackTx(dbCtx)
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := c.Server.Store.CommitTx(dbCtx); err != nil {
			c.Server.Store.RollbackTx(dbCtx)
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}

		response := &structures.GetCabinetResponse{
			ID:   cabinetsData.ID,
			Name: cabinetsData.Name,
			Type: cabinetsData.Type,
		}
		return req.Status(fiber.StatusOK).JSON(response)

	case "groups":
		data, err := json.Marshal(request.Table)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		var groupsData model.Group
		if err := json.Unmarshal(data, &groupsData); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		dbCtx, err := c.Server.Store.BeginTx(req.Context())
		if err != nil {
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}
		defer c.Server.Store.RollbackTx(dbCtx)

		if err := c.Server.Store.Group().Update(dbCtx, &groupsData); err != nil {
			c.Server.Store.RollbackTx(dbCtx)
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := c.Server.Store.CommitTx(dbCtx); err != nil {
			c.Server.Store.RollbackTx(dbCtx)
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}

		response := &structures.GetGroupResponse{
			ID:             groupsData.ID,
			Specialization: groupsData.Specialization,
			Name:           groupsData.Name,
			MaxPairs:       groupsData.MaxPairs,
		}
		return req.Status(fiber.StatusOK).JSON(response)

	case "specializations":
		data, err := json.Marshal(request.Table)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		var specializationsData model.Specialization
		if err := json.Unmarshal(data, &specializationsData); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		dbCtx, err := c.Server.Store.BeginTx(req.Context())
		if err != nil {
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}
		defer c.Server.Store.RollbackTx(dbCtx)

		if err := c.Server.Store.Specialization().Update(dbCtx, &specializationsData); err != nil {
			c.Server.Store.RollbackTx(dbCtx)
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := c.Server.Store.CommitTx(dbCtx); err != nil {
			c.Server.Store.RollbackTx(dbCtx)
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}

		response := &structures.GetSpecializationResponse{

			ID:        specializationsData.ID,
			Name:      specializationsData.Name,
			Course:    specializationsData.Course,
			PlanId:    specializationsData.PlanId,
			ShortPlan: specializationsData.ShortPlan,
		}
		return req.Status(fiber.StatusOK).JSON(response)

	case "subjects":
		data, err := json.Marshal(request.Table)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		var subjectsData model.Subject
		if err := json.Unmarshal(data, &subjectsData); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		dbCtx, err := c.Server.Store.BeginTx(req.Context())
		if err != nil {
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}
		defer c.Server.Store.RollbackTx(dbCtx)

		if err := c.Server.Store.Subject().Update(dbCtx, &subjectsData); err != nil {
			c.Server.Store.RollbackTx(dbCtx)
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "update subjects failed",
			})
		}

		if err := c.Server.Store.CommitTx(dbCtx); err != nil {
			c.Server.Store.RollbackTx(dbCtx)
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}

		response := &structures.GetSubjectResponse{
			ID:               subjectsData.ID,
			Name:             subjectsData.Name,
			RecommendCabType: subjectsData.RecommendCabType,
		}
		return req.Status(fiber.StatusOK).JSON(response)

	case "teachers":
		data, err := json.Marshal(request.Table)
		fmt.Println(data)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		var teachersData model.Teacher
		if err := json.Unmarshal(data, &teachersData); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		dbCtx, err := c.Server.Store.BeginTx(req.Context())
		if err != nil {
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		defer c.Server.Store.RollbackTx(dbCtx)

		if err := c.Server.Store.Teacher().Update(dbCtx, &teachersData); err != nil {
			c.Server.Store.RollbackTx(dbCtx)
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := c.Server.Store.CommitTx(dbCtx); err != nil {
			c.Server.Store.RollbackTx(dbCtx)
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		response := &structures.GetTeacherResponse{
			ID:               teachersData.ID,
			Name:             teachersData.Name,
			RecommendSchCap_: teachersData.RecommendSchCap_,
			LinksID:          teachersData.LinksID,
			Sl:               teachersData.SL,
		}
		return req.Status(fiber.StatusOK).JSON(response)

	default:
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}
}

func (c *AdminPanelController) Create(req *fiber.Ctx) error {
	request := &structures.CreateRequest{}
	if err := req.BodyParser(request); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to parse request")
	}
	switch request.Table {
	case "cabinets":
		data, err := json.Marshal(request.Data)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		var cabinetData *model.Cabinet
		if err := json.Unmarshal(data, &cabinetData); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		dbCtx, err := c.Server.Store.BeginTx(req.Context())
		if err != nil {
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		defer c.Server.Store.RollbackTx(dbCtx)

		cabinetData, err = c.Server.Store.Cabinet().Create(dbCtx, cabinetData)
		if err != nil {
			c.Server.Store.RollbackTx(dbCtx)
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := c.Server.Store.CommitTx(dbCtx); err != nil {
			c.Server.Store.RollbackTx(dbCtx)
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}

		response := &structures.GetCabinetResponse{
			ID:   cabinetData.ID,
			Name: cabinetData.Name,
			Type: cabinetData.Type,
		}

		return req.Status(fiber.StatusOK).JSON(response)

	case "groups":
		data, err := json.Marshal(request.Data)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		var groupsData *model.Group
		if err := json.Unmarshal(data, &groupsData); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		dbCtx, err := c.Server.Store.BeginTx(req.Context())
		if err != nil {
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		defer c.Server.Store.RollbackTx(dbCtx)

		groupsData, err = c.Server.Store.Group().Create(dbCtx, groupsData)
		if err != nil {
			c.Server.Store.RollbackTx(dbCtx)
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := c.Server.Store.CommitTx(dbCtx); err != nil {
			c.Server.Store.RollbackTx(dbCtx)
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}

		response := &structures.GetGroupResponse{
			ID:             groupsData.ID,
			Specialization: groupsData.Specialization,
			Name:           groupsData.Name,
			MaxPairs:       groupsData.MaxPairs,
		}

		return req.Status(fiber.StatusOK).JSON(response)

	case "specializations":

		data, err := json.Marshal(request.Data)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		var specializationsData *model.Specialization
		if err := json.Unmarshal(data, &specializationsData); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		dbCtx, err := c.Server.Store.BeginTx(req.Context())
		if err != nil {
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		defer c.Server.Store.RollbackTx(dbCtx)

		specializationsData, err = c.Server.Store.Specialization().Create(dbCtx, specializationsData)
		if err != nil {
			c.Server.Store.RollbackTx(dbCtx)
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := c.Server.Store.CommitTx(dbCtx); err != nil {
			c.Server.Store.RollbackTx(dbCtx)
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}

		response := &structures.GetSpecializationResponse{
			ID:        specializationsData.ID,
			Name:      specializationsData.Name,
			Course:    specializationsData.Course,
			PlanId:    specializationsData.PlanId,
			ShortPlan: specializationsData.ShortPlan,
		}

		return req.Status(fiber.StatusOK).JSON(response)

	case "subjects":
		data, err := json.Marshal(request.Data)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		var subjectsData *model.Subject
		if err := json.Unmarshal(data, &subjectsData); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		dbCtx, err := c.Server.Store.BeginTx(req.Context())
		if err != nil {
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		defer c.Server.Store.RollbackTx(dbCtx)

		subjectsData, err = c.Server.Store.Subject().Create(dbCtx, subjectsData)
		if err != nil {
			c.Server.Store.RollbackTx(dbCtx)
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := c.Server.Store.CommitTx(dbCtx); err != nil {
			c.Server.Store.RollbackTx(dbCtx)
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}

		response := &structures.GetSubjectResponse{
			ID:               subjectsData.ID,
			Name:             subjectsData.Name,
			RecommendCabType: subjectsData.RecommendCabType,
		}

		return req.Status(fiber.StatusOK).JSON(response)

	case "teachers":
		data, err := json.Marshal(request.Data)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		var teachersData *model.Teacher
		if err := json.Unmarshal(data, &teachersData); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		dbCtx, err := c.Server.Store.BeginTx(req.Context())
		if err != nil {
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		defer c.Server.Store.RollbackTx(dbCtx)

		teachersData, err = c.Server.Store.Teacher().Create(dbCtx, teachersData)
		if err != nil {
			c.Server.Store.RollbackTx(dbCtx)
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := c.Server.Store.CommitTx(dbCtx); err != nil {
			c.Server.Store.RollbackTx(dbCtx)
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		response := &structures.GetTeacherResponse{
			ID:               teachersData.ID,
			Name:             teachersData.Name,
			LinksID:          teachersData.LinksID,
			RecommendSchCap_: teachersData.RecommendSchCap_,
			Sl:               teachersData.SL,
		}
		return req.Status(fiber.StatusOK).JSON(response)

	default:
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}
}
