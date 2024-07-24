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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	//TODO: make v correct
	schedule.Normalize()
	fmt.Println(schedule)
	err = CreateMongoSchedule(schedule)
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
	client, dbCtx, cancel := mongoDB.ConnectMongoDB("")
	defer client.Disconnect(dbCtx)
	defer cancel()

	schedulesCollection := client.Database("eljur").Collection("schedules")
	id, err := primitive.ObjectIDFromHex(request.ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	var mongoSchedule mongostructures.MongoSchedule
	err = schedulesCollection.FindOne(dbCtx, bson.M{"_id": id}).Decode(&mongoSchedule)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok",
		"schedule": mongoSchedule})
}

func (c *ConstructorController) GetList(ctx *fiber.Ctx) error {
	client, dbCtx, cancel := mongoDB.ConnectMongoDB("")
	defer client.Disconnect(dbCtx)
	defer cancel()
	schedulesCollection := client.Database("eljur").Collection("schedules")

	projection := bson.D{{Key: "_id", Value: 1}, {Key: "name", Value: 1}}
	cursor, err := schedulesCollection.Find(dbCtx, bson.D{}, options.Find().SetProjection(projection))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer cursor.Close(dbCtx)
	response := make(map[string]string)
	var nameFiller int
	for cursor.Next(dbCtx) {
		var q bson.M
		err = cursor.Decode(&q)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		k, ok := q["name"].(string)
		if !ok {
			k = fmt.Sprintf("name_filled_%d", nameFiller)
			nameFiller++
		}
		v, ok := q["_id"].(primitive.ObjectID)
		if !ok {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot decode value _ID"})
		}
		response[k] = v.Hex()
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok", "schedules": response})
}
