package apis

import (
	"time"

	"github.com/duyanhitbe/library-golang/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateBookRequest struct {
	Name            string    `json:"name" binding:"required"`
	CategoryID      uuid.UUID `json:"category_id"`
	Author          string    `json:"author"`
	PublicationDate time.Time `json:"publication_date"`
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

	result := []*BookResponse{}

	for _, book := range books {
		rsp, err := server.parseBookResponse(book)
		if err != nil {
			server.ThrowDbException(DbException{
				Err: err,
			})
		}
		result = append(result, rsp)
	}

	server.PaginatedResponse(req, result, total)
}

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
	CategoryID      uuid.UUID `json:"category_id"`
	Author          string    `json:"author"`
	PublicationDate time.Time `json:"publication_date"`
}

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
