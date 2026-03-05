package verifRepository

import (
	"database/sql"
)

type verifRepository struct {
	db *sql.DB
}

type VerifRepository interface {
	// write in here
}

func InitVerifRepository(db *sql.DB) VerifRepository {
	return &verifRepository{db: db}
}

// Entry
	