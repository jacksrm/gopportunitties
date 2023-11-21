package service

import (
	"github.com/jacksrm/gopportunitties/internal/opening/dto"
	"github.com/jacksrm/gopportunitties/internal/opening/entity"
	"github.com/jacksrm/gopportunitties/internal/opening/repository"
)

type Service struct {
	repo repository.OpeningRepository
}

func New(repo repository.OpeningRepository) *Service {
	return &Service{repo}
}

func (s Service) Create(data dto.CreateOpening) (entity.Opening, error) {
	newOpening := entity.Opening{
		Role:     data.Role,
		Company:  data.Company,
		Location: data.Location,
		Remote:   *data.Remote,
		Link:     data.Link,
		Salary:   data.Salary,
	}

	result, err := s.repo.Create(newOpening)

	if err != nil {
		return entity.Opening{}, err
	}

	return result, nil
}

func (s Service) Get(id uint64) (entity.Opening, error) {
	opening, err := s.repo.FindById(id)

	if err != nil {
		return entity.Opening{}, err
	}

	return opening, nil
}

func (s Service) GetAll() ([]entity.Opening, error) {
	openings, err := s.repo.FindAll()

	if err != nil {
		return nil, err
	}

	return openings, nil
}

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
