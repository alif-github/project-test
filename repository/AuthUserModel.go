package repository

import "database/sql"

type AuthUserModel struct {
	ID        sql.NullInt64
	FullName  sql.NullString
	Username  sql.NullString
	Password  sql.NullString
	CreatedBy sql.NullInt64
	UpdatedBy sql.NullInt64
}
