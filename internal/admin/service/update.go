package service

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/kyogai2281337/cns_eljur/internal/admin/structures"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
)

// SetObj handles the set object request by parsing the request body, determining which table to update in,
// updating the entry with the provided ID, and returning a JSON response with the updated entry's data.
//
// Parameters:
//   - req: a pointer to a fiber.Ctx object representing the HTTP request context.
//
// Returns:
//   - error: an error if there was an issue parsing the request, finding the user role, creating the user,
//     or returning the JSON response. Otherwise, nil.
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
		//fmt.Println(data)
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
