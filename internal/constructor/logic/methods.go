package constructor

import (
	"context"

	"github.com/kyogai2281337/cns_eljur/pkg/mongo/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func MakeLecture(g *Group, c *Cabinet, t *Teacher, s *Subject) *Lecture {
	l := &Lecture{
		Cabinet: c,
		Teacher: t,
		Group:   g,
		Subject: s,
	}
	return l
}

func saveSchCabSorted(collection *mongo.Collection, sch SchCabSorted, ctx context.Context) error {
	simpleSch := toSimpleSchCabSorted(sch)
	_, err := collection.InsertOne(ctx, simpleSch)
	return err
}

func loadSchCabSorted(collection *mongo.Collection, id primitive.ObjectID, cabinetsCollection, teachersCollection, groupsCollection, subjectsCollection *mongo.Collection, ctx context.Context) (SchCabSorted, error) {
	var simpleSch structs.SimpleSchCabSorted
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&simpleSch)
	if err != nil {
		return SchCabSorted{}, err
	}

	sch, err := toSchCabSorted(simpleSch, cabinetsCollection, teachersCollection, groupsCollection, subjectsCollection, ctx)
	if err != nil {
		return SchCabSorted{}, err
	}

	return sch, nil
}
