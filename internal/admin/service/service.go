package service

import (
	"encoding/json"

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
			PlanId: dbRequest.PlanId,
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
			return err
		}
		response := &structures.GetTeacherResponse{
			ID:   dbRequest.ID,
			Name: dbRequest.Name,
			//Links:            dbRequest.Links,
			LinksID:          dbRequest.LinksID,
			RecommendSchCap_: dbRequest.RecommendSchCap_,
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
				ID:     n.ID,
				Name:   n.Name,
				Course: n.Course,
				PlanId: n.PlanId,
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
			teachersResponse := &structures.GetSubjectResponse{
				ID:               n.ID,
				Name:             n.Name,
				RecommendCabType: model.CabType(n.RecommendSchCap_),
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

		if err := c.Server.Store.User().Update(&userData); err != nil {
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "update users failed",
			})
		}

		response := &structures.GetUserResponse{
			ID:        userData.ID,
			Email:     userData.Email,
			FirstName: userData.FirstName,
			LastName:  userData.LastName,
			Role:      userData.Role,
			IsActive:  userData.IsActive,
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
		if err := c.Server.Store.Cabinet().Update(&cabinetsData); err != nil {
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
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
		if err := c.Server.Store.Group().Update(&groupsData); err != nil {
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "update groups failed",
			})
		}
		response := &structures.GetGroupResponse{
			ID:             groupsData.ID,
			Specialization: groupsData.Specialization,
			Name:           groupsData.Name,
			MaxPairs:       groupsData.MaxPairs,
			//SpecPlan:       groupsData.SpecPlan,
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
		if err := c.Server.Store.Specialization().Update(&specializationsData); err != nil {
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "update specializations failed",
			})
		}
		response := &structures.GetSpecializationResponse{

			ID:     specializationsData.ID,
			Name:   specializationsData.Name,
			Course: specializationsData.Course,
			//EduPlan: specializationsData.EduPlan,
			PlanId: specializationsData.PlanId,
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
		if err := c.Server.Store.Subject().Update(&subjectsData); err != nil {
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "update subjects failed",
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
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		var teachersData model.Teacher
		if err := json.Unmarshal(data, &teachersData); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		if err := c.Server.Store.Teacher().Update(&teachersData); err != nil {
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "update teachers failed",
			})
		}
		response := &structures.GetTeacherResponse{
			ID:   teachersData.ID,
			Name: teachersData.Name,
			//Links:            teachersData.Links,
			LinksID:          teachersData.LinksID,
			RecommendSchCap_: teachersData.RecommendSchCap_,
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
		cabinetData, err = c.Server.Store.Cabinet().Create(cabinetData)
		if err != nil {
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "create cabinet failed",
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
		groupsData, err = c.Server.Store.Group().Create(groupsData)
		if err != nil {
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "create group failed",
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
		specializationsData, err = c.Server.Store.Specialization().Create(specializationsData)
		if err != nil {
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "create specialization failed",
			})
		}
		response := &structures.GetSpecializationResponse{
			ID:     specializationsData.ID,
			Name:   specializationsData.Name,
			Course: specializationsData.Course,
			PlanId: specializationsData.PlanId,
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
		subjectsData, err = c.Server.Store.Subject().Create(subjectsData)
		if err != nil {
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "create subject failed",
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
		teachersData, err = c.Server.Store.Teacher().Create(teachersData)
		if err != nil {
			return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "create teacher failed",
			})
		}
		response := &structures.GetTeacherResponse{
			ID:               teachersData.ID,
			Name:             teachersData.Name,
			LinksID:          teachersData.LinksID,
			RecommendSchCap_: teachersData.RecommendSchCap_,
		}
		return req.Status(fiber.StatusOK).JSON(response)

	default:
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}
}
