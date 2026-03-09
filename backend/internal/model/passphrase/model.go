package passphraseModel

import "github.com/google/uuid"

type Passphrase struct {
	ID   uuid.UUID `json:"id" db:"id" structure:"UUID;primary key;default;gen_random_uuid()"`
	Pass string    `json:"pass" db:"pass" structure:"varchar(255);not null"`
}
