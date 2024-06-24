package types

import (
	"github.com/google/uuid"
	"github.com/notzree/uprank-backend/inference-backend/ent"
)

type UpworkRankingMessage struct {
	Message        string
	Job_id         uuid.UUID
	User_id        string
	Receipt_handle string
}

type UpworkRankingResult struct {
	//this should probably return some sort of array w/ the top rankings?
	Top_N []*ent.UpworkFreelancer
}
