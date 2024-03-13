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

type PaginationResponse struct {
	Limit int64 `json:"limit"`
	Page  int64 `json:"page"`
	Total int64 `json:"total"`
}

type SuccessResponse struct {
	StatusCode int                 `json:"status_code"`
	Success    bool                `json:"success"`
	Message    string              `json:"message"`
	Data       interface{}         `json:"data"`
	Pagination *PaginationResponse `json:"pagination"`
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
	pagination := &PaginationResponse{
		Limit: req.Limit,
		Page:  req.Page,
		Total: total,
	}
	server.Response(&SuccessResponse{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    SuccessMessage,
		Data:       data,
		Pagination: pagination,
	})
}

type BookResponse struct {
	Book     *db.Book     `json:"book"`
	Category *db.Category `json:"category"`
	BookInfo *db.BookInfo `json:"book_info"`
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
	ID        uuid.UUID   `json:"id"`
	Username  string      `json:"username"`
	Role      db.RoleEnum `json:"role"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	DeletedAt *time.Time  `json:"deleted_at"`
	IsActive  bool        `json:"is_active"`
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
