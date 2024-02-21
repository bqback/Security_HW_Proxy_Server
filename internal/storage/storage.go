package storage

import (
	"database/sql"
	"proxy_server/internal/storage/postgresql"
)

type Storages struct {
	Request  IRequestStorage
	Response IResponseStorage
}

func NewPostgresStorages(db *sql.DB) *Storages {
	return &Storages{
		Request:  postgresql.NewRequestStorage(db),
		Response: postgresql.NewResponseStorage(db),
	}
}
