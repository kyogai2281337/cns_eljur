package utils

import "github.com/kyogai2281337/cns_eljur/pkg/sql/model"

func UpdateGroups(current *model.Group, updated *model.Group) (string, []interface{}) {
	var values []interface{}
	query := "UPDATE groups SET"

	if current.Name != updated.Name {
		query += " name = ?,"
		values = append(values, updated.Name)
	}
	if current.Specialization.ID != updated.Specialization.ID {
		query += " spec_id = ?,"
		values = append(values, updated.Specialization.ID)
	}

	if current.MaxPairs != updated.MaxPairs {
		query += " max_pairs = ?,"
		values = append(values, updated.MaxPairs)
	}
	if len(values) != 0 {
		query = query[:len(query)-1] + " WHERE id = ?"
		values = append(values, updated.ID)
	}

	return query, values
}
