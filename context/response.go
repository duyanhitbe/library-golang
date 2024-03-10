package context

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func Response(ctx *gin.Context, response *SuccessResponse) {
	ctx.JSON(response.StatusCode, response)
}

func CreatedResponse(data interface{}) *SuccessResponse {
	return &SuccessResponse{
		StatusCode: http.StatusCreated,
		Success:    true,
		Message:    SuccessMessage,
		Data:       data,
	}
}

func OkResponse(data interface{}) *SuccessResponse {
	return &SuccessResponse{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    SuccessMessage,
		Data:       data,
	}
}

func PaginatedResponse(req *ListRequest, data interface{}, total int64) *SuccessResponse {
	pagination := PaginationResponse{
		Limit: req.Limit,
		Page:  req.Page,
		Total: total,
	}
	return &SuccessResponse{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    SuccessMessage,
		Data:       data,
		Pagination: pagination,
	}
}
