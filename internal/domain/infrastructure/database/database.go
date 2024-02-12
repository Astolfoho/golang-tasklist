package database

import (
	"task-list/internal/domain/taskconfiguration"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=123mudar dbname=task-list port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("fail to connect database")
	}

	db.AutoMigrate(&taskconfiguration.TaskConfiguration{})

	return db
}
