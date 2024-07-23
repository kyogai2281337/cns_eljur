package service

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	constructor "github.com/kyogai2281337/cns_eljur/internal/constructor/logic"
	"github.com/kyogai2281337/cns_eljur/internal/constructor/structures"
	mongoDB "github.com/kyogai2281337/cns_eljur/internal/mongo"
	mongostructures "github.com/kyogai2281337/cns_eljur/internal/mongo/structs"
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
		user := ctx.Locals("user").(*model.User)
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

	groups := make([]*model.Group, 0)
	cabs := make([]*model.Cabinet, 0)
	teachers := make([]*model.Teacher, 0)
	plans := make([]*model.Specialization, 0)

	for _, groupID := range request.Groups {
		group, err := c.Server.Store.Group().Find(groupID)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		groups = append(groups, group)
	}

	for _, cabinetID := range request.Cabinets {
		cabinet, err := c.Server.Store.Cabinet().Find(cabinetID)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		cabs = append(cabs, cabinet)
	}

	for _, teacherID := range request.Teachers {
		teacher, err := c.Server.Store.Teacher().Find(teacherID)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		teachers = append(teachers, teacher)
	}

	for _, planID := range request.Plans {
		plan, err := c.Server.Store.Specialization().Find(planID)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		plans = append(plans, plan)
	}

	schedule := constructor.MakeSchedule(request.Limits.Days, request.Limits.Pairs, groups, teachers, cabs, plans, request.Limits.MaxWeeks, request.Limits.MaxDays)
	//TODO: make v correct
	err := schedule.Make()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	mongoSchedule := mongostructures.ToMongoSchedule(schedule)
	fmt.Println(schedule.String() + "\n\n\n\n\n")
	fmt.Println(mongoSchedule)
	client, dbCtx, cancel := mongoDB.ConnectMongoDB("")
	defer client.Disconnect(dbCtx)
	defer cancel()

	schedulesCollection := client.Database("eljur").Collection("schedules")

	_, err = schedulesCollection.InsertOne(dbCtx, mongoSchedule)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "ok"})
}
