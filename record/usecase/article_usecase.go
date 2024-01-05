package usecase

import (
	"context"
	"time"

	"github.com/reinhardjs/mezink/domain"
	"github.com/reinhardjs/mezink/domain/dto/request"
)

type recordUsecase struct {
	recordRepo     domain.RecordRepository
	contextTimeout time.Duration
}

// NewRecordUsecase will create new an recordUsecase object representation of domain.RecordUsecase interface
func NewRecordUsecase(repo domain.RecordRepository, timeout time.Duration) domain.RecordUsecase {
	return &recordUsecase{
		recordRepo:     repo,
		contextTimeout: timeout,
	}
}

func (r *recordUsecase) GetBySumRange(c context.Context, request *request.GetRecordRequest) (res []domain.Record, err error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	res, err = r.recordRepo.GetBySumRange(ctx, request)

	return
}
