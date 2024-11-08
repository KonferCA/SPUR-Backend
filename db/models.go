// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Company struct {
	ID          pgtype.UUID
	OwnerUserID pgtype.UUID
	Name        string
	Description pgtype.Text
	IsVerified  bool
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type User struct {
	ID            pgtype.UUID
	Email         string
	PasswordHash  string
	FirstName     pgtype.Text
	LastName      pgtype.Text
	Role          string
	WalletAddress pgtype.Text
	CreatedAt     pgtype.Timestamp
	UpdatedAt     pgtype.Timestamp
}
