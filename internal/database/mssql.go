package database

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// var DB *gorm.DB
type Connection struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `ysml:"database"`
}
var DB *gorm.DB
func InitMssqlDB() error {
	yamlFile, err := ioutil.ReadFile("configs/mssql.yaml")
	if err != nil {
		return err
	}

	connection := Connection{}
	yaml.Unmarshal(yamlFile, &connection)
	log.Println(connection)
	dataSourceName := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		connection.User, connection.Password, connection.Host, connection.Port, connection.Database)
	log.Println(dataSourceName)
	DB, err = gorm.Open(sqlserver.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return err
	}

	if DB.Error != nil {
		return DB.Error
	}

	log.Println("connnect success")
	return nil
}
