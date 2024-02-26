package postgresql

import (
	"context"
	"database/sql"
	"proxy_server/internal/apperrors"
	"proxy_server/internal/pkg/dto"
	"proxy_server/internal/pkg/entities"
	"proxy_server/internal/utils"

	"github.com/Masterminds/squirrel"
	chimw "github.com/go-chi/chi/v5/middleware"
)

type PgRequestStorage struct {
	db *sql.DB
}

func NewRequestStorage(db *sql.DB) *PgRequestStorage {
	return &PgRequestStorage{
		db: db,
	}
}

func (s PgRequestStorage) StoreRequest(ctx context.Context, request *dto.IncomingRequest) (*dto.RequestID, error) {
	funcName := "StoreRequest"
	logger := *utils.GetReqLogger(ctx)
	if logger == nil {
		return nil, apperrors.ErrLoggerMissingFromContext
	}
	requestID := chimw.GetReqID(ctx)

	sql, args, err := squirrel.
		Insert("public.request").
		Columns(allRequestInsertFields...).
		Values(request.Method, request.Scheme, request.Host, request.Path, request.GetParams, request.Headers, request.Cookies, request.PostParams).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID, funcName, nodeName)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	// logger.DebugFmt("Built query\n\t"+sql+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID, funcName, nodeName)

	result := dto.RequestID{}

	query := s.db.QueryRow(sql, args...)

	if err := query.Scan(&result.Value); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s PgRequestStorage) GetRequestByID(ctx context.Context, requestID *dto.RequestID) (*entities.Request, error) {
	// TODO Implement
	return nil, nil
}

func (s PgRequestStorage) GetAllRequests(ctx context.Context) (*[]entities.Request, error) {
	// TODO Implement
	return nil, nil
}
