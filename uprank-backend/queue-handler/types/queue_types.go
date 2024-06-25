package types

import (
	"github.com/google/uuid"
)

type UpworkRankingMessage struct {
	Message        string
	Job_id         uuid.UUID
	Platform       string
	Platform_id    string
	User_id        string
	Receipt_handle string
}
