package apis

import (
	"database/sql"
	"errors"
	"time"

	"github.com/duyanhitbe/library-golang/db"
	"github.com/duyanhitbe/library-golang/validations"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateBookRequest struct {
	Name            string    `json:"name" binding:"required"`
	CategoryID      uuid.UUID `json:"category_id" binding:"required"`
	Author          string    `json:"author" binding:"required"`
	PublicationDate time.Time `json:"publication_date" binding:"required"`
}

// CreateBook godoc
// @Summary Create one book
// @Tags Book API
// @Accept application/json
// @Produce application/json
// @Param body body apis.CreateBookRequest true "Create book request"
// @Security BearerAuth
// @Success 200 {object} apis.SuccessResponse{data=apis.BookResponse} "success"
// @Failure 400 {object} apis.ExceptionResponse "client error"
// @Failure 500 {object} apis.ExceptionResponse "database error"
// @Router /v1/book [post]
func (server *HttpServer) CreateBook(ctx *gin.Context) {
	var req CreateBookRequest
	if ok := server.BindJSON(&req); !ok {
		return
	}

	category, err := server.store.GetOneCategoryById(server.ctx, req.CategoryID)
	if err != nil {
		server.ThrowDbException(DbException{
			Err:             err,
			NotFoundMessage: "Category not found",
		})
		return
	}
	bookInfo, err := server.store.CreateBookInfo(ctx, db.CreateBookInfoParams{
		Name:            req.Name,
		Author:          req.Author,
		PublicationDate: req.PublicationDate,
	})
	if err != nil {
		server.ThrowDbException(DbException{
			Err:             err,
			NotFoundMessage: "Category not found",
		})
		return
	}

	book, err := server.store.CreateBook(ctx, db.CreateBookParams{
		CategoryID: req.CategoryID,
		BookInfoID: bookInfo.ID,
	})
	if err != nil {
		server.ThrowDbException(DbException{
			Err:             err,
			NotFoundMessage: "Category not found",
		})
		return
	}

	rsp := BookResponse{
		Book:     book,
		Category: category,
		BookInfo: bookInfo,
	}
	server.OkResponse(rsp)
}

// ListBook godoc
// @Summary Get a list of books
// @Tags Book API
// @Accept application/json
// @Produce application/json
// @Param query query apis.SwaggerListRequest false "List query request"
// @Success 200 {object} apis.PaginationResponse{data=[]apis.BookResponse} "success"
// @Failure 400 {object} apis.ExceptionResponse "client error"
// @Failure 500 {object} apis.ExceptionResponse "database error"
// @Router /v1/book [get]
func (server *HttpServer) ListBook(ctx *gin.Context) {
	req := server.BindPagination()
	if req == nil {
		return
	}

	params := db.ListBookParams{
		Limit:  int32(req.Limit),
		Offset: int32(req.Offset),
	}
	books, err := server.store.ListBook(ctx, params)
	if err != nil {
		server.ThrowDbException(DbException{
			Err: err,
		})
		return
	}
	total, err := server.store.CountBook(ctx)
	if err != nil {
		server.ThrowDbException(DbException{
			Err: err,
		})
		return
	}

	var result []*BookResponse

	for _, book := range books {
		rsp, err := server.parseBookResponse(book)
		if err != nil {
			server.ThrowDbException(DbException{
				Err: err,
			})
		}
		result = append(result, rsp)
	}

	if result == nil {
		result = []*BookResponse{}
	}

	server.PaginatedResponse(req, result, total)
}

// GetOneBookById godoc
// @Summary Get one book by id
// @Tags Book API
// @Accept application/json
// @Produce application/json
// @Param id path string true "book id"
// @Success 200 {object} apis.SuccessResponse{data=apis.BookResponse} "success"
// @Failure 400 {object} apis.ExceptionResponse "client error"
// @Failure 500 {object} apis.ExceptionResponse "database error"
// @Router /v1/book/{id} [get]
func (server *HttpServer) GetOneBookById(ctx *gin.Context) {
	id, ok := server.BindID()
	if !ok {
		return
	}

	book, err := server.store.GetOneBookById(ctx, *id)
	if err != nil {
		server.ThrowDbException(DbException{
			Err:             err,
			NotFoundMessage: "Book not found",
		})
		return
	}
	rsp, err := server.parseBookResponse(book)
	if err != nil {
		server.ThrowInternalServerException(err)
	}
	server.OkResponse(rsp)
}

type UpdateOneBookByIdRequest struct {
	Name            string    `json:"name" binding:"required"`
	CategoryID      uuid.UUID `json:"category_id" binding:"required"`
	Author          string    `json:"author" binding:"required"`
	PublicationDate time.Time `json:"publication_date" binding:"required"`
}

// UpdateOneBookById godoc
// @Summary Update one book by id
// @Tags Book API
// @Accept application/json
// @Produce application/json
// @Param id path string true "book id"
// @Param body body apis.UpdateOneBookByIdRequest true "Update book request"
// @Security BearerAuth
// @Success 200 {object} apis.SuccessResponse{data=apis.BookResponse} "success"
// @Failure 400 {object} apis.ExceptionResponse "client error"
// @Failure 500 {object} apis.ExceptionResponse "database error"
// @Router /v1/book/{id} [patch]
func (server *HttpServer) UpdateOneBookById(ctx *gin.Context) {
	var req UpdateOneBookByIdRequest
	if ok := server.BindJSON(&req); !ok {
		return
	}
	id, ok := server.BindID()
	if !ok {
		return
	}

	category, err := server.store.GetOneCategoryById(ctx, req.CategoryID)
	if err != nil {
		server.ThrowDbException(DbException{
			Err:             err,
			NotFoundMessage: "Category not found",
		})
		return
	}
	book, err := server.store.UpdateOneBookById(ctx, db.UpdateOneBookByIdParams{
		ID:         *id,
		CategoryID: req.CategoryID,
	})
	if err != nil {
		server.ThrowDbException(DbException{
			Err:             err,
			NotFoundMessage: "Book not found",
		})
		return
	}
	bookInfo, err := server.store.UpdateOneBookInfoById(ctx, db.UpdateOneBookInfoByIdParams{
		ID:              book.BookInfoID,
		Name:            req.Name,
		Author:          req.Author,
		PublicationDate: req.PublicationDate,
	})
	if err != nil {
		server.ThrowDbException(DbException{
			Err:             err,
			NotFoundMessage: "Book info not found",
		})
		return
	}

	rsp := BookResponse{
		Book:     book,
		Category: category,
		BookInfo: bookInfo,
	}

	server.OkResponse(rsp)
}

// DeleteOneBookById godoc
// @Summary Delete one book by id
// @Tags Book API
// @Accept application/json
// @Produce application/json
// @Param id path string true "book id"
// @Security BearerAuth
// @Success 200 {object} apis.SuccessResponse{data=apis.BookResponse} "success"
// @Failure 400 {object} apis.ExceptionResponse "client error"
// @Failure 500 {object} apis.ExceptionResponse "database error"
// @Router /v1/book/{id} [delete]
func (server *HttpServer) DeleteOneBookById(ctx *gin.Context) {
	id, ok := server.BindID()
	if !ok {
		return
	}

	book, err := server.store.DeleteOneBookById(ctx, *id)
	if err != nil {
		server.ThrowDbException(DbException{
			Err:             err,
			NotFoundMessage: "Book not found",
		})
		return
	}

	rsp, err := server.parseBookResponse(book)
	if err != nil {
		server.ThrowDbException(DbException{
			Err: err,
		})
	}
	server.OkResponse(rsp)
}

type BorrowBookRequest struct {
	BookID  string `json:"book_id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type BorrowBookResponse struct {
	*BookResponse
	Borrower *db.Borrower `json:"borrower"`
}

// BorrowBook godoc
// @Summary Borrow a book
// @Tags Book API
// @Accept application/json
// @Produce application/json
// @Param body body apis.BorrowBookRequest true "Borrow book request"
// @Security BearerAuth
// @Success 200 {object} apis.SuccessResponse{data=apis.BorrowBookResponse} "success"
// @Failure 400 {object} apis.ExceptionResponse "client error"
// @Failure 500 {object} apis.ExceptionResponse "database error"
// @Router /v1/book/borrow [post]
func (server *HttpServer) BorrowBook(ctx *gin.Context) {
	var req BorrowBookRequest
	if ok := server.BindJSON(&req); !ok {
		return
	}
	//Parse bookID to uuid
	bookID, err := validations.ParseUUID(req.BookID)
	if err != nil {
		server.ThrowBadRequestException(err)
	}
	//Get one book by id
	book, err := server.store.GetOneBookById(ctx, *bookID)
	if err != nil {
		server.ThrowDbException(DbException{
			Err:             err,
			NotFoundMessage: "Book not found",
		})
		return
	}
	//Get one or create new borrower
	//If exist borrower, update borrower info
	borrower, err := server.store.GetOneBorrowerByPhone(ctx, req.Phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			borrower, err = server.store.CreateBorrower(ctx, db.CreateBorrowerParams{
				Name:    req.Name,
				Phone:   req.Phone,
				Address: req.Address,
			})
			if err != nil {
				server.ThrowDbException(DbException{
					Err: err,
				})
				return
			}
		} else {
			server.ThrowDbException(DbException{
				Err: err,
			})
			return
		}
	} else {
		borrower, err = server.store.UpdateOneBorrowerById(ctx, db.UpdateOneBorrowerByIdParams{
			ID:      borrower.ID,
			Name:    req.Name,
			Phone:   req.Phone,
			Address: req.Address,
		})
		if err != nil {
			server.ThrowDbException(DbException{
				Err: err,
			})
			return
		}
	}

	bookBorrower, err := server.store.GetOneBookBorrower(ctx, db.GetOneBookBorrowerParams{
		BorrowerID: borrower.ID,
		BookID:     book.ID,
	})
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			server.ThrowDbException(DbException{
				Err: err,
			})
			return
		}
	}
	if bookBorrower == nil {
		_, err = server.store.CreateBookBorrower(ctx, db.CreateBookBorrowerParams{
			BorrowerID: borrower.ID,
			BookID:     book.ID,
		})
		if err != nil {
			server.ThrowDbException(DbException{
				Err: err,
			})
			return
		}
	}

	bookRsp, err := server.parseBookResponse(book)
	if err != nil {
		server.ThrowDbException(DbException{
			Err: err,
		})
		return
	}
	rsp := BorrowBookResponse{
		BookResponse: bookRsp,
		Borrower:     borrower,
	}
	server.OkResponse(rsp)
}

// ListBookByBorrowerId godoc
// @Summary Get list of books by borrowId
// @Tags Book API
// @Accept application/json
// @Produce application/json
// @Param id path string true "borrower id"
// @Param query query apis.SwaggerListRequest false "List query request"
// @Security BearerAuth
// @Success 200 {object} apis.PaginationResponse{data=apis.BookResponse} "success"
// @Failure 400 {object} apis.ExceptionResponse "client error"
// @Failure 500 {object} apis.ExceptionResponse "database error"
// @Router /v1/book/borrow/{id} [get]
func (server *HttpServer) ListBookByBorrowerId(ctx *gin.Context) {
	borrowerID, ok := server.BindID()
	if !ok {
		return
	}

	req := server.BindPagination()
	if req == nil {
		return
	}

	bookBorrowers, err := server.store.GetAllBookBorrowerByBorrowerId(ctx, *borrowerID)
	if err != nil {
		server.ThrowDbException(DbException{
			Err: err,
		})
		return
	}
	var bookIds []uuid.UUID
	for _, bookBorrower := range bookBorrowers {
		bookIds = append(bookIds, bookBorrower.BookID)
	}

	books, err := server.store.ListBookByIds(ctx, db.ListBookByIdsParams{
		Limit:  int32(req.Limit),
		Offset: int32(req.Offset),
		Ids:    bookIds,
	})
	if err != nil {
		server.ThrowDbException(DbException{
			Err: err,
		})
		return
	}

	var result []*BookResponse

	for _, book := range books {
		rsp, err := server.parseBookResponse(book)
		if err != nil {
			server.ThrowDbException(DbException{
				Err: err,
			})
			return
		}
		result = append(result, rsp)
	}

	total, err := server.store.CountBookByIds(ctx, bookIds)
	if err != nil {
		server.ThrowDbException(DbException{
			Err: err,
		})
		return
	}

	server.PaginatedResponse(req, result, total)
}
