CREATE TYPE "role_enum" AS ENUM (
  'ADMIN',
  'MANAGER'
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "username" NVARCHAR NOT NULL,
  "password" NVARCHAR NOT NULL,
  "role" role_enum NOT NULL DEFAULT 'MANAGER',
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "deleted_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "name" NVARCHAR NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "deleted_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "books" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "category_id" uuid NOT NULL,
  "book_info_id" uuid NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "deleted_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "book_infos" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "name" NVARCHAR NOT NULL,
  "author" NVARCHAR NOT NULL,
  "publication_date" TIMESTAMPTZ NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "deleted_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "borrowers" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "name" NVARCHAR NOT NULL,
  "phone" NVARCHAR NOT NULL,
  "address" NVARCHAR NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "deleted_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "book_borrower" (
  "borrower_id" uuid NOT NULL,
  "book_id" uuid NOT NULL,
  PRIMARY KEY ("borrower_id", "book_id")
);

ALTER TABLE "books" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "books" ADD FOREIGN KEY ("book_info_id") REFERENCES "book_infos" ("id");

ALTER TABLE "book_borrower" ADD FOREIGN KEY ("borrower_id") REFERENCES "borrowers" ("id");

ALTER TABLE "book_borrower" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");
