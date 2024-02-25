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

func InitDB(env *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		env.Db.Host,
		env.Db.Username,
		env.Db.Password,
		env.Db.Name,
		env.Db.Port,
	)
	if utils.IsLocal() {
		dsn += " sslmode=disable"
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := runMigrations(db); err != nil {
		return nil, err
	}
	return db, nil
}

func runMigrations(db *gorm.DB) error {
	dbSQL, err := db.DB()
	if err != nil {
		return err
	}
	if err := goose.Up(dbSQL, "store/db/migrations"); err != nil {
		log.Fatalf("Could not run migrations: %v", err)
		return err
	}
	return nil
}
