package database

import (
	"github.com/flpcastro/opportunities-api-go/internal/configs"
	"github.com/flpcastro/opportunities-api-go/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	logger := configs.GetLogger("postgre")

	conn := "host=localhost port=5432 user=root password=root dbname=apigolang sslmode=disable"
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		logger.Errorf("postgre error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&entity.Opportunity{})
	if err != nil {
		logger.Errorf("postgre automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
