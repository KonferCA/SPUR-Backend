// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserRole string

const (
	UserRoleAdmin        UserRole = "admin"
	UserRoleStartupOwner UserRole = "startup_owner"
	UserRoleInvestor     UserRole = "investor"
)

func (e *UserRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserRole(s)
	case string:
		*e = UserRole(s)
	default:
		return fmt.Errorf("unsupported scan type for UserRole: %T", src)
	}
	return nil
}

type NullUserRole struct {
	UserRole UserRole
	Valid    bool // Valid is true if UserRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserRole) Scan(value interface{}) error {
	if value == nil {
		ns.UserRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserRole), nil
}

func (e UserRole) Valid() bool {
	switch e {
	case UserRoleAdmin,
		UserRoleStartupOwner,
		UserRoleInvestor:
		return true
	}
	return false
}

func AllUserRoleValues() []UserRole {
	return []UserRole{
		UserRoleAdmin,
		UserRoleStartupOwner,
		UserRoleInvestor,
	}
}

type Company struct {
	ID          string
	OwnerUserID string
	Name        string
	Description *string
	IsVerified  bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   pgtype.Timestamp
}

type CompanyDocument struct {
	ID           string
	CompanyID    string
	DocumentType string
	FileUrl      string
	CreatedAt    pgtype.Timestamp
	UpdatedAt    pgtype.Timestamp
}

type CompanyFinancial struct {
	ID             string
	CompanyID      string
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
	ID         string
	CompanyID  string
	QuestionID string
	AnswerText string
	CreatedAt  pgtype.Timestamp
	UpdatedAt  pgtype.Timestamp
	DeletedAt  pgtype.Timestamp
}

type Employee struct {
	ID        string
	CompanyID string
	Name      string
	Email     string
	Role      string
	Bio       *string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type FundingTransaction struct {
	ID                string
	ProjectID         string
	Amount            pgtype.Numeric
	Currency          string
	TransactionHash   string
	FromWalletAddress string
	ToWalletAddress   string
	Status            string
	CreatedAt         pgtype.Timestamp
	UpdatedAt         pgtype.Timestamp
}

type Meeting struct {
	ID                string
	ProjectID         string
	ScheduledByUserID string
	StartTime         time.Time
	EndTime           time.Time
	MeetingUrl        *string
	Location          *string
	Notes             *string
	CreatedAt         pgtype.Timestamp
	UpdatedAt         pgtype.Timestamp
}

type Project struct {
	ID          string
	CompanyID   string
	Title       string
	Description *string
	Status      string
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type ProjectComment struct {
	ID        string
	ProjectID string
	UserID    string
	Comment   string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type ProjectFile struct {
	ID        string
	ProjectID string
	FileType  string
	FileUrl   string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type ProjectLink struct {
	ID        string
	ProjectID string
	LinkType  string
	Url       string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type ProjectTag struct {
	ID        string
	ProjectID string
	TagID     string
	CreatedAt pgtype.Timestamp
}

type Question struct {
	ID           string
	QuestionText string
	CreatedAt    pgtype.Timestamp
	UpdatedAt    pgtype.Timestamp
	DeletedAt    pgtype.Timestamp
}

type ResourceRequest struct {
	ID           string
	CompanyID    string
	ResourceType string
	Description  *string
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Tag struct {
	ID        string
	Name      string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type User struct {
	ID            string
	Email         string
	PasswordHash  string
	FirstName     *string
	LastName      *string
	WalletAddress *string
	CreatedAt     pgtype.Timestamp
	UpdatedAt     pgtype.Timestamp
	Role          UserRole
}
