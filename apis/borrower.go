package apis

import (
	"github.com/duyanhitbe/library-golang/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (server *HttpServer) ListBorrowerByBookId(ctx *gin.Context) {
	bookID, ok := server.BindID()
	if !ok {
		return
	}

	req := server.BindPagination()
	if req == nil {
		return
	}

	bookBorrowers, err := server.store.GetAllBookBorrowerByBookId(ctx, *bookID)
	if err != nil {
		server.ThrowDbException(DbException{
			Err: err,
		})
		return
	}
	borrowerIds := []uuid.UUID{}
	for _, bookBorrower := range bookBorrowers {
		borrowerIds = append(borrowerIds, bookBorrower.BorrowerID)
	}

	borrowers, err := server.store.ListBorrowerByIds(ctx, db.ListBorrowerByIdsParams{
		Limit:  int32(req.Limit),
		Offset: int32(req.Offset),
		Ids:    borrowerIds,
	})
	if err != nil {
		server.ThrowDbException(DbException{
			Err: err,
		})
		return
	}

	total, err := server.store.CountBorrowerByIds(ctx, borrowerIds)
	if err != nil {
		server.ThrowDbException(DbException{
			Err: err,
		})
		return
	}

	server.PaginatedResponse(req, borrowers, total)
}
