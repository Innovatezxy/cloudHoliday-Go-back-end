package models

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type Admin struct {
	Id      string `json:"id"`
	Pwd     string `json:"pwd"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Male    string `json:"male"`
	Academy string `json:"academy"`
	Phone   string `json:"phone"`
}

type Teacher struct {
	Openid  string `json:"openid"`
	Unionid string `json:"unionid"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Male    string `json:"male"`
	Academy string `json:"academy"`
	Phone   string `json:"phone"`
	Number  string `json:"number"`
}

type Students struct {
	Openid  string `json:"openid"`
	Unionid string `json:"unionid"`
	FormId  string `json:"formid"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Male    string `json:"male"`
	Stuid   string `json:"stuid"`
	Academy string `json:"academy"`
	Class   string `json:"class"`
	Phone   string `json:"phone"`
	Number  string `json:"number"`
}

func QueryAdminInfo(c *gin.Context) {
	id := c.PostForm("id")
	DB, err := sqlx.Open("mysql", database)
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "DB_error",
			"results": err.Error(),
		})
	}

	var persons []Admin
	err = DB.Select(&persons,"SELECT * FROM admin WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "DB_error",
			"results": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"results": persons,
		})
	}
}

func QueryTeacherInfo(c *gin.Context) {
	openid := c.PostForm("id")
	DB, err := sqlx.Open("mysql", database)
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "DB_error",
			"results": err.Error(),
		})
	}

	var persons []Teacher
	err = DB.Select(&persons,"SELECT * FROM teacher WHERE openid=?", openid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "DB_error",
			"results": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"results": persons,
		})
	}
}

func QueryStdInfo(c *gin.Context) {
	openid := c.PostForm("id")
	DB, err := sqlx.Open("mysql", database)
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "DB_error",
			"results": err.Error(),
		})
	}

	var persons []Students
	err = DB.Select(&persons,"SELECT * FROM students WHERE openid=?", openid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "DB_error",
			"results": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"results": persons,
		})
	}
}

func EditTeacherPhone(c *gin.Context) {
	openid := c.PostForm("id")
	phone := c.PostForm("phone")

	DB,err := sqlx.Open("mysql",database)
	defer DB.Close()

	DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "DB_error",
			"results": err.Error(),
		})
	}

	result,err := DB.Exec("UPDATE teacher SET phone=? WHERE openid=?",phone,openid)

	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"status": "error",
			"results": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"status": "success",
			"results": result,
		})
	}
}

func EditStdPhone(c *gin.Context) {
	openid := c.PostForm("id")
	phone := c.PostForm("phone")

	DB,err := sqlx.Open("mysql",database)
	defer DB.Close()

	DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "DB_error",
			"results": err.Error(),
		})
	}

	result,err := DB.Exec("UPDATE students SET phone=? WHERE openid=?",phone,openid)

	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"status": "error",
			"results": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"status": "success",
			"results": result,
		})
	}
}

func EditStdClass(c *gin.Context) {
	openid := c.PostForm("id")
	class := c.PostForm("class")

	DB,err := sqlx.Open("mysql",database)
	defer DB.Close()

	DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "DB_error",
			"results": err.Error(),
		})
	}

	result,err := DB.Exec("UPDATE students SET class=? WHERE openid=?",class,openid)

	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"status": "error",
			"results": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"status": "success",
			"results": result,
		})
	}
}
