package handler

import (
	"github.com/AykoSousa/go-opportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db *gorm.DB
)

func InitialzeHander() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}