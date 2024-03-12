
-- name: CreateBookBorrower :one
INSERT INTO "book_borrower" ("borrower_id", "book_id")
VALUES ($1, $2)
RETURNING *;

-- name: GetOneBookBorrower :one
SELECT * FROM "book_borrower"
WHERE "borrower_id" = $1 AND "book_id" = $2;

-- name: GetAllBookBorrowerByBorrowerId :many
SELECT * FROM "book_borrower"
WHERE "borrower_id" = $1;

-- name: GetAllBookBorrowerByBookId :many
SELECT * FROM "book_borrower"
WHERE "book_id" = $1;