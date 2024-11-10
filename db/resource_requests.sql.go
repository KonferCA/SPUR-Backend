// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: resource_requests.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createResourceRequest = `-- name: CreateResourceRequest :one
INSERT INTO resource_requests (
    id,
    company_id,
    resource_type,
    description,
    status,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
)
RETURNING id, company_id, resource_type, description, status, created_at, updated_at
`

type CreateResourceRequestParams struct {
	CompanyID    pgtype.UUID
	ResourceType string
	Description  pgtype.Text
	Status       string
}

func (q *Queries) CreateResourceRequest(ctx context.Context, arg CreateResourceRequestParams) (ResourceRequest, error) {
	row := q.db.QueryRow(ctx, createResourceRequest,
		arg.CompanyID,
		arg.ResourceType,
		arg.Description,
		arg.Status,
	)
	var i ResourceRequest
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.ResourceType,
		&i.Description,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteResourceRequest = `-- name: DeleteResourceRequest :exec
DELETE FROM resource_requests
WHERE id = $1
`

func (q *Queries) DeleteResourceRequest(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteResourceRequest, id)
	return err
}

const getResourceRequestByID = `-- name: GetResourceRequestByID :one
SELECT id, company_id, resource_type, description, status, created_at, updated_at FROM resource_requests
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetResourceRequestByID(ctx context.Context, id pgtype.UUID) (ResourceRequest, error) {
	row := q.db.QueryRow(ctx, getResourceRequestByID, id)
	var i ResourceRequest
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.ResourceType,
		&i.Description,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listResourceRequests = `-- name: ListResourceRequests :many
SELECT id, company_id, resource_type, description, status, created_at, updated_at FROM resource_requests
ORDER BY updated_at DESC
`

func (q *Queries) ListResourceRequests(ctx context.Context) ([]ResourceRequest, error) {
	rows, err := q.db.Query(ctx, listResourceRequests)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ResourceRequest
	for rows.Next() {
		var i ResourceRequest
		if err := rows.Scan(
			&i.ID,
			&i.CompanyID,
			&i.ResourceType,
			&i.Description,
			&i.Status,
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

const listResourceRequestsByCompany = `-- name: ListResourceRequestsByCompany :many
SELECT id, company_id, resource_type, description, status, created_at, updated_at FROM resource_requests
WHERE company_id = $1
ORDER BY updated_at DESC
`

func (q *Queries) ListResourceRequestsByCompany(ctx context.Context, companyID pgtype.UUID) ([]ResourceRequest, error) {
	rows, err := q.db.Query(ctx, listResourceRequestsByCompany, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ResourceRequest
	for rows.Next() {
		var i ResourceRequest
		if err := rows.Scan(
			&i.ID,
			&i.CompanyID,
			&i.ResourceType,
			&i.Description,
			&i.Status,
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

const updateResourceRequestStatus = `-- name: UpdateResourceRequestStatus :one
UPDATE resource_requests
SET 
    status = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, company_id, resource_type, description, status, created_at, updated_at
`

type UpdateResourceRequestStatusParams struct {
	ID     pgtype.UUID
	Status string
}

func (q *Queries) UpdateResourceRequestStatus(ctx context.Context, arg UpdateResourceRequestStatusParams) (ResourceRequest, error) {
	row := q.db.QueryRow(ctx, updateResourceRequestStatus, arg.ID, arg.Status)
	var i ResourceRequest
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.ResourceType,
		&i.Description,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
