package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Department struct {
	Id     int    `db:"id"`
	First  string `db:"first"`
	Second string `db:"second"`
	Name   string `db:"name"`
}

type NoteTemplate struct {
	NoteType string `json:"noteType"`
	Template string `json:"template"`
}

type NoteInfo struct {
	PatientId int    `json:"patientId"`
	NoteType  string `json:"noteType"`
	Note      string `json:"note"`
}

func main() {
	conf, err := loadMySqlConfig("conf/mysql-config.json")
	if err != nil {
		log.Panic(err)
	}

	db, err := initDB(conf)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	rows, e := db.Query("SELECT id, first, second, name FROM department")
	if e == nil {
		errors.New("query incur error")
	}

	defer rows.Close()
	for rows.Next() {
		var department Department
		err := rows.Scan(&department.Id, &department.First, &department.Second, &department.Name)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Println(department)
	}

	router := gin.Default()
	router.Use(Cors())

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It works")
	})

	router.GET("/departments", func(c *gin.Context) {
		rows, err := db.Query("SELECT id, first, second, name FROM department")

		if err != nil {
			log.Fatalln(err)
		}
		defer rows.Close()

		departments := make([]Department, 0)

		for rows.Next() {
			var department Department
			err := rows.Scan(&department.Id, &department.First, &department.Second, &department.Name)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return
			}

			departments = append(departments, department)
		}
		if err = rows.Err(); err != nil {
			log.Fatalln(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"departments": departments,
		})

	})

	router.POST("/note-template", func(c *gin.Context) {
		noteTemplate := NoteTemplate{}
		c.BindJSON(&noteTemplate)
		fmt.Println(noteTemplate.NoteType)
		fmt.Println(noteTemplate.Template)
		rs, err := db.Exec("INSERT INTO medical_note_templates (note_type, template) VALUES (?, ?)", noteTemplate.NoteType, noteTemplate.Template)
		if err != nil {
			log.Fatalln(err)
		}

		id, err := rs.LastInsertId()
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("insert person Id {}", id)
		msg := fmt.Sprintf("insert successful %d", id)
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
	})

	router.POST("/note", func(c *gin.Context) {
		noteInfo := NoteInfo{}
		c.BindJSON(&noteInfo)

		rs, err := db.Exec("INSERT INTO medical_notes (patient_id, note_type, note) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE note = ?", noteInfo.PatientId, noteInfo.NoteType, noteInfo.Note, noteInfo.Note)
		if err != nil {
			log.Fatalln(err)
		}

		id, err := rs.LastInsertId()
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("insert person Id {}", id)
		msg := fmt.Sprintf("insert successful %d", id)
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
	})

	router.GET("/note", func(c *gin.Context) {
		patientId := c.Query("patientId")
		noteType := c.Query("noteType")
		fmt.Printf("patientId: %s noteType: %s", patientId, noteType)

		var note string
		err := db.QueryRow("SELECT note FROM medical_notes WHERE patient_id = ? AND note_type = ?", patientId, noteType).Scan(&note)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"note": nil,
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"note": note,
		})

	})
	router.Run(":8000")
}

// 开启跨域函数
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", " Origin, X-Requested-With, Content-Type, Accept")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.Writer.Header().Set("Content-Type", "application/json;charset=utf-8")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		defer func() {
			if err := recover(); err != nil {
				log.Fatalf("Panic info is: %v", err)
				log.Fatalf("Panic info is: %s", debug.Stack())
			}
		}()

		c.Next()
	}
}

func loadMySqlConfig(path string) (map[string]string, error) {
	jsonFile, err := os.Open("conf/mysql-config.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	fmt.Println(string(byteValue))

	var jsonConfig map[string]map[string]string
	json.Unmarshal(byteValue, &jsonConfig)
	return jsonConfig["ConnectionConfig"], nil
}

func initDB(conf map[string]string) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", conf["user"], conf["password"], conf["host"], conf["database"])
	fmt.Println(dataSourceName)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	//设置数据库最大连接数
	db.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(10)
	//验证连接
	if err := db.Ping(); err != nil {
		fmt.Println("open database fail")
		return nil, err
	}

	fmt.Println("connnect success")
	return db, nil
}
