package service

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	constructor "github.com/kyogai2281337/cns_eljur/internal/constructor/logic"
	"github.com/kyogai2281337/cns_eljur/internal/constructor/structures"
	"github.com/kyogai2281337/cns_eljur/internal/mongo/primitives"
	"github.com/kyogai2281337/cns_eljur/pkg/server"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
)

type ConstructorController struct {
	Server *server.Server
}

func NewConstructorController(server *server.Server) *ConstructorController {
	return &ConstructorController{
		Server: server,
	}
}

func (c *ConstructorController) Authentication() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userData := ctx.Locals("user")
		if userData == nil {
			return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
		}

		user, ok := userData.(*model.User)
		if !ok {
			return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
		}

		if user.Role.Name != "superuser" {
			return fiber.NewError(fiber.StatusForbidden, "forbidden")
		}

		return ctx.Next()
	}
}

func (c *ConstructorController) Create(ctx *fiber.Ctx) error {
	request := &structures.CreateRequest{}
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	groups, cabs, teachers, plans, err := c.makeLists(request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	schedule := constructor.MakeSchedule(request.Name, request.Limits.Days, request.Limits.Pairs, groups, teachers, cabs, plans, request.Limits.MaxDays, request.Limits.MaxWeeks)
	schedule.Make()
	err = primitives.NewMongoConn().Schedule().Make(schedule)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "ok"})
}

func (c *ConstructorController) Get(ctx *fiber.Ctx) error {
	request := &structures.GetByIDRequest{}
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	mongoSchedule, err := primitives.NewMongoConn().Schedule().Find(request.ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok",
		"schedule": mongoSchedule})
}

func (c *ConstructorController) GetList(ctx *fiber.Ctx) error {
	response, err := primitives.NewMongoConn().Schedule().GetList()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok", "schedules": response})
}

func (c *ConstructorController) Update(ctx *fiber.Ctx) error {
	request := &structures.UpdateRequest{}
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	mongoSchedule, err := primitives.NewMongoConn().Schedule().Find(request.ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	schedule, err := c.RecoverToFull(mongoSchedule)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	switch request.OperationType {
	case "insert":
		data, err := json.Marshal(request.Value)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		var updData *structures.UpdateInsertRequest
		if err := json.Unmarshal(data, &updData); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		schedule.Insert(updData.Day, updData.Pair, schedule.RecoverLectureData(&struct {
			Group   string
			Teacher string
			Cabinet string
			Subject string
		}{
			Group:   updData.Lecture.Group,
			Teacher: updData.Lecture.Teacher,
			Cabinet: updData.Lecture.Cabinet,
			Subject: updData.Lecture.Subject,
		}))
		if err = primitives.NewMongoConn().Schedule().Update(request.ID, schedule); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok"})

	case "delete":
		data, err := json.Marshal(request.Value)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		var updData *structures.UpdateDeleteRequest
		if err := json.Unmarshal(data, &updData); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		schedule.Delete(updData.Day, updData.Pair, schedule.RecoverObject(updData.Name, updData.Type))

		if err = primitives.NewMongoConn().Schedule().Update(request.ID, schedule); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok"})
	default:
		return fiber.NewError(fiber.StatusInternalServerError, "unknown operation type")
	}
}

func (c *ConstructorController) Delete(ctx *fiber.Ctx) error {
	request := &structures.GetByIDRequest{}
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := primitives.NewMongoConn().Schedule().Delete(request.ID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok"})
}

func (c *ConstructorController) Rename(ctx *fiber.Ctx) error {
	request := &structures.RenameRequest{}
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := primitives.NewMongoConn().Schedule().Rename(request.ID, request.Name); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok"})
}
