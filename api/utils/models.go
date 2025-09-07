package utils

import (
	"database/sql"

	"github.com/TDiblik/project-template/api/models"
)

func SQLNullStringFromString(value string) models.SQLNullString {
	return SQLNullStringFromStringRef(&value)
}

func SQLNullStringFromStringRef(value *string) models.SQLNullString {
	return models.SQLNullString{
		NullString: sql.NullString{
			String: DerefOrEmpty(value),
			Valid:  IsNotNil(value),
		},
	}
}
