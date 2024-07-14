package constructor

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/kyogai2281337/cns_eljur/pkg/mongo/structs"
)

func ToSimpleSchCabSorted(sch *SchCabSorted) structs.SimpleSchCabSorted {
	simpleSch := structs.SimpleSchCabSorted{
		Days: make([][]map[string]structs.Lecture, len(sch.Days)),
	}

	for i, day := range sch.Days {
		simpleSch.Days[i] = make([]map[string]structs.Lecture, len(day))
		for j, pair := range day {
			simpleSch.Days[i][j] = make(map[string]structs.Lecture)
			for cab, lec := range pair {
				cabID := cab.ID.Hex()
				simpleSch.Days[i][j][cabID] = structs.Lecture{
					CabinetID: cab.ID,
					TeacherID: lec.Teacher.ID,
					GroupID:   lec.Group.ID,
					SubjectID: lec.Subject.ID,
				}
			}
		}
	}

	return simpleSch
}

func toSchCabSorted(simpleSch structs.SimpleSchCabSorted, cabinetsCollection, teachersCollection, groupsCollection, subjectsCollection *mongo.Collection, ctx context.Context) (SchCabSorted, error) {
	sch := SchCabSorted{
		Days: make([][]map[*Cabinet]*Lecture, len(simpleSch.Days)),
	}

	for i, day := range simpleSch.Days {
		sch.Days[i] = make([]map[*Cabinet]*Lecture, len(day))
		for j, pair := range day {
			sch.Days[i][j] = make(map[*Cabinet]*Lecture)
			for cabID, simpleLec := range pair {
				var cab Cabinet
				var teacher Teacher
				var group Group
				var subject Subject

				objID, _ := primitive.ObjectIDFromHex(cabID)

				cabinetsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&cab)
				teachersCollection.FindOne(ctx, bson.M{"_id": simpleLec.TeacherID}).Decode(&teacher)
				groupsCollection.FindOne(ctx, bson.M{"_id": simpleLec.GroupID}).Decode(&group)
				subjectsCollection.FindOne(ctx, bson.M{"_id": simpleLec.SubjectID}).Decode(&subject)

				sch.Days[i][j][&cab] = &Lecture{
					Cabinet: &cab,
					Teacher: &teacher,
					Group:   &group,
					Subject: &subject,
				}
			}
		}
	}

	return sch, nil
}
