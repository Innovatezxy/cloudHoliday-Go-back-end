package models

import (
	"CloudHoliday/conf"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"net/http"
)

var database = conf.Database

func AdminRegister(c *gin.Context) {
	id := c.PostForm("id")
	pwd := c.PostForm("pwd")
	person_type := c.PostForm("type")
	name :=c.PostForm("name")
	male := c.PostForm("male")
	academy := c.PostForm("academy")
	phone := c.PostForm("phone")

	DB, err := sqlx.Open("mysql",database)
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "DB_error",
			"results": err.Error(),
		})
	}

	result, err := DB.Exec("INSERT INTO admin (id, pwd, type, name, male, academy, phone) VALUES (?, ?, ?, ?, ?, ?, ?)",id,pwd,person_type,name,male,academy,phone)
	 if err != nil {
		 c.JSON(http.StatusOK,gin.H{
			 "status": "existed",
		 })
	 }else{
	 	c.JSON(http.StatusOK,gin.H{
	 		"status": "success",
	 		"results": result,
		})
	 }
}

func TeacherRegister(c *gin.Context) {
	openid := c.PostForm("openid")
	unionid := c.PostForm("unionid")
	person_type := c.PostForm("type")
	name :=c.PostForm("name")
	male := c.PostForm("male")
	academy := c.PostForm("academy")
	phone := c.PostForm("phone")
	number := c.PostForm("number")

	DB, err := sqlx.Open("mysql",database)
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "DB_error",
			"results": err.Error(),
		})
	}

	result, err := DB.Exec("INSERT INTO teacher (openid, unionid, type, name, male, academy, phone, number) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",openid,unionid,person_type,name,male,academy,phone,number)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"status": "existed",
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"status": "success",
			"results": result,
		})
	}
}

func StdRegister(c *gin.Context) {
	openid := c.PostForm("openid")
	unionid := c.PostForm("unionid")
	formid := c.PostForm("formid")
	person_type := c.PostForm("type")
	name :=c.PostForm("name")
	male := c.PostForm("male")
	stuid := c.PostForm("stuid")
	academy := c.PostForm("academy")
	class := c.PostForm("class")
	phone := c.PostForm("phone")
	number := c.PostForm("number")

	DB, err := sqlx.Open("mysql",database)
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "DB_error",
			"results": err.Error(),
		})
	}

	result, err := DB.Exec("INSERT INTO students (openid, unionid, formid, type, name, male, stuid, academy, class, phone, number) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",openid,unionid,formid,person_type,name,male,stuid,academy,class,phone,number)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"status": "existed",
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"status": "success",
			"results": result,
		})
	}
}
