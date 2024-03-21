package apis

import (
	"net/http"
	"time"

	"github.com/duyanhitbe/library-golang/db"
	"github.com/google/uuid"
)

const (
	SuccessMessage = "success"
)

type PaginatedResponse struct {
	Limit int64 `json:"limit,omitempty"`
	Page  int64 `json:"page,omitempty"`
	Total int64 `json:"total,omitempty"`
}

type PaginationResponse struct {
	StatusCode int                `json:"status_code,omitempty"`
	Success    bool               `json:"success,omitempty"`
	Message    string             `json:"message,omitempty"`
	Data       interface{}        `json:"data,omitempty"`
	Pagination *PaginatedResponse `json:"pagination,omitempty"`
}

type SuccessResponse struct {
	StatusCode int         `json:"status_code,omitempty"`
	Success    bool        `json:"success,omitempty"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func (server *HttpServer) Response(response *SuccessResponse) {
	server.ctx.JSON(response.StatusCode, response)
}

func (server *HttpServer) CreatedResponse(data interface{}) {
	server.Response(&SuccessResponse{
		StatusCode: http.StatusCreated,
		Success:    true,
		Message:    SuccessMessage,
		Data:       data,
	})
}

func (server *HttpServer) OkResponse(data interface{}) {
	server.Response(&SuccessResponse{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    SuccessMessage,
		Data:       data,
	})
}

func (server *HttpServer) PaginatedResponse(req *ListRequest, data interface{}, total int64) {
	pagination := &PaginatedResponse{
		Limit: req.Limit,
		Page:  req.Page,
		Total: total,
	}
	response := &PaginationResponse{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    SuccessMessage,
		Data:       data,
		Pagination: pagination,
	}
	server.ctx.JSON(response.StatusCode, response)
}

type BookResponse struct {
	Book     *db.Book     `json:"book,omitempty"`
	Category *db.Category `json:"category,omitempty"`
	BookInfo *db.BookInfo `json:"book_info,omitempty"`
}

func (server *HttpServer) parseBookResponse(book *db.Book) (*BookResponse, error) {
	bookInfo, err := server.store.GetOneBookInfoById(server.ctx, book.BookInfoID)
	if err != nil {
		return nil, err
	}
	category, err := server.store.GetOneCategoryById(server.ctx, book.CategoryID)
	if err != nil {
		return nil, err
	}
	rsp := BookResponse{
		Book:     book,
		Category: category,
		BookInfo: bookInfo,
	}
	return &rsp, nil
}

type UserResponse struct {
	ID        uuid.UUID   `json:"id,omitempty"`
	Username  string      `json:"username,omitempty"`
	Role      db.RoleEnum `json:"role,omitempty"`
	CreatedAt time.Time   `json:"created_at,omitempty"`
	UpdatedAt time.Time   `json:"updated_at,omitempty"`
	DeletedAt *time.Time  `json:"deleted_at,omitempty"`
	IsActive  bool        `json:"is_active,omitempty"`
}

func (server *HttpServer) parseUserResponse(user *db.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
		IsActive:  user.IsActive,
	}
}

type LoginUserResponse struct {
	AccessToken          string    `json:"access_token,omitempty"`
	AccessTokenExpiresAt time.Time `json:"access_token_exp,omitempty"`
	AccessTokenIssuedAt  time.Time `json:"access_token_iat,omitempty"`
}
