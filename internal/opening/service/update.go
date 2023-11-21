package service

import (
	"github.com/jacksrm/gopportunitties/internal/opening/dto"
	"github.com/jacksrm/gopportunitties/internal/opening/entity"
)

func (s Service) Update(id uint64, data dto.UpdateOpening) (entity.Opening, error) {
	toUpdate, err := s.repo.FindById(id)

	if err != nil {
		return entity.Opening{}, err
	}

	if data.Role != nil {
		toUpdate.Role = *data.Role
	}

	if data.Company != nil {
		toUpdate.Company = *data.Company
	}

	if data.Location != nil {
		toUpdate.Location = *data.Location
	}

	if data.Remote != nil {
		toUpdate.Remote = *data.Remote
	}

	if data.Link != nil {
		toUpdate.Link = *data.Link
	}

	if data.Salary != nil {
		toUpdate.Salary = *data.Salary
	}

	result, err := s.repo.Update(toUpdate)

	if err != nil {
		return entity.Opening{}, err
	}

	return result, nil
}
