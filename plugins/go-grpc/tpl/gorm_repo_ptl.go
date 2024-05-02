package tpl

var GormRepo = `package repo

import (
	"fmt"
	"os"
	"{{ .ProjectName }}/db/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IRepo interface {
	Migrate()
	// put your database methods here
}

type Repo struct {
	*gorm.DB
}

func (r Repo) Migrate() {
	err := r.AutoMigrate(
		// put your models here
	)
	if err != nil {
		fmt.Printf("cant migrate db: %s", err)
		os.Exit(1)
	}
}

func New() IRepo {
	db := mustConnectDb()
	return &Repo{
		DB: db,
	}
}

func mustConnectDb() *gorm.DB {
	dsn := ""
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("cant connect to db: %s", err)
		os.Exit(1)
	}
	return db
}
`
