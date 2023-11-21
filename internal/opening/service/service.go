package service

import (
	"github.com/jacksrm/gopportunitties/internal/opening/repository"
)

type Service struct {
	repo repository.OpeningRepository
}

func New(repo repository.OpeningRepository) *Service {
	return &Service{repo}
}
