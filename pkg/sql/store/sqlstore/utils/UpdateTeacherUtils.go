package utils

import (
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
)

func UpdateTeachers(current *model.Teacher, updated *model.Teacher) (string, []interface{}) {
	var values []interface{}
	query := "UPDATE teachers SET"

	if current.Name != updated.Name {
		query += " name = ?,"
		values = append(values, updated.Name)
	}
	if current.RecommendSchCap_ != updated.RecommendSchCap_ {
		query += " capacity = ?,"
		values = append(values, updated.RecommendSchCap_)
	}

	if len(values) != 0 {
		query = query[:len(query)-1] + " WHERE id = ?"
		values = append(values, updated.ID)
	}

	return query, values
}
