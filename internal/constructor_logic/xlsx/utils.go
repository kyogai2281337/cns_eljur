package xlsx

func IntFieldExists(fields []int, field int) bool {
	for _, f := range fields {
		if f == field {
			return true
		}
	}
	return false
}

func StringFieldExists(fields []string, field string) bool {
	for _, f := range fields {
		if f == field {
			return true
		}
	}
	return false
}
