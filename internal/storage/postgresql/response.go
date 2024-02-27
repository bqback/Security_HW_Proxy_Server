package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"proxy_server/internal/apperrors"
	"proxy_server/internal/pkg/dto"
	"proxy_server/internal/pkg/entities"
	"proxy_server/internal/utils"

	"github.com/Masterminds/squirrel"
	chimw "github.com/go-chi/chi/v5/middleware"
)

type PgResponseStorage struct {
	db *sql.DB
}

func NewResponseStorage(db *sql.DB) *PgResponseStorage {
	return &PgResponseStorage{
		db: db,
	}
}

func (s PgResponseStorage) StoreResponse(ctx context.Context, response *dto.IncomingResponse, reqID *dto.RequestID) error {
	funcName := "StoreResponse"
	logger := *utils.GetReqLogger(ctx)
	if logger == nil {
		return apperrors.ErrLoggerMissingFromContext
	}
	requestID := chimw.GetReqID(ctx)

	respQuery, args, err := squirrel.
		Insert("public.response").
		Columns(allResponseInsertFields...).
		Values(response.Code, response.Message, response.Headers, response.RawBody, response.TextBody).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID, funcName, nodeName)
		return apperrors.ErrCouldNotBuildQuery
	}

	result := dto.ResponseID{}

	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		logger.DebugFmt("Failed to start transaction with error "+err.Error(), requestID, funcName, nodeName)
		return apperrors.ErrCouldNotBeginTransaction
	}
	logger.DebugFmt("Transaction started", requestID, funcName, nodeName)

	query := tx.QueryRow(respQuery, args...)

	if err := query.Scan(&result.Value); err != nil {
		logger.DebugFmt("Storing response failed with error "+err.Error(), requestID, funcName, nodeName)
		err = tx.Rollback()
		if err != nil {
			logger.DebugFmt("Transaction rollback failed with error "+err.Error(), requestID, funcName, nodeName)
			return apperrors.ErrCouldNotRollback
		}
		return err
	}
	logger.DebugFmt("Response stored", requestID, funcName, nodeName)

	linkQuery, args, err := squirrel.
		Insert("public.request_response").
		Columns(requestResponseFields...).
		Values(reqID.Value, result.Value).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID, funcName, nodeName)
		return apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built req-resp query\n\t"+linkQuery+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID, funcName, nodeName)

	_, err = tx.Exec(linkQuery, args...)

	if err != nil {
		logger.DebugFmt("Linking response to request failed with error "+err.Error(), requestID, funcName, nodeName)
		err = tx.Rollback()
		if err != nil {
			logger.DebugFmt("Transaction rollback failed with error "+err.Error(), requestID, funcName, nodeName)
			return apperrors.ErrCouldNotRollback
		}
		return err
	}
	logger.DebugFmt("Response linked to request", requestID, funcName, nodeName)

	err = tx.Commit()
	if err != nil {
		logger.DebugFmt("Failed to commit changes with:"+err.Error(), requestID, funcName, nodeName)
		return apperrors.ErrCouldNotCommit
	}
	logger.DebugFmt("Changes commited", requestID, funcName, nodeName)

	return nil
}

func (s PgResponseStorage) GetResponseByRequestID(ctx context.Context, reqID *dto.RequestID) (*entities.Response, error) {
	// TODO Implement
	return nil, nil
}

func (s PgResponseStorage) GetResponseByResponseID(ctx context.Context, respID *dto.ResponseID) (*entities.Response, error) {
	// TODO Implement
	return nil, nil
}
