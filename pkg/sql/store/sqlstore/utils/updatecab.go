package utils

import "github.com/kyogai2281337/cns_eljur/pkg/sql/model"

func UpdateCabs(current *model.Cabinet, updated *model.Cabinet) (string, []interface{}) {
	var values []interface{}
	query := "UPDATE cabinets SET"

	if current.Name != updated.Name {
		query += " name = ?,"
		values = append(values, updated.Name)
	}
	if current.Type != updated.Type {
		query += " type = ?,"
		values = append(values, updated.Type)
	}
	if len(values) != 0 {
		query = query[:len(query)-1] + " WHERE id = ?"
		values = append(values, updated.ID)
	}

	return query, values
}
