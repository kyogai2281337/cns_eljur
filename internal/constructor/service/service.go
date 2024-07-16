package service

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	constructor "github.com/kyogai2281337/cns_eljur/internal/constructor/logic"
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

func (c *ConstructorController) Add(req *fiber.Ctx) error {
	var request structures.AddScheduleRequest
	if err := req.BodyParser(&request); err != nil {
		return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	client, ctx, cancel := mongoDB.ConnectMongoDB(c.MongoURL)
	defer cancel()
	defer client.Disconnect(ctx)

	db := client.Database("schedule")

	switch request.ObjType {
	case "cabinet":
		cabinetsCollection := db.Collection("cabinets")
		_, err := cabinetsCollection.InsertOne(ctx, request.Data)
		if err != nil {
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to insert data",
			})
		}
	case "group":
		groupsCollection := db.Collection("groups")
		_, err := groupsCollection.InsertOne(ctx, request.Data)
		if err != nil {
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to insert data",
			})
		}
	case "teacher":
		teachersCollection := db.Collection("teachers")
		_, err := teachersCollection.InsertOne(ctx, request.Data)
		if err != nil {
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to insert data",
			})
		}
	case "subject":
		subjectsCollection := db.Collection("subjects")
		_, err := subjectsCollection.InsertOne(ctx, request.Data)
		if err != nil {
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to insert data",
			})
		}
	case "specialization":
		specializationsCollection := db.Collection("specializations")
		_, err := specializationsCollection.InsertOne(ctx, request.Data)
		if err != nil {
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to insert data",
			})
		}
	default:
		return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid object type",
		})
	}

	return req.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
	})
}

func (c *ConstructorController) Delete(req *fiber.Ctx) error {
	var request structures.DelScheduleRequest
	if err := req.BodyParser(&request); err != nil {
		return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	client, ctx, cancel := mongoDB.ConnectMongoDB(c.MongoURL)
	defer cancel()
	defer client.Disconnect(ctx)

	db := client.Database("schedule")

	switch request.ObjType {
	case "cabinet":
		cabinetsCollection := db.Collection("cabinets")
		_, err := cabinetsCollection.DeleteOne(ctx, bson.M{"name": request.ObjID})
		if err != nil {
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to delete data",
			})
		}
	case "group":
		groupsCollection := db.Collection("groups")
		_, err := groupsCollection.DeleteOne(ctx, bson.M{"name": request.ObjID})
		if err != nil {
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to delete data",
			})
		}
	case "teacher":
		teachersCollection := db.Collection("teachers")
		_, err := teachersCollection.DeleteOne(ctx, bson.M{"id": request.ObjID})
		if err != nil {
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to delete data",
			})
		}
	case "subject":
		subjectsCollection := db.Collection("subjects")
		_, err := subjectsCollection.DeleteOne(ctx, bson.M{"id": request.ObjID})
		if err != nil {
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to delete data",
			})
		}
	case "specialization":
		specializationsCollection := db.Collection("specializations")
		_, err := specializationsCollection.DeleteOne(ctx, bson.M{"name": request.ObjID})
		if err != nil {
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to delete data",
			})
		}
	default:
		return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid object type",
		})
	}

	return req.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}

func (c *ConstructorController) Update(req *fiber.Ctx) error {
	var request structures.UpdateScheduleRequest
	if err := req.BodyParser(&request); err != nil {
		return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	client, ctx, cancel := mongoDB.ConnectMongoDB(c.MongoURL)
	defer cancel()
	defer client.Disconnect(ctx)

	db := client.Database("schedule")

	switch request.ObjType {
	case "cabinet":
		cabinetsCollection := db.Collection("cabinets")
		_, err := cabinetsCollection.UpdateOne(ctx, bson.M{"name": request.ObjID}, bson.M{"$set": request.Data})
		if err != nil {
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to update data",
			})
		}
	case "group":
		groupsCollection := db.Collection("groups")
		_, err := groupsCollection.UpdateOne(ctx, bson.M{"name": request.ObjID}, bson.M{"$set": request.Data})
		if err != nil {
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to update data",
			})
		}
	case "teacher":
		teachersCollection := db.Collection("teachers")
		_, err := teachersCollection.UpdateOne(ctx, bson.M{"id": request.ObjID}, bson.M{"$set": request.Data})
		if err != nil {
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to update data",
			})
		}
	case "subject":
		subjectsCollection := db.Collection("subjects")
		_, err := subjectsCollection.UpdateOne(ctx, bson.M{"id": request.ObjID}, bson.M{"$set": request.Data})
		if err != nil {
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to update data",
			})
		}
	case "specialization":
		specializationsCollection := db.Collection("specializations")
		_, err := specializationsCollection.UpdateOne(ctx, bson.M{"name": request.ObjID}, bson.M{"$set": request.Data})
		if err != nil {
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to update data",
			})
		}
	case "lecture":
		lecturesCollection := db.Collection("lectures")
		_, err := lecturesCollection.UpdateOne(ctx, bson.M{"_id": request.ObjID}, bson.M{"$set": request.Data})
		if err != nil {
			return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to update data",
			})
		}
	default:
		return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid object type",
		})
	}

	return req.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}

