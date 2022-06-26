package database

import (
	"time"

	. "github.com/hotrungnhan/go-fiber-template/pkg/configs"
	"github.com/hotrungnhan/go-fiber-template/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgreSQLConnection func for connection to PostgreSQL database.
func PostgreSQLConnection() (*gorm.DB, error) {
	dsn, err := utils.ConnectionURLBuilder("postgres")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(Get().DB.MAX_IDLE_CONN)
	sqlDB.SetMaxOpenConns(Get().DB.MAX_CONN)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db, nil
}
