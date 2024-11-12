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

type CompanyDocument struct {
	ID           pgtype.UUID
	CompanyID    pgtype.UUID
	DocumentType string
	FileUrl      string
	CreatedAt    pgtype.Timestamp
	UpdatedAt    pgtype.Timestamp
}

type CompanyFinancial struct {
	ID             pgtype.UUID
	CompanyID      pgtype.UUID
	FinancialYear  int32
	Revenue        pgtype.Numeric
	Expenses       pgtype.Numeric
	Profit         pgtype.Numeric
	Sales          pgtype.Numeric
	AmountRaised   pgtype.Numeric
	Arr            pgtype.Numeric
	GrantsReceived pgtype.Numeric
	CreatedAt      pgtype.Timestamp
	UpdatedAt      pgtype.Timestamp
}

type CompanyQuestionAnswer struct {
	ID         pgtype.UUID
	CompanyID  pgtype.UUID
	QuestionID pgtype.UUID
	AnswerText string
	CreatedAt  pgtype.Timestamp
	UpdatedAt  pgtype.Timestamp
}

type Employee struct {
	ID        pgtype.UUID
	CompanyID pgtype.UUID
	Name      string
	Email     string
	Role      string
	Bio       pgtype.Text
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Question struct {
	ID           pgtype.UUID
	QuestionText string
	CreatedAt    pgtype.Timestamp
	UpdatedAt    pgtype.Timestamp
}

type ResourceRequest struct {
	ID           pgtype.UUID
	CompanyID    pgtype.UUID
	ResourceType string
	Description  pgtype.Text
	Status       string
	CreatedAt    pgtype.Timestamp
	UpdatedAt    pgtype.Timestamp
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
