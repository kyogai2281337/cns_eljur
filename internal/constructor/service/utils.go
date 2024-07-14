package service

import (
	mongoDB "github.com/kyogai2281337/cns_eljur/pkg/mongo"
	"github.com/kyogai2281337/cns_eljur/pkg/set"
	"go.mongodb.org/mongo-driver/bson"
)

type _dump struct {
	list []any
}

func (d *_dump) toSet() (*set.Set, error) {

	s := set.New()
	for _, item := range d.list {
		if err := s.Push(item); err != nil {
			return nil, err
		}
	}
	return s, nil
}

func allObjects(url string, collname string) ([]any, error) {
	client, ctx, cancel := mongoDB.ConnectMongoDB(url)
	defer cancel()
	defer client.Disconnect(ctx)

	collection := client.Database("schedule").Collection(collname)

	// Извлечение всех документов из коллекции
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var response []interface{}
	if err = cursor.All(ctx, &response); err != nil {
		return nil, err
	}
	return response, nil
}
