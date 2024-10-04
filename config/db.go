package config

import (
	"log"

	"github.com/muhammadsarimin/simple-api-xmu/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(con types.DBConfig) *gorm.DB {

	db, err := gorm.Open(postgres.Open(con.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if con.AutoMigrate {
		db.AutoMigrate(&types.Movie{})
	}

	return db

}
