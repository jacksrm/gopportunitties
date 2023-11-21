package sqlite

import (
	"github.com/jacksrm/gopportunitties/config"
	"github.com/jacksrm/gopportunitties/internal/opening/entity"
	"gorm.io/gorm"
)

type SqliteOpeningRepository struct {
	db *gorm.DB
}

func New() SqliteOpeningRepository {
	return SqliteOpeningRepository{config.GetSqlite()}
}

func (s SqliteOpeningRepository) FindAll() ([]entity.Opening, error) {
	var openings []entity.Opening
	result := s.db.Find(&openings)

	if result.Error != nil {
		return nil, result.Error
	}

	return openings, nil
}

func (s SqliteOpeningRepository) FindById(id uint) (entity.Opening, error) {
	var opening entity.Opening
	result := s.db.First(&opening, id)

	if result.Error != nil {
		return entity.Opening{}, result.Error
	}

	return opening, nil
}

func (s SqliteOpeningRepository) Create(opening entity.Opening) (entity.Opening, error) {
	result := s.db.Create(&opening)

	if result.Error != nil {
		return entity.Opening{}, result.Error
	}

	return opening, nil
}

func (s SqliteOpeningRepository) Update(opening entity.Opening) (entity.Opening, error) {
	result := s.db.Save(&opening)

	if result.Error != nil {
		return entity.Opening{}, result.Error
	}

	return opening, nil
}

func (s SqliteOpeningRepository) Delete(opening entity.Opening) (entity.Opening, error) {
	result := s.db.Delete(&opening)

	if result.Error != nil {
		return entity.Opening{}, result.Error
	}

	return opening, nil
}
