package postgresql

import (
	"database/sql"
	"proxy_server/internal/pkg/dto"
	"proxy_server/internal/pkg/entities"
)

type PgRequestStorage struct {
	db *sql.DB
}

func NewRequestStorage(db *sql.DB) *PgRequestStorage {
	return &PgRequestStorage{
		db: db,
	}
}

func (s PgRequestStorage) StoreRequest(request *dto.IncomingRequest) error {
	// TODO Implement
	return nil
}

func (s PgRequestStorage) GetRequestByID(requestID *dto.RequestID) (*entities.Request, error) {
	// TODO Implement
	return nil, nil
}

func (s PgRequestStorage) GetAllRequests() (*[]entities.Request, error) {
	// TODO Implement
	return nil, nil
}
