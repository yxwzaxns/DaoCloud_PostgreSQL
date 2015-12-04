package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"os"
)

func main() {
	host := os.Getenv("POSTGRESQL_PORT_5432_TCP_ADDR")
	port := os.Getenv("POSTGRESQL_PORT_5432_TCP_PORT")
	username := os.Getenv("POSTGRESQL_USERNAME")
	password := os.Getenv("POSTGRESQL_PASSWORD")

	dbname := os.Getenv("POSTGRESQL_INSTANCE_NAME")

	connection_info := "host=" + host + " port=" + port + " user=" + username + " password=" + password + " dbname=" + dbname + " sslmode=disable"
	fmt.Println(connection_info)
	db, err := sql.Open("POSTGRESQL", connection_info)
	if err != nil {
		fmt.Println("Open Error")
		return
	}

	if _, err := db.Exec("CREATE TABLE Persons(Name varchar(255))"); err != nil {
		fmt.Println(err.Error())
	}

	if _, err := db.Exec("INSERT INTO Persons (name) VALUES ('DaoCloud')"); err != nil {
		fmt.Println(err.Error())
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/", func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM Persons")
		if err != nil {
			fmt.Println(err.Error())
		}
		var data string

		for rows.Next() {
			var name string
			_ = rows.Scan(&name)
			data += name +"\n"
		}

		fmt.Println(data)
		c.JSON(200, data)
	})

	r.Run(":8080")
}
