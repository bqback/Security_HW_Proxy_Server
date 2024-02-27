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
		Values(request.Method, request.Scheme, request.Host, request.Path,
			request.GetParams, request.Headers, request.Cookies, request.PostParams, request.RawBody, request.TextBody).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID, funcName, nodeName)
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	result := dto.RequestID{}

	query := s.db.QueryRow(sql, args...)

	if err := query.Scan(&result.Value); err != nil {
		return nil, err
	}

	logger.DebugFmt("Request stored", requestID, funcName, nodeName)

	return &result, nil
}

func (s PgRequestStorage) GetRequestByID(ctx context.Context, requestID *dto.RequestID) (*entities.Request, error) {
	funcName := "GetRequestByID"
	logger := *utils.GetReqLogger(ctx)
	if logger == nil {
		return nil, apperrors.ErrLoggerMissingFromContext
	}
	reqID := chimw.GetReqID(ctx)

	sql, args, err := squirrel.Select(allRequestSelectFields...).
		From("public.request").
		Where(squirrel.Eq{"public.request.id": requestID.Value}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	row := s.db.QueryRow(sql, args...)

	request := entities.Request{}
	err = row.Scan(
		&request.ID,
		&request.Method,
		&request.Scheme,
		&request.Host,
		&request.Path,
		&request.GetParams,
		&request.Headers,
		&request.Cookies,
		&request.PostParams,
		&request.RawBody,
		&request.TextBody,
	)
	if err != nil {
		logger.Error("Parsing error: " + err.Error())
		return nil, apperrors.ErrCouldNotGetRequest
	}

	logger.DebugFmt("Collected request", reqID, funcName, nodeName)

	return &request, nil
}

func (s PgRequestStorage) GetAllRequests(ctx context.Context) (*[]entities.Request, error) {
	funcName := "GetAllRequests"
	logger := *utils.GetReqLogger(ctx)
	if logger == nil {
		return nil, apperrors.ErrLoggerMissingFromContext
	}
	requestID := chimw.GetReqID(ctx)

	sql, args, err := squirrel.Select(allRequestSelectFields...).
		From("public.request").
		ToSql()
	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	requestRows, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, apperrors.ErrCouldNotGetRequest
	}
	defer requestRows.Close()
	logger.DebugFmt("Got request info rows", requestID, funcName, nodeName)

	requests := []entities.Request{}
	for requestRows.Next() {
		var request entities.Request

		err = requestRows.Scan(
			&request.ID,
			&request.Method,
			&request.Scheme,
			&request.Host,
			&request.Path,
			&request.GetParams,
			&request.Headers,
			&request.Cookies,
			&request.PostParams,
			&request.RawBody,
			&request.TextBody,
		)
		if err != nil {
			logger.Error("Parsing error: " + err.Error())
			return nil, apperrors.ErrCouldNotGetRequest
		}
		requests = append(requests, request)
	}
	logger.DebugFmt("Collected requests", requestID, funcName, nodeName)

	return &requests, nil
}
