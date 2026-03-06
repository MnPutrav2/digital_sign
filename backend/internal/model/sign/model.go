package signModel

import (
	"time"

	"github.com/google/uuid"
)

type Sign struct {
	ID        uuid.UUID `json:"id" db:"id" structure:"UUID;primary key;default;gen_random_uuid()"`
	Signature string    `json:"signature" db:"signature" structure:"varchar(500);not null"`
	Timestamp time.Time `json:"timestamp" db:"timestamp" structure:"timestamp;not null;default;current_timestamp"`
}
