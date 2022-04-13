package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// "sqlserver://sa:czacza-20001207@localhost:1433?database=TestDB"

func InitDB() error {
	jsonFile, err := os.Open("configs/mysql.json")
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var jsonConfig map[string]map[string]string
	json.Unmarshal(byteValue, &jsonConfig)

	conf := jsonConfig["ConnectionConfig"]
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", conf["user"], conf["password"], conf["host"], conf["database"])
	fmt.Println(dataSourceName)
	DB, err = gorm.Open(mysql.Open(dataSourceName),&gorm.Config{})
	if err != nil {
		return err
	}

	if DB.Error != nil {
        return DB.Error
    }

	log.Println("connnect success")
	return nil
}
