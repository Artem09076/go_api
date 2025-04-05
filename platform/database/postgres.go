package database

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Artem09076/go_api.git/pkg/utils"
)

var DB *gorm.DB

func PostgreSQLConnection() {
	var err error
	url := utils.ConnectionURLDatabase()

	DB, err = gorm.Open(postgres.Open(url))
	if err != nil {
		panic(err)
	}

}
