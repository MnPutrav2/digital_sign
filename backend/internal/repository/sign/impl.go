package signRepository

import (
	"context"
	"database/sql"
	"time"

	signModel "github.com/MnPutrav2/go_architecture/internal/model/sign"
	"github.com/MnPutrav2/go_architecture/pkg/query"
	"github.com/google/uuid"
)

type signRepository struct {
	db *sql.DB
}

type SignRepository interface {
	InsertSignature(ctx context.Context, sign string) error
}

func InitSignRepository(db *sql.DB) SignRepository {
	return &signRepository{db: db}
}

// Entry
func (q *signRepository) InsertSignature(ctx context.Context, sign string) error {
	payload := signModel.Sign{
		ID:        uuid.New(),
		Signature: sign,
		Timestamp: time.Now(),
	}

	if err := query.Init[signModel.Sign](q.db).Insert(payload).Exec(ctx); err != nil {
		return err
	}

	return nil
}
