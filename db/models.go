// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type RoleEnum string

const (
	RoleEnumADMIN   RoleEnum = "ADMIN"
	RoleEnumMANAGER RoleEnum = "MANAGER"
)

func (e *RoleEnum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = RoleEnum(s)
	case string:
		*e = RoleEnum(s)
	default:
		return fmt.Errorf("unsupported scan type for RoleEnum: %T", src)
	}
	return nil
}

type NullRoleEnum struct {
	RoleEnum RoleEnum `json:"role_enum"`
	Valid    bool     `json:"valid"` // Valid is true if RoleEnum is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRoleEnum) Scan(value interface{}) error {
	if value == nil {
		ns.RoleEnum, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.RoleEnum.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRoleEnum) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.RoleEnum), nil
}

type Book struct {
	ID         uuid.UUID  `json:"id"`
	CategoryID uuid.UUID  `json:"category_id"`
	BookInfoID uuid.UUID  `json:"book_info_id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
	IsActive   bool       `json:"is_active"`
}

type BookBorrower struct {
	BorrowerID uuid.UUID `json:"borrower_id"`
	BookID     uuid.UUID `json:"book_id"`
}

type BookInfo struct {
	ID              uuid.UUID  `json:"id"`
	Name            string     `json:"name"`
	Author          string     `json:"author"`
	PublicationDate time.Time  `json:"publication_date"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
	IsActive        bool       `json:"is_active"`
}

type Borrower struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Phone     string     `json:"phone"`
	Address   string     `json:"address"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	IsActive  bool       `json:"is_active"`
}

type Category struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	IsActive  bool       `json:"is_active"`
}

type User struct {
	ID        uuid.UUID  `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	Role      RoleEnum   `json:"role"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	IsActive  bool       `json:"is_active"`
}
