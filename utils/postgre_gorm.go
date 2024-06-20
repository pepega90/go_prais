package utils

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBPostgreGorm() *gorm.DB {
	dsn := "postgresql://prais:prais@localhost:5432/db_prais"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalf("error while connecting to database: %v", err.Error())
		return nil
	}
	return db
}
