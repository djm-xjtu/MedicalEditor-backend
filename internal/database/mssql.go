package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// var DB *gorm.DB

// "sqlserver://sa:czacza-20001207@localhost:1433?database=TestDB"

func InitMssqlDB() error {
	jsonFile, err := os.Open("configs/mssql.json")
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var jsonConfig map[string]map[string]string
	json.Unmarshal(byteValue, &jsonConfig)

	conf := jsonConfig["ConnectionConfig"]
	dataSourceName := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", conf["user"], conf["password"], conf["host"], conf["database"])

	DB, err = gorm.Open(sqlserver.Open(dataSourceName),&gorm.Config{})
	if err != nil {
		return err
	}

	if DB.Error != nil {
        return DB.Error
    }

	log.Println("connnect success")
	return nil
}
