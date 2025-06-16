package dao

func InsertNullOnZero(value interface{}) interface{} {
	if value == 0 {
		return nil
	}
	return value
}
