package handler

import (
	"github.com/flpcastro/opportunities-api-go/internal/configs"
	"github.com/flpcastro/opportunities-api-go/pkg/database"
	"gorm.io/gorm"
)

var (
	logger *configs.Logger
	db     *gorm.DB
)

func InitHandler() {
	logger = configs.GetLogger("handler")
	db = database.GetPostgreSQL()
}
