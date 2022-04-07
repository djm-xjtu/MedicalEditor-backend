package database

import (
	"editor-backend/internal/entities"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	jsonFile, err := os.Open("configs/mysql-config.json")
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var jsonConfig map[string]map[string]string
	json.Unmarshal(byteValue, &jsonConfig)

	conf := jsonConfig["ConnectionConfig"]
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", conf["user"], conf["password"], conf["host"], conf["database"])

	DB, err = gorm.Open(mysql.Open(dataSourceName),&gorm.Config{})
	if err != nil {
		return err
	}

	if DB.Error != nil {
        return DB.Error
    }

	log.Println("connnect success")
	DB.AutoMigrate(&entities.Department{})
	return nil
}
