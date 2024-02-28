package repository

import (
	"fmt"
	"go-fiber/config"
	"go-fiber/datamodel"
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(cfg *config.DBConfig) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.DBNAME,
		cfg.Port,
		cfg.SSL,
		cfg.TimeZone,
	)
	gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	err = db.AutoMigrate(&datamodel.User{})
	if err != nil {
		panic("Failed to migrate database")
	}

	slog.Info("Database connected successfully")
	DB = db
}
