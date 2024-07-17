package service

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kyogai2281337/cns_eljur/internal/constructor/structures"
	mongoDB "github.com/kyogai2281337/cns_eljur/pkg/mongo"
	"github.com/kyogai2281337/cns_eljur/pkg/mongo/structs"
	"github.com/kyogai2281337/cns_eljur/pkg/server"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConstructorController struct {
	Server   *server.Server
	MongoURL string
}

func NewConstructorController(s *server.Server, link string) *ConstructorController {
	return &ConstructorController{
		Server:   s,
		MongoURL: link,
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

func (c *ConstructorController) Find(req *fiber.Ctx) error {
	var request structures.FindScheduleRequest
	if err := req.BodyParser(&request); err != nil {
		return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	client, ctx, cancel := mongoDB.ConnectMongoDB(c.MongoURL)
	defer cancel()
	defer client.Disconnect(ctx)

	db := client.Database("schedule")
	lecturesCollection := db.Collection("lectures")
	cabinetsCollection := db.Collection("cabinets")
	groupsCollection := db.Collection("groups")
	teachersCollection := db.Collection("teachers")
	subjectsCollection := db.Collection("subjects")

	var lecture structs.Lecture
	objID, err := primitive.ObjectIDFromHex(request.ObjID)
	if err != nil {
		return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid ObjectID",
		})
	}

	filter := bson.M{
		"day":  request.Day,
		"pair": request.Pair,
	}

	switch request.ObjType {
	case "cabinet":
		filter["cabinetId"] = objID
	case "group":
		filter["groupId"] = objID
	case "teacher":
		filter["teacherId"] = objID
	default:
		return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid object type",
		})
	}

	err = lecturesCollection.FindOne(ctx, filter).Decode(&lecture)
	if err != nil {
		return req.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "schedule not found",
		})
	}

	var cabinet structs.Cabinet
	var group structs.Group
	var teacher structs.Teacher
	var subject structs.Subject

	cabinetsCollection.FindOne(ctx, bson.M{"_id": lecture.CabinetID}).Decode(&cabinet)
	groupsCollection.FindOne(ctx, bson.M{"_id": lecture.GroupID}).Decode(&group)
	teachersCollection.FindOne(ctx, bson.M{"_id": lecture.TeacherID}).Decode(&teacher)
	subjectsCollection.FindOne(ctx, bson.M{"_id": lecture.SubjectID}).Decode(&subject)

	response := structures.FindScheduleResponse{
		CabName:     strconv.Itoa(cabinet.Name),
		GroupName:   group.Name,
		TeachName:   teacher.Name,
		SubjectName: subject.Name,
		Place:       structures.Placement{request.Day, request.Pair},
	}

	return req.JSON(response)
}

// func (c *ConstructorController) Create(req *fiber.Ctx) error {
// 	var request structures.CreateScheduleRequest
// 	if err := req.BodyParser(&request); err != nil {
// 		return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "cannot parse JSON",
// 		})
// 	}
// 	// инициализация этой залупы
// 	schedule := constructor.NewSchCab(request.Days, request.Pairs)
// 	gd := _dump{}
// 	gd.list, _ = allObjects(c.MongoURL, "groups")

// 	td := _dump{}
// 	td.list, _ = allObjects(c.MongoURL, "teachers")

// 	cd := _dump{}
// 	cd.list, _ = allObjects(c.MongoURL, "cabinets")

// 	// преобразование в член быка
// 	groups, err := gd.toSet()
// 	if err != nil {
// 		return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "failed to convert data",
// 		})
// 	}
// 	teachers, err := td.toSet()
// 	if err != nil {
// 		return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "failed to convert data",
// 		})
// 	}
// 	cabinets, err := cd.toSet()
// 	if err != nil {
// 		return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "failed to convert data",
// 		})
// 	}

// 	if err = schedule.AssignLecturesViaCabinet(groups, teachers, cabinets); err != nil {
// 		return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": groups.Set,
// 		})
// 	}

// 	dump := constructor.ToSimpleSchCabSorted(schedule)

// 	client, ctx, cancel := mongoDB.ConnectMongoDB(c.MongoURL)
// 	defer cancel()
// 	defer client.Disconnect(ctx)

// 	dumps := client.Database("schedule").Collection("dumps")
// 	_, err = dumps.InsertOne(ctx, dump)
// 	if err != nil {
// 		return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "failed to insert data",
// 		})
// 	}

// 	return req.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"message": "success",
// 	})
// }
