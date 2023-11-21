package service

import "github.com/jacksrm/gopportunitties/internal/opening/entity"

func (s Service) GetAll() ([]entity.Opening, error) {
	openings, err := s.repo.FindAll()

	if err != nil {
		return nil, err
	}

	return openings, nil
}
