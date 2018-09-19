package gostrutils

import "database/sql"

// ToSQLString convert string to sql.NullString
func ToSQLString(str string) sql.NullString {
	return sql.NullString{Valid: str != "", String: str}
}

// ToSQLStringNotNullable convert string to sql.NullString not nullable
func ToSQLStringNotNullable(str string) sql.NullString {
	return sql.NullString{Valid: true, String: str}
}
