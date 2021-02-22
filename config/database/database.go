package database

import (
	"fmt"
	"log"

	"github.com/iamaul/fatbellies/config"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var dbConnection *gorm.DB

func ConnectDatabase(c *config.Configuration) (*gorm.DB, error) {
	dbargs := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", c.DbHost, c.DbPort, c.DbUsername, c.DbName, c.DbPassword)
	connect, err := gorm.Open("postgres", dbargs)
	if err != nil {
		log.Fatal(err)
	}

	connect.DB().SetMaxIdleConns(20)
	connect.DB().SetMaxOpenConns(200)

	connect.LogMode(true)

	dbConnection = connect

	logrus.Info("Database connected.")

	return dbConnection, err
}
