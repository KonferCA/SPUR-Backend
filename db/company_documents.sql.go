// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: company_documents.sql

package db

import (
	"context"
)

const createCompanyDocument = `-- name: CreateCompanyDocument :one
INSERT INTO company_documents (
    company_id,
    document_type,
    file_url
) VALUES (
    $1, $2, $3
)
RETURNING id, company_id, document_type, file_url, created_at, updated_at
`

type CreateCompanyDocumentParams struct {
	CompanyID    string
	DocumentType string
	FileUrl      string
}

func (q *Queries) CreateCompanyDocument(ctx context.Context, arg CreateCompanyDocumentParams) (CompanyDocument, error) {
	row := q.db.QueryRow(ctx, createCompanyDocument, arg.CompanyID, arg.DocumentType, arg.FileUrl)
	var i CompanyDocument
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.DocumentType,
		&i.FileUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteCompanyDocument = `-- name: DeleteCompanyDocument :exec
DELETE FROM company_documents
WHERE id = $1
`

func (q *Queries) DeleteCompanyDocument(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteCompanyDocument, id)
	return err
}

const getCompanyDocumentByID = `-- name: GetCompanyDocumentByID :one
SELECT id, company_id, document_type, file_url, created_at, updated_at FROM company_documents
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCompanyDocumentByID(ctx context.Context, id string) (CompanyDocument, error) {
	row := q.db.QueryRow(ctx, getCompanyDocumentByID, id)
	var i CompanyDocument
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.DocumentType,
		&i.FileUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listCompanyDocuments = `-- name: ListCompanyDocuments :many
SELECT id, company_id, document_type, file_url, created_at, updated_at FROM company_documents
WHERE company_id = $1
ORDER BY created_at DESC
`

func (q *Queries) ListCompanyDocuments(ctx context.Context, companyID string) ([]CompanyDocument, error) {
	rows, err := q.db.Query(ctx, listCompanyDocuments, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CompanyDocument
	for rows.Next() {
		var i CompanyDocument
		if err := rows.Scan(
			&i.ID,
			&i.CompanyID,
			&i.DocumentType,
			&i.FileUrl,
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

const listDocumentsByType = `-- name: ListDocumentsByType :many
SELECT id, company_id, document_type, file_url, created_at, updated_at FROM company_documents
WHERE company_id = $1 AND document_type = $2
ORDER BY created_at DESC
`

type ListDocumentsByTypeParams struct {
	CompanyID    string
	DocumentType string
}

func (q *Queries) ListDocumentsByType(ctx context.Context, arg ListDocumentsByTypeParams) ([]CompanyDocument, error) {
	rows, err := q.db.Query(ctx, listDocumentsByType, arg.CompanyID, arg.DocumentType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CompanyDocument
	for rows.Next() {
		var i CompanyDocument
		if err := rows.Scan(
			&i.ID,
			&i.CompanyID,
			&i.DocumentType,
			&i.FileUrl,
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

const updateCompanyDocument = `-- name: UpdateCompanyDocument :one
UPDATE company_documents
SET 
    document_type = $2,
    file_url = $3,
    updated_at = NOW()
WHERE id = $1
RETURNING id, company_id, document_type, file_url, created_at, updated_at
`

type UpdateCompanyDocumentParams struct {
	ID           string
	DocumentType string
	FileUrl      string
}

func (q *Queries) UpdateCompanyDocument(ctx context.Context, arg UpdateCompanyDocumentParams) (CompanyDocument, error) {
	row := q.db.QueryRow(ctx, updateCompanyDocument, arg.ID, arg.DocumentType, arg.FileUrl)
	var i CompanyDocument
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.DocumentType,
		&i.FileUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