func (c *ConstructorController) Create(req *fiber.Ctx) error {
	var request structures.CreateScheduleRequest
	if err := req.BodyParser(&request); err != nil {
		return req.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}
	// инициализация этой залупы
	schedule := constructor.NewSchCab(request.Days, request.Pairs)
	gd := _dump{}
	gd.list, _ = allObjects(c.MongoURL, "groups")

	td := _dump{}
	td.list, _ = allObjects(c.MongoURL, "teachers")

	cd := _dump{}
	cd.list, _ = allObjects(c.MongoURL, "cabinets")

	// преобразование в член быка
	groups, err := gd.toSet()
	if err != nil {
		return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to convert data",
		})
	}
	teachers, err := td.toSet()
	if err != nil {
		return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to convert data",
		})
	}
	cabinets, err := cd.toSet()
	if err != nil {
		return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to convert data",
		})
	}

	schedule.AssignLecturesViaCabinet(groups, teachers, cabinets)

	dump := constructor.ToSimpleSchCabSorted(schedule)

	client, ctx, cancel := mongoDB.ConnectMongoDB(c.MongoURL)
	defer cancel()
	defer client.Disconnect(ctx)

	dumps := client.Database("schedule").Collection("dumps")
	_, err = dumps.InsertOne(ctx, dump)
	if err != nil {
		return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to insert data",
		})
	}

	return req.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}

func (c *ConstructorController) GetAllCabinets(req *fiber.Ctx) error {
	// Подключение к MongoDB
	client, ctx, cancel := mongoDB.ConnectMongoDB(c.MongoURL)
	defer cancel()
	defer client.Disconnect(ctx)

	collection := client.Database("schedule").Collection("cabinets")

	// Извлечение всех документов из коллекции
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to find documents",
		})
	}
	defer cursor.Close(ctx)

	var cabinets []structs.Cabinet
	if err = cursor.All(ctx, &cabinets); err != nil {
		return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to decode documents",
		})
	}

	return req.Status(fiber.StatusOK).JSON(cabinets)
}

// Аналогичные функции для других коллекций
func (c *ConstructorController) GetAllGroups(req *fiber.Ctx) error {
	client, ctx, cancel := mongoDB.ConnectMongoDB(c.MongoURL)
	defer cancel()
	defer client.Disconnect(ctx)

	collection := client.Database("schedule").Collection("groups")

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to find documents",
		})
	}
	defer cursor.Close(ctx)

	var groups []structs.Group
	if err = cursor.All(ctx, &groups); err != nil {
		return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to decode documents",
		})
	}

	return req.Status(fiber.StatusOK).JSON(groups)
}

func (c *ConstructorController) GetAllTeachers(req *fiber.Ctx) error {
	client, ctx, cancel := mongoDB.ConnectMongoDB(c.MongoURL)
	defer cancel()
	defer client.Disconnect(ctx)

	collection := client.Database("schedule").Collection("teachers")

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to find documents",
		})
	}
	defer cursor.Close(ctx)

	var teachers []structs.Teacher
	if err = cursor.All(ctx, &teachers); err != nil {
		return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to decode documents",
		})
	}

	return req.Status(fiber.StatusOK).JSON(teachers)
}

func (c *ConstructorController) GetAllSubjects(req *fiber.Ctx) error {
	client, ctx, cancel := mongoDB.ConnectMongoDB(c.MongoURL)
	defer cancel()
	defer client.Disconnect(ctx)

	collection := client.Database("schedule").Collection("subjects")

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to find documents",
		})
	}
	defer cursor.Close(ctx)

	var subjects []structs.Subject
	if err = cursor.All(ctx, &subjects); err != nil {
		return req.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to decode documents",
		})
	}

	return req.Status(fiber.StatusOK).JSON(subjects)
}
