package models

import (
	"CloudHoliday/conf"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"net/http"
)

var appid = conf.AppID
var secret = conf.AppSecret

type AccessToken struct {
	Token   string `json:"token"`
	Endtime string `json:"endtime"`
}

type Code2Session struct {
	SessionKey string `json:"session_key"`
	Openid     string `json:"openid"`
}

type NewAccessToken struct {
	Token   string `json:"access_token"`
	Expires string `json:"expires_in"`
}

type PushMessage struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type TemplateMsg struct {
	Touser     string        `json:"touser,omitempty"`     //接收者的OpenID
	TemplateID string        `json:"template_id,omitempty"` //模板消息ID
	URL        string        `json:"page,omitempty"`
	FormID     string        `json:"form_id,omitempty"`
	Data       *TemplateData `json:"data,omitempty"`
}

type TemplateData struct {
	Keyword1 KeyWordData `json:"keyword1,omitempty"`
	Keyword2 KeyWordData `json:"keyword2,omitempty"`
	Keyword3 KeyWordData `json:"keyword3,omitempty"`
	Keyword4 KeyWordData `json:"keyword4,omitempty"`
	Keyword5 KeyWordData `json:"keyword5,omitempty"`
	Keyword6 KeyWordData `json:"keyword5,omitempty"`
	Keyword7 KeyWordData `json:"keyword5,omitempty"`
}

type KeyWordData struct {
	Value string `json:"value"`
}

func GetOpenid(c *gin.Context) {
	var result Code2Session
	jscode := c.PostForm("jscode")
	conn, err := http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=" + appid + "&secret=" + secret + "&js_code=" + jscode + "&grant_type=authorization_code")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"results": err.Error(),
		})
	}
	defer conn.Body.Close()
	body, _ := ioutil.ReadAll(conn.Body)
	err = json.Unmarshal([]byte(string(body)), &result)
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"results": result,
	})
}

func GetAccessToken(c *gin.Context) {
	var result NewAccessToken
	conn, err := http.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appid + "&secret=" + secret)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"results": err.Error(),
		})
	}
	defer conn.Body.Close()
	body, _ := ioutil.ReadAll(conn.Body)
	err = json.Unmarshal([]byte(string(body)), &result)
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"results": result,
	})
}

func PushNotification(c *gin.Context) {
	var result PushMessage
	var msg TemplateMsg
	templateMsg := &TemplateMsg{}
	tempData := &TemplateData{}

	accessToken := c.PostForm("access_token")
	templateMsg.Touser = c.PostForm("touser")
	templateMsg.TemplateID = c.PostForm("template_id")
	templateMsg.URL = c.PostForm("page")
	templateMsg.FormID = c.PostForm("form_id")

	tempData.Keyword1.Value = c.PostForm("k1")
	tempData.Keyword2.Value = c.PostForm("k2")
	tempData.Keyword3.Value = c.PostForm("k3")
	tempData.Keyword4.Value = c.PostForm("k4")
	tempData.Keyword5.Value = c.PostForm("k5")
	tempData.Keyword6.Value = c.PostForm("k6")
	tempData.Keyword7.Value = c.PostForm("k7")
	templateMsg.Data = tempData

	data, _ := json.Marshal(msg)

	conn, err := http.Post("https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token="+accessToken, "application/json", bytes.NewBuffer(data))

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"results": err.Error(),
		})
	}
	defer conn.Body.Close()
	body, _ := ioutil.ReadAll(conn.Body)
	err = json.Unmarshal(body, &result)
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"results": result,
	})
}

func QueryAccessToken(c *gin.Context) {
	DB, err := sqlx.Open("mysql", database)
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "DB_error",
			"results": err.Error(),
		})
	}

	var result AccessToken
	err = DB.Get(&result, "SELECT token,endtime FROM access_token WHERE id=1")

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

func UpdateAccessToken(c *gin.Context) {
	accessToken := c.PostForm("access_token")
	endTime := c.PostForm("endtime")

	DB, err := sqlx.Open("mysql", database)
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "DB_error",
			"results": err.Error(),
		})
	}

	result, err := DB.Exec("UPDATE access_token SET token=?,endtime=? WHERE id=1", accessToken, endTime)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "no_expired",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"results": result,
		})
	}
}

func UpdateFormId(c *gin.Context) {
	openid := c.PostForm("openid")
	formid := c.PostForm("formid")

	DB, err := sqlx.Open("mysql", database)
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "DB_error",
			"results": err.Error(),
		})
	}

	result, err := DB.Exec("UPDATE students SET formid=? WHERE openid=?", formid, openid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "error",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"results": result,
		})
	}
}
