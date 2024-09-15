package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kyogai2281337/cns_eljur/internal/admin/structures"
)

// GetList handles the get list request by parsing the request body, determining which table to search in,
// finding the entries with the provided pagination, and returning a JSON response with the entries' data.
//
// Parameters:
//   - req: a pointer to a fiber.Ctx object representing the HTTP request context.
//
// Returns:
//   - error: an error if there was an issue parsing the request, finding the entries, or returning the JSON response.
//     Otherwise, nil.
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
				ID: n.ID,
				//Specialization: n.Specialization,
				Name: n.Name,
				//MaxPairs:       n.MaxPairs,
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
