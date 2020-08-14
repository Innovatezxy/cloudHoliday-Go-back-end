package models

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type Application struct {
	Id             string `json:"id"`
	Type           string `json:"type"`
	SOpenid        string `json:"s_openid" db:"s_openid"`
	Name           string `json:"name"`
	Male           string `json:"male"`
	Stuid          string `json:"stuid"`
	Academy        string `json:"academy"`
	Class          string `json:"class"`
	Start          string `json:"start"`
	End            string `json:"end"`
	Remark         string `json:"remark"`
	Reason         string `json:"reason"`
	ReasonImg	   string `json:"reason_img" db:"reason_img"`
	RName          string `json:"r_name" db:"r_name"`
	ROpenid        string `json:"r_openid" db:"r_openid"`
	RPhone         string `json:"r_phone"  db:"r_phone"`
	BName          string `json:"b_name" db:"b_name"`
	BOpenid        string `json:"b_openid" db:"b_openid"`
	BPhone        string `json:"b_phone"  db:"b_phone"`
	BResult        string `json:"b_result" db:"b_result"`
	BRefuseReason  string `json:"b_refuse_reason" db:"b_refuse_reason"`
	LName          string `json:"l_name" db:"l_name"`
	LOpenid        string `json:"l_openid" db:"l_openid"`
	LPhone        string `json:"l_phone"  db:"l_phone"`
	LResult        string `json:"l_result" db:"l_result"`
	LRefuseReason  string `json:"l_refuse_reason" db:"l_refuse_reason"`
	SPhone         string `json:"s_phone" db:"s_phone"`
}

type HApplication struct {
	Id             string `json:"id"`
	Type           string `json:"type"`
	SOpenid        string `json:"s_openid" db:"s_openid"`
	Name           string `json:"name"`
	Male           string `json:"male"`
	Stuid          string `json:"stuid"`
	Academy        string `json:"academy"`
	Class          string `json:"class"`
	Start          string `json:"start"`
	End            string `json:"end"`
	Remark         string `json:"remark"`
	Reason         string `json:"reason"`
	BName          string `json:"b_name" db:"b_name"`
	BOpenid        string `json:"b_openid" db:"b_openid"`
	BPhone        string `json:"b_phone"  db:"b_phone"`
	BResult        string `json:"b_result" db:"b_result"`
	BRefuseReason  string `json:"b_refuse_reason" db:"b_refuse_reason"`
	SPhone         string `json:"s_phone" db:"s_phone"`
}

type IsHoliday struct {
	Id string `json:"id"`
	IsHoliday string `json:"isholiday"`
	Remark string `json:"remark"`
}

type RollNotice struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Content string `json:"content"`
	Path string `json:"path"`
}


func QueryApplication(c *gin.Context) {
	openid := c.PostForm("id")

	DB, err := sqlx.Open("mysql", database)
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "DB_error",
			"results": err.Error(),
		})
	}

	var ask []Application
	err = DB.Select(&ask, "SELECT * FROM daily_2019_2020_up WHERE s_openid=? OR r_openid=? OR b_openid=? OR l_openid=?", openid, openid, openid, openid)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "DB_error",
			"results": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"number": len(ask),
			"results": ask,
		})
	}
}

func QueryHApplication(c *gin.Context) {
	openid := c.PostForm("id")

	DB, err := sqlx.Open("mysql", database)
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "DB_error",
			"results": err.Error(),
		})
	}

	var ask []HApplication
	err = DB.Select(&ask, "SELECT * FROM holiday_2019_2020_up WHERE s_openid=? OR b_openid=?", openid, openid)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "DB_error",
			"results": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"number": len(ask),
			"results": ask,
		})
	}
}

func QueryIsHoliday(c *gin.Context) {
	DB, err := sqlx.Open("mysql", database)
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "DB_error",
			"results": err.Error(),
		})
	}

	var result IsHoliday
	err = DB.Get(&result, "SELECT * FROM isholiday")

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "none",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"results": result,
		})
	}
}

func QueryRollNotice(c *gin.Context) {
	DB, err := sqlx.Open("mysql", database)
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "DB_error",
			"results": err.Error(),
		})
	}

	var result []RollNotice
	err = DB.Select(&result, "SELECT * FROM roll_notice")

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "none",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"results": result,
		})
	}
}
