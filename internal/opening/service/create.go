package service

import (
	"github.com/jacksrm/gopportunitties/internal/opening/dto"
	"github.com/jacksrm/gopportunitties/internal/opening/entity"
)

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
