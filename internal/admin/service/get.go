package service

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kyogai2281337/cns_eljur/internal/admin/structures"
)

// GetObj handles the get object request by parsing the request body, determining which table to search in,
// finding the entry with the provided ID, and returning a JSON response with the entry's data.
//
// Parameters:
//   - req: a pointer to a fiber.Ctx object representing the HTTP request context.
//
// Returns:
//   - error: an error if there was an issue parsing the request, finding the user role, creating the user,
//     or returning the JSON response. Otherwise, nil.
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
			ID:       dbRequest.ID,
			Name:     dbRequest.Name,
			Type:     dbRequest.Type,
			Capacity: dbRequest.Capacity,
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
			return fmt.Errorf("failed to find person %s", err.Error())
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
