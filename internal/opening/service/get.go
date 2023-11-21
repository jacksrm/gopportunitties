package service

import "github.com/jacksrm/gopportunitties/internal/opening/entity"

func (s Service) Get(id uint64) (entity.Opening, error) {
	opening, err := s.repo.FindById(id)

	if err != nil {
		return entity.Opening{}, err
	}

	return opening, nil
}
