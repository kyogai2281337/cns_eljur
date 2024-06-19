package utils

import "github.com/kyogai2281337/cns_eljur/pkg/sql/model"

func PrepareUpdateQueryAndValues(current *model.User, updated *model.User) (string, []interface{}) {
	var values []interface{}
	query := "UPDATE users SET"

	if current.Email != updated.Email {
		query += " email = ?,"
		values = append(values, updated.Email)
	}
	if current.EncPass != updated.EncPass {
		query += " encrypted_password = ?,"
		values = append(values, updated.EncPass)
	}
	if current.FirstName != updated.FirstName {
		query += " first_name = ?,"
		values = append(values, updated.FirstName)
	}
	if current.LastName != updated.LastName {
		query += " last_name = ?,"
		values = append(values, updated.LastName)
	}
	if current.Role.ID != updated.Role.ID {
		query += " role_id = ?,"
		values = append(values, updated.Role.ID)
	}
	if current.IsActive != updated.IsActive {
		query += " is_active = ?,"
		values = append(values, updated.IsActive)
	}
	if len(values) != 0 {
		query = query[:len(query)-1] + " WHERE id = ?"
		values = append(values, updated.ID)
	}

	return query, values
}
