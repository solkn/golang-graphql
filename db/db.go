package database

import (
	"src/graph/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {

	connStr := "host=localhost user=postgres password=123 dbname=sol port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	dbConn, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	TableMigrations(dbConn)

	return dbConn, err
}

func TableMigrations(db *gorm.DB) {

	if !db.Migrator().HasTable(&model.BaseModel{}) {
		db.Migrator().CreateTable(&model.BaseModel{})
	}

	if !db.Migrator().HasTable(&model.Campaign{}) {
		db.Migrator().CreateTable(&model.Campaign{})
	}

	if !db.Migrator().HasTable(&model.Event{}) {
		db.Migrator().CreateTable(&model.Event{})
	}

	if !db.Migrator().HasTable(&model.Location{}) {
		db.Migrator().CreateTable(&model.Location{})
	}

	if !db.Migrator().HasTable(&model.Project{}) {
		db.Migrator().CreateTable(&model.Project{})
	}

	if !db.Migrator().HasTable(&model.Task{}) {
		db.Migrator().CreateTable(&model.Task{})
	}

	if !db.Migrator().HasTable(&model.TaskAssignment{}) {
		db.Migrator().CreateTable(&model.TaskAssignment{})

	}

	if !db.Migrator().HasTable(&model.User{}) {
		db.Migrator().CreateTable(&model.User{})
	}

}
