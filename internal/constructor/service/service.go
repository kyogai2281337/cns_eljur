package service

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	mongoDB "github.com/kyogai2281337/cns_eljur/pkg/mongo"
	"github.com/kyogai2281337/cns_eljur/pkg/mongo/structs"
	"github.com/kyogai2281337/cns_eljur/pkg/server"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/kyogai2281337/cns_eljur/internal/constructor/structures"
)

type ConstructorController struct {
	Server *server.Server
	Mongo  *mongo.Database
}

func NewConstructorController(s *server.Server, client *mongo.Client) *ConstructorController {
	db := client.Database("schedule")
	return &ConstructorController{
		Server: s,
		Mongo:  db,
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

	// Подключение к MongoDB
	client, ctx, cancel := mongoDB.ConnectMongoDB("mongodb://localhost:27017")
	defer cancel()
	defer client.Disconnect(ctx)

	db := client.Database("schedule")
	lecturesCollection := db.Collection("lectures")
	cabinetsCollection := db.Collection("cabinets")
	groupsCollection := db.Collection("groups")
	teachersCollection := db.Collection("teachers")
	subjectsCollection := db.Collection("subjects")

	// Поиск лекции
	var lecture *structs.Lecture
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

	// Получение данных для ответа
	var cabinet structs.Cabinet
	var group structs.Group
	var teacher structs.Teacher
	var subject structs.Subject

	cabinetsCollection.FindOne(ctx, bson.M{"_id": lecture.CabinetID}).Decode(&cabinet)
	groupsCollection.FindOne(ctx, bson.M{"_id": lecture.GroupID}).Decode(&group)
	teachersCollection.FindOne(ctx, bson.M{"_id": lecture.TeacherID}).Decode(&teacher)
	subjectsCollection.FindOne(ctx, bson.M{"_id": lecture.SubjectID}).Decode(&subject)

	// Формирование ответа
	response := structures.FindScheduleResponse{
		CabName:     strconv.Itoa(cabinet.Name),
		GroupName:   group.Name,
		TeachName:   teacher.Name,
		SubjectName: subject.Name,
		Place:       structures.Placement{request.Day, request.Pair},
	}

	return req.JSON(response)
}

// TODO: ЕБАНЫЙ В РОТ МЕТОДЫ СОЗДАНИЯ БЛЯТЬ СНАЧАЛА АТОМАРНЫЕ ДЛЯ ВНОСА ПОТОМ В ХИПЫ ОБЬЕДИНЯТЬ А ПОТОМ
// КРУД РАЗМЕРОМ С МОЙ ХУЙ ТАК ЕБАНЫЙ В РОТ ПОТОМ ПРЕОБРАЗОВЫВАТЬ НАДО АААААААААААААААААААААААААААААААА
