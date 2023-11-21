package service

import "github.com/jacksrm/gopportunitties/internal/opening/entity"

func (s Service) Delete(id uint64) (entity.Opening, error) {
	toDelete, err := s.repo.FindById(id)

	if err != nil {
		return entity.Opening{}, err
	}

	result, err := s.repo.Delete(toDelete)

	if err != nil {
		return entity.Opening{}, err
	}

	return result, nil
}
