package apis

import (
	"net/http"
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
	StatusCode int                `json:"status_code"`
	Success    bool               `json:"success"`
	Message    string             `json:"message"`
	Data       interface{}        `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
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
	pagination := PaginationResponse{
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
