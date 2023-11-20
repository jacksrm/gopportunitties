package repository

import (
	"github.com/jacksrm/gopportunitties/config"
	"github.com/jacksrm/gopportunitties/internal/opening/entity"
	"gorm.io/gorm"
)

var sqliteDB *gorm.DB = config.GetSqlite()

type OpeningRepository interface {
	FindAll() ([]entity.Opening, error)
	FindById(id uint) (entity.Opening, error)
	Create(opening entity.Opening) (entity.Opening, error)
	Update(opening entity.Opening) (entity.Opening, error)
	Delete(opening entity.Opening) (entity.Opening, error)
}
