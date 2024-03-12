// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: borrower.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const countBorrower = `-- name: CountBorrower :one
SELECT COUNT(*) FROM "borrowers"
WHERE "deleted_at" IS NULL
`

func (q *Queries) CountBorrower(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.countBorrowerStmt, countBorrower)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countBorrowerByIds = `-- name: CountBorrowerByIds :one
SELECT COUNT(*) FROM "borrowers"
WHERE "id" = ANY($1::uuid[]) AND "deleted_at" IS NULL
`

func (q *Queries) CountBorrowerByIds(ctx context.Context, ids []uuid.UUID) (int64, error) {
	row := q.queryRow(ctx, q.countBorrowerByIdsStmt, countBorrowerByIds, pq.Array(ids))
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createBorrower = `-- name: CreateBorrower :one
INSERT INTO "borrowers" ("name", "phone", "address")
VALUES ($1, $2, $3)
RETURNING id, name, phone, address, created_at, updated_at, deleted_at, is_active
`

type CreateBorrowerParams struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (q *Queries) CreateBorrower(ctx context.Context, arg CreateBorrowerParams) (*Borrower, error) {
	row := q.queryRow(ctx, q.createBorrowerStmt, createBorrower, arg.Name, arg.Phone, arg.Address)
	var i Borrower
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Phone,
		&i.Address,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.IsActive,
	)
	return &i, err
}

const deleteOneBorrowerById = `-- name: DeleteOneBorrowerById :one
UPDATE "borrowers"
SET "deleted_at" = NOW()
WHERE "id" = $1 AND "deleted_at" IS NULL
RETURNING id, name, phone, address, created_at, updated_at, deleted_at, is_active
`

func (q *Queries) DeleteOneBorrowerById(ctx context.Context, id uuid.UUID) (*Borrower, error) {
	row := q.queryRow(ctx, q.deleteOneBorrowerByIdStmt, deleteOneBorrowerById, id)
	var i Borrower
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Phone,
		&i.Address,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.IsActive,
	)
	return &i, err
}

const getOneBorrowerById = `-- name: GetOneBorrowerById :one
SELECT id, name, phone, address, created_at, updated_at, deleted_at, is_active FROM "borrowers"
WHERE "id" = $1 AND "deleted_at" IS NULL
`

func (q *Queries) GetOneBorrowerById(ctx context.Context, id uuid.UUID) (*Borrower, error) {
	row := q.queryRow(ctx, q.getOneBorrowerByIdStmt, getOneBorrowerById, id)
	var i Borrower
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Phone,
		&i.Address,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.IsActive,
	)
	return &i, err
}

const getOneBorrowerByPhone = `-- name: GetOneBorrowerByPhone :one
SELECT id, name, phone, address, created_at, updated_at, deleted_at, is_active FROM "borrowers"
WHERE "phone" = $1 AND "deleted_at" IS NULL
`

func (q *Queries) GetOneBorrowerByPhone(ctx context.Context, phone string) (*Borrower, error) {
	row := q.queryRow(ctx, q.getOneBorrowerByPhoneStmt, getOneBorrowerByPhone, phone)
	var i Borrower
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Phone,
		&i.Address,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.IsActive,
	)
	return &i, err
}

const listBorrower = `-- name: ListBorrower :many
SELECT id, name, phone, address, created_at, updated_at, deleted_at, is_active FROM "borrowers"
WHERE "deleted_at" IS NULL
LIMIT $1
OFFSET $2
`

type ListBorrowerParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListBorrower(ctx context.Context, arg ListBorrowerParams) ([]*Borrower, error) {
	rows, err := q.query(ctx, q.listBorrowerStmt, listBorrower, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Borrower{}
	for rows.Next() {
		var i Borrower
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Phone,
			&i.Address,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.IsActive,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listBorrowerByIds = `-- name: ListBorrowerByIds :many
SELECT id, name, phone, address, created_at, updated_at, deleted_at, is_active FROM "borrowers"
WHERE "id" = ANY($3::uuid[]) AND "deleted_at" IS NULL
LIMIT $1
OFFSET $2
`

type ListBorrowerByIdsParams struct {
	Limit  int32       `json:"limit"`
	Offset int32       `json:"offset"`
	Ids    []uuid.UUID `json:"ids"`
}

func (q *Queries) ListBorrowerByIds(ctx context.Context, arg ListBorrowerByIdsParams) ([]*Borrower, error) {
	rows, err := q.query(ctx, q.listBorrowerByIdsStmt, listBorrowerByIds, arg.Limit, arg.Offset, pq.Array(arg.Ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Borrower{}
	for rows.Next() {
		var i Borrower
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Phone,
			&i.Address,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.IsActive,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOneBorrowerById = `-- name: UpdateOneBorrowerById :one
UPDATE "borrowers"
SET "name" = $2, "phone" = $3, "address" = $4
WHERE "id" = $1 AND "deleted_at" IS NULL
RETURNING id, name, phone, address, created_at, updated_at, deleted_at, is_active
`

type UpdateOneBorrowerByIdParams struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Phone   string    `json:"phone"`
	Address string    `json:"address"`
}

func (q *Queries) UpdateOneBorrowerById(ctx context.Context, arg UpdateOneBorrowerByIdParams) (*Borrower, error) {
	row := q.queryRow(ctx, q.updateOneBorrowerByIdStmt, updateOneBorrowerById,
		arg.ID,
		arg.Name,
		arg.Phone,
		arg.Address,
	)
	var i Borrower
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Phone,
		&i.Address,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.IsActive,
	)
	return &i, err
}
