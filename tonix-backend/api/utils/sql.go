package utils

import "database/sql"

func NullableToString(s sql.NullString) *string {
	if !s.Valid {
		return nil
	}

	return &s.String
}
