// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: company.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createCompany = `-- name: CreateCompany :one
INSERT INTO companies (
    id,
    owner_user_id,
    name,
    description,
    is_verified,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    $1, 
    $2, 
    $3, 
    false, 
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
)
RETURNING id, owner_user_id, name, description, is_verified, created_at, updated_at
`

type CreateCompanyParams struct {
	OwnerUserID pgtype.UUID
	Name        string
	Description pgtype.Text
}

func (q *Queries) CreateCompany(ctx context.Context, arg CreateCompanyParams) (Company, error) {
	row := q.db.QueryRow(ctx, createCompany, arg.OwnerUserID, arg.Name, arg.Description)
	var i Company
	err := row.Scan(
		&i.ID,
		&i.OwnerUserID,
		&i.Name,
		&i.Description,
		&i.IsVerified,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteCompany = `-- name: DeleteCompany :exec
DELETE FROM companies
WHERE id = $1
`

// TODO: Add + use auth to ensure only company owners can delete
func (q *Queries) DeleteCompany(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteCompany, id)
	return err
}

const getCompanyByID = `-- name: GetCompanyByID :one
SELECT id, owner_user_id, name, description, is_verified, created_at, updated_at
FROM companies
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCompanyByID(ctx context.Context, id pgtype.UUID) (Company, error) {
	row := q.db.QueryRow(ctx, getCompanyByID, id)
	var i Company
	err := row.Scan(
		&i.ID,
		&i.OwnerUserID,
		&i.Name,
		&i.Description,
		&i.IsVerified,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listCompanies = `-- name: ListCompanies :many
SELECT id, owner_user_id, name, description, is_verified, created_at, updated_at
FROM companies
ORDER BY updated_at DESC
`

func (q *Queries) ListCompanies(ctx context.Context) ([]Company, error) {
	rows, err := q.db.Query(ctx, listCompanies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Company
	for rows.Next() {
		var i Company
		if err := rows.Scan(
			&i.ID,
			&i.OwnerUserID,
			&i.Name,
			&i.Description,
			&i.IsVerified,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
