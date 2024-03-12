
-- name: CreateBorrower :one
INSERT INTO "borrowers" ("name", "phone", "address")
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListBorrower :many
SELECT * FROM "borrowers"
WHERE "deleted_at" IS NULL
LIMIT $1
OFFSET $2;

-- name: ListBorrowerByIds :many
SELECT * FROM "borrowers"
WHERE "id" = ANY(sqlc.arg(ids)::uuid[]) AND "deleted_at" IS NULL
LIMIT $1
OFFSET $2;

-- name: CountBorrower :one
SELECT COUNT(*) FROM "borrowers"
WHERE "deleted_at" IS NULL;

-- name: CountBorrowerByIds :one
SELECT COUNT(*) FROM "borrowers"
WHERE "id" = ANY(sqlc.arg(ids)::uuid[]) AND "deleted_at" IS NULL;

-- name: GetOneBorrowerById :one
SELECT * FROM "borrowers"
WHERE "id" = $1 AND "deleted_at" IS NULL;

-- name: GetOneBorrowerByPhone :one
SELECT * FROM "borrowers"
WHERE "phone" = $1 AND "deleted_at" IS NULL;

-- name: UpdateOneBorrowerById :one
UPDATE "borrowers"
SET "name" = $2, "phone" = $3, "address" = $4
WHERE "id" = $1 AND "deleted_at" IS NULL
RETURNING *;

-- name: DeleteOneBorrowerById :one
UPDATE "borrowers"
SET "deleted_at" = NOW()
WHERE "id" = $1 AND "deleted_at" IS NULL
RETURNING *;