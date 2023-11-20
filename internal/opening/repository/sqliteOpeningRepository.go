package repository

import "github.com/jacksrm/gopportunitties/internal/opening/entity"

type SqliteOpeningRepository struct{}

func (SqliteOpeningRepository) FindAll() ([]entity.Opening, error) {
	var openings []entity.Opening
	result := sqliteDB.Find(&openings)

	if result.Error != nil {
		return nil, result.Error
	}

	return openings, nil
}

func (SqliteOpeningRepository) FindById(id uint) (entity.Opening, error) {
	var opening entity.Opening
	result := sqliteDB.First(&opening, id)

	if result.Error != nil {
		return entity.Opening{}, result.Error
	}

	return opening, nil
}

func (SqliteOpeningRepository) Create(opening entity.Opening) (entity.Opening, error) {
	result := sqliteDB.Create(&opening)

	if result.Error != nil {
		return entity.Opening{}, result.Error
	}

	return opening, nil
}

func (SqliteOpeningRepository) Update(opening entity.Opening) (entity.Opening, error) {
	result := sqliteDB.Save(&opening)

	if result.Error != nil {
		return entity.Opening{}, result.Error
	}

	return opening, nil
}

func (SqliteOpeningRepository) Delete(opening entity.Opening) (entity.Opening, error) {
	result := sqliteDB.Delete(&opening)

	if result.Error != nil {
		return entity.Opening{}, result.Error
	}

	return opening, nil
}
