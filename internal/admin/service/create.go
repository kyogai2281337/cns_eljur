package service

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/kyogai2281337/cns_eljur/internal/admin/structures"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
)

// Create handles the creation request by parsing the request body, determining which table to create an entry in,
// creating a new entry with the provided information, and returning a JSON response with the created entry's data.
//
// Parameters:
//   - req: a pointer to a fiber.Ctx object representing the HTTP request context.
//
// Returns:
//   - error: an error if there was an issue parsing the request, finding the user role, creating the user,
//     or returning the JSON response. Otherwise, nil.
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
			ID:       cabinetData.ID,
			Name:     cabinetData.Name,
			Type:     cabinetData.Type,
			Capacity: cabinetData.Capacity,
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
