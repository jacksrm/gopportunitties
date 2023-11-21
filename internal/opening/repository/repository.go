package repository

import (
	"github.com/jacksrm/gopportunitties/internal/opening/entity"
)

type OpeningRepository interface {
	FindAll() ([]entity.Opening, error)
	FindById(id uint64) (entity.Opening, error)
	Create(opening entity.Opening) (entity.Opening, error)
	Update(opening entity.Opening) (entity.Opening, error)
	Delete(opening entity.Opening) (entity.Opening, error)
}
