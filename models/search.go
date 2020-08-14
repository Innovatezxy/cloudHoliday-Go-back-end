package models

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func SearchTeacher(c *gin.Context) {
	name := c.PostForm("name")

	DB, err := sqlx.Open("mysql",database)
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "DB_error",
			"results": err.Error(),
		})
	}

	var persons []Teacher
	err = DB.Select(&persons,"SELECT * FROM teacher WHERE name LIKE ? AND type NOT IN ('l')","%" + name + "%")

	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"status": "DB_error",
			"results": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"results": persons,
		})
	}
}

func SearchLeader(c *gin.Context) {
	name := c.PostForm("name")

	DB, err := sqlx.Open("mysql",database)
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "DB_error",
			"results": err.Error(),
		})
	}

	var persons []Teacher
	err = DB.Select(&persons,"SELECT * FROM teacher WHERE name LIKE ? AND type='l'","%" + name + "%")

	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"status": "DB_error",
			"results": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"results": persons,
		})
	}
}