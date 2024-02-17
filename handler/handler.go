package handler

import (
	"github.com/joaokogs/desafio-estagio/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQlite()
}
