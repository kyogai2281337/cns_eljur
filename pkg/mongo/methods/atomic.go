package inserts

import (
	"context"

	"github.com/kyogai2281337/cns_eljur/pkg/mongo/structs"
	"go.mongodb.org/mongo-driver/mongo"
)

func insertCabinet(collection *mongo.Collection, cabinet structs.Cabinet) (*mongo.InsertOneResult, error) {
	return collection.InsertOne(context.TODO(), cabinet)
}

func insertSubject(collection *mongo.Collection, subject structs.Subject) (*mongo.InsertOneResult, error) {
	return collection.InsertOne(context.TODO(), subject)
}

func insertSpecialization(collection *mongo.Collection, specialization structs.Specialization) (*mongo.InsertOneResult, error) {
	return collection.InsertOne(context.TODO(), specialization)
}

func insertGroup(collection *mongo.Collection, group structs.Group) (*mongo.InsertOneResult, error) {
	return collection.InsertOne(context.TODO(), group)
}

func insertTeacher(collection *mongo.Collection, teacher structs.Teacher) (*mongo.InsertOneResult, error) {
	return collection.InsertOne(context.TODO(), teacher)
}

func insertLecture(collection *mongo.Collection, lecture structs.Lecture) (*mongo.InsertOneResult, error) {
	return collection.InsertOne(context.TODO(), lecture)
}
