package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type CabType int

const (
	Normal CabType = iota
	Flowable
	Laboratory
	Computered
	Sport
)

type Cabinet struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name int                `bson:"name"`
	Type CabType            `bson:"type"`
}

type Subject struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	Name             string             `bson:"name"`
	RecommendCabType CabType            `bson:"recommendCabType"`
}

type Specialization struct {
	ID      primitive.ObjectID         `bson:"_id,omitempty"`
	Name    string                     `bson:"name"`
	Course  int                        `bson:"course"`
	EduPlan map[primitive.ObjectID]int `bson:"eduPlan"`
}

type Group struct {
	ID               primitive.ObjectID         `bson:"_id,omitempty"`
	SpecializationID primitive.ObjectID         `bson:"specializationId"`
	Name             string                     `bson:"name"`
	MaxPairs         int                        `bson:"maxPairs"`
	SpecPlan         map[primitive.ObjectID]int `bson:"specPlan"`
}

type Teacher struct {
	ID              primitive.ObjectID                          `bson:"_id,omitempty"`
	Name            string                                      `bson:"name"`
	Links           map[primitive.ObjectID][]primitive.ObjectID `bson:"links"`
	RecommendSchCap int                                         `bson:"recommendSchCap"`
}

type Lecture struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CabinetID primitive.ObjectID `bson:"cabinetId"`
	TeacherID primitive.ObjectID `bson:"teacherId"`
	GroupID   primitive.ObjectID `bson:"groupId"`
	SubjectID primitive.ObjectID `bson:"subjectId"`
}

type SimpleSchCabSorted struct {
	Days [][]map[string]Lecture `bson:"days"`
}
