package models

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func AdminLogin(c *gin.Context) {
	accountId := c.PostForm("id")
	accountPwd := c.PostForm("pwd")
	DB, err := sqlx.Open("mysql",database)
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "DB_error",
			"results": err.Error(),
		})
	}

	var id,pwd string
	row := DB.QueryRow("SELECT id,pwd FROM admin WHERE id=? OR pwd=?", accountId,accountPwd)
	err = row.Scan(&id,&pwd)

	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"status": "none",
		})
	} else {
		if accountId == id && accountPwd == pwd {
			c.JSON(http.StatusOK,gin.H{
				"status": "success",
			})
		} else {
			c.JSON(http.StatusOK,gin.H{
				"status": "error",
			})
		}
	}
}

func TSLogin(c *gin.Context) {
	id := c.PostForm("id")
	DB, err := sqlx.Open("mysql",database)
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "DB_error",
			"results": err.Error(),
		})
	}

	var openid1,openid2 string
	row := DB.QueryRow("SELECT teacher.openid,students.openid FROM teacher,students WHERE teacher.openid=? OR students.openid=?", id ,id)
	err = row.Scan(&openid1,&openid2)

	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"status": "none",
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"status": "success",
		})
	}
}
