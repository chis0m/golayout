package db

import (
	"fmt"
	"github.com/pressly/goose/v3"
	"go-layout/config"
	"go-layout/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var AppDb *gorm.DB

func InitDB(env *config.Config) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		env.DBHost,
		env.DbUsername,
		env.DbPassword,
		env.DbName,
		env.DbPort,
	)
	if utils.IsLocal() {
		dsn += " sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	AppDb = db

	if err := runMigrations(); err != nil {
		return err
	}
	return nil
}

func runMigrations() error {
	dbSQL, err := AppDb.DB()
	if err != nil {
		return err
	}

	if err := goose.Up(dbSQL, "storage/db/migrations"); err != nil {
		log.Fatalf("Could not apply migrations: %v", err)
		return err
	}

	return nil
}
