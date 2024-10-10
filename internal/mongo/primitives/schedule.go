package primitives

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"

	constructor "github.com/kyogai2281337/cns_eljur/internal/constructor/logic"
	mongostructures "github.com/kyogai2281337/cns_eljur/internal/mongo/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Schedule struct {
	Store *MongoStore
}

func (s *Schedule) Make(schedule *constructor.Schedule) error {
	mongoSchedule := mongostructures.ToMongoSchedule(schedule)
	schedulesCollection := s.Store.Client.Database("eljur").Collection("schedules")
	_, err := schedulesCollection.InsertOne(s.Store.Ctx, mongoSchedule)
	if err != nil {
		return err
	}
	return nil
}

func (s *Schedule) Update(id string, schedule *constructor.Schedule) error {
	mongoSchedule := mongostructures.ToMongoSchedule(schedule)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("error converting string to ObjectID: %s", err.Error())
	}
	schedulesCollection := s.Store.Client.Database("eljur").Collection("schedules")
	_, err = schedulesCollection.UpdateOne(s.Store.Ctx, bson.M{"_id": objectID}, bson.M{"$set": mongoSchedule})
	if err != nil {
		return err
	}
	return nil
}

func (s *Schedule) Find(id string) (*mongostructures.MongoSchedule, error) {
	schedulesCollection := s.Store.Client.Database("eljur").Collection("schedules")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("error converting string to ObjectID: %s", err.Error())
	}

	var mongoSchedule *mongostructures.MongoSchedule
	err = schedulesCollection.FindOne(s.Store.Ctx, bson.M{"_id": objectID}).Decode(&mongoSchedule)
	if err != nil {
		return nil, err
	}

	return mongoSchedule, nil
}

func (s *Schedule) GetList() (map[string]string, error) {
	schedulesCollection := s.Store.Client.Database("eljur").Collection("schedules")

	projection := bson.D{{Key: "_id", Value: 1}, {Key: "name", Value: 1}}
	cursor, err := schedulesCollection.Find(s.Store.Ctx, bson.D{}, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(s.Store.Ctx)
	response := make(map[string]string)
	var nameFiller int
	for cursor.Next(s.Store.Ctx) {
		var q bson.M
		err = cursor.Decode(&q)
		if err != nil {
			return nil, err
		}
		k, ok := q["name"].(string)
		if !ok {
			k = fmt.Sprintf("name_filled_%d", nameFiller)
			nameFiller++
		}
		v, ok := q["_id"].(primitive.ObjectID)
		if !ok {
			return nil, err
		}
		response[k] = v.Hex()
	}
	return response, nil
}

func (s *Schedule) Delete(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("error converting string to ObjectID: %s", err.Error())
	}
	schedulesCollection := s.Store.Client.Database("eljur").Collection("schedules")
	_, err = schedulesCollection.DeleteOne(s.Store.Ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	return nil
}

func (s *Schedule) Rename(id string, name string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("error converting string to ObjectID: %s", err.Error())
	}
	schedulesCollection := s.Store.Client.Database("eljur").Collection("schedules")
	_, err = schedulesCollection.UpdateOne(s.Store.Ctx, bson.M{"_id": objectID}, bson.M{"$set": bson.M{"name": name}})
	if err != nil {
		return err
	}
	return nil
}

func (s *Schedule) GetHash(id string) (string, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", fmt.Errorf("error converting string to ObjectID: %s", err.Error())
	}
	schedulesCollection := s.Store.Client.Database("eljur").Collection("schedules")
	var mongoSchedule *mongostructures.MongoSchedule
	err = schedulesCollection.FindOne(s.Store.Ctx, bson.M{"_id": objectID}).Decode(&mongoSchedule)
	if err != nil {
		return "", err
	}

	return hashStruct(mongoSchedule), nil
}

func hashStruct(data interface{}) string {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(buf.Bytes())
	return hex.EncodeToString(hash[:])
}
