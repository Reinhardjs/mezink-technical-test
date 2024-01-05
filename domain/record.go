package domain

import (
	"context"
	"time"

	"github.com/reinhardjs/mezink/domain/dto/request"
)

// Record is representing the Record data struct
type Record struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Marks     *[]int    `json:"marks" db:"-"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// RecordUsecase represent the record's usecases
type RecordUsecase interface {
	GetBySumRange(ctx context.Context, request *request.GetRecordRequest) ([]Record, error)
}

// RecordRepository represent the record's repository contract
type RecordRepository interface {
	GetBySumRange(ctx context.Context, request *request.GetRecordRequest) ([]Record, error)
}
