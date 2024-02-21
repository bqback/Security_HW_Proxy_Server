package postgresql

import (
	"database/sql"
	"proxy_server/internal/pkg/dto"
	"proxy_server/internal/pkg/entities"
)

type PgResponseStorage struct {
	db *sql.DB
}

func NewResponseStorage(db *sql.DB) *PgResponseStorage {
	return &PgResponseStorage{
		db: db,
	}
}

func (s PgResponseStorage) StoreResponse(*dto.IncomingResponse) error {
	// TODO Implement
	return nil
}

func (s PgResponseStorage) GetResponseByRequestID(*dto.RequestID) (*entities.Response, error) {
	// TODO Implement
	return nil, nil
}

func (s PgResponseStorage) GetResponseByResponseID(*dto.ResponseID) (*entities.Response, error) {
	// TODO Implement
	return nil, nil
}
