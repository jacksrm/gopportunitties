package controller

import (
	"github.com/jacksrm/gopportunitties/config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() {
	db = config.GetSqlite()
}
