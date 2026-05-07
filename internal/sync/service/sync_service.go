package service

import (
	"context"

	"github.com/PPlaner/Backend/internal/dto"
	"github.com/PPlaner/Backend/internal/sync/repository"
)

type SyncService struct {
	repo *repository.SyncRepository
}

func NewSyncService(repo *repository.SyncRepository) *SyncService {
	return &SyncService{
		repo: repo,
	}
}

func (s *SyncService) Sync(ctx context.Context, userID int, req dto.SyncRequest) (dto.SyncResponse, error) {
	return s.repo.SyncData(ctx, userID, req)
}
