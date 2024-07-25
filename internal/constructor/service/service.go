package service

import (
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
