package models

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type UploadImg struct {
	Path string `json:"path"`
}

func AskForLeave(c *gin.Context){
	id := c.PostForm("id")
	askType := c.PostForm("type")
	SOpenid := c.PostForm("s_openid")
	name := c.PostForm("name")
	male := c.PostForm("male")
	stuid := c.PostForm("stuid")
	academy := c.PostForm("academy")
	class := c.PostForm("class")
	start := c.PostForm("start")
	end := c.PostForm("end")
	remark := c.PostForm("remark")
	reason := c.PostForm("reason")
	reasonImg := c.PostForm("reason_img")
	RName := c.PostForm("r_name")
	ROpenid := c.PostForm("r_openid")
	RPhone := c.PostForm("r_phone")
	BName := c.PostForm("b_name")
	BOpenid := c.PostForm("b_openid")
	BPhone := c.PostForm("b_phone")
	BResult := c.PostForm("b_result")
	BRefuseReason := c.PostForm("b_refuse_reason")
	LName := c.PostForm("l_name")
	LOpenid := c.PostForm("l_openid")
	LPhone := c.PostForm("l_phone")
	LResult := c.PostForm("l_result")
	LRefuseReason := c.PostForm("l_refuse_reason")
	SPhone := c.PostForm("s_phone")

	DB, err := sqlx.Open("mysql",database)
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"status": "error",
			"results": err.Error(),
		})
	}

	result, err := DB.Exec("INSERT INTO daily_2019_2020_up (id,type,s_openid,name,male,stuid,academy,class,start,end,remark,reason,reason_img,r_name,r_openid,r_phone,b_name,b_openid,b_phone,b_result,b_refuse_reason,l_name,l_openid,l_phone,l_result,l_refuse_reason,s_phone) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",id,askType,SOpenid,name,male,stuid,academy,class,start,end,remark,reason,reasonImg,RName,ROpenid,RPhone,BName,BOpenid,BPhone,BResult,BRefuseReason,LName,LOpenid,LPhone,LResult,LRefuseReason,SPhone)

	if err != nil{
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

func AskForHoliday(c *gin.Context){
	id := c.PostForm("id")
	askType := c.PostForm("type")
	SOpenid := c.PostForm("s_openid")
	name := c.PostForm("name")
	male := c.PostForm("male")
	stuid := c.PostForm("stuid")
	academy := c.PostForm("academy")
	class := c.PostForm("class")
	start := c.PostForm("start")
	end := c.PostForm("end")
	remark := c.PostForm("remark")
	reason := c.PostForm("reason")
	BName := c.PostForm("b_name")
	BOpenid := c.PostForm("b_openid")
	BPhone := c.PostForm("b_phone")
	BResult := c.PostForm("b_result")
	BRefuseReason := c.PostForm("b_refuse_reason")
	SPhone := c.PostForm("s_phone")

	DB, err := sqlx.Open("mysql",database)
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"status": "error",
			"results": err.Error(),
		})
	}

	result, err := DB.Exec("INSERT INTO holiday_2019_2020_up (id,type,s_openid,name,male,stuid,academy,class,start,end,remark,reason,b_name,b_openid,b_phone,b_result,b_refuse_reason,s_phone) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",id,askType,SOpenid,name,male,stuid,academy,class,start,end,remark,reason,BName,BOpenid,BPhone,BResult,BRefuseReason,SPhone)

	if err != nil{
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

func ApproveHApplication(c *gin.Context){
	id := c.PostForm("id")
	BResult := c.PostForm("b_result")
	BRefuseReason := c.PostForm("b_refuse_reason")

	DB,err := sqlx.Open("mysql",database)
	defer DB.Close()

	DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "DB_error",
			"results": err.Error(),
		})
	}

	result,err := DB.Exec("UPDATE holiday_2019_2020_up SET b_result=?,b_refuse_reason=? WHERE id=?",BResult,BRefuseReason,id)

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

func CancelApplication(c *gin.Context) {
	id := c.PostForm("id")
	DB,err := sqlx.Open("mysql",database)
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"status": "error",
			"results": err.Error(),
		})
	}

	result,err := DB.Exec("DELETE FROM daily_2019_2020_up WHERE id=?",id)

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

func CancelHApplication(c *gin.Context) {
        id := c.PostForm("id")
        DB,err := sqlx.Open("mysql",database)
        defer DB.Close()

        err = DB.Ping()
        if err != nil {
                c.JSON(http.StatusOK,gin.H{
                       	"status": "error",
                       	"results": err.Error(),
               	})
        }

        result,err := DB.Exec("DELETE FROM holiday_2019_2020_up WHERE id=?",id)

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


func LApproveApplication(c *gin.Context){
	id := c.PostForm("id")
	LResult := c.PostForm("l_result")
	LRefuseReason := c.PostForm("l_refuse_reason")
	BResult := c.PostForm("b_result")
	BRefuseReason := c.PostForm("b_refuse_reason")

	DB,err := sqlx.Open("mysql",database)
	defer DB.Close()

	DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "DB_error",
			"results": err.Error(),
		})
	}

	result,err := DB.Exec("UPDATE daily_2019_2020_up SET b_result=?,b_refuse_reason=?,l_result=?,l_refuse_reason=? WHERE id=?",BResult,BRefuseReason,LResult,LRefuseReason,id)

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

func BApproveApplication(c *gin.Context){
	id := c.PostForm("id")
	BResult := c.PostForm("b_result")
	BRefuseReason := c.PostForm("b_refuse_reason")

	DB,err := sqlx.Open("mysql",database)
	defer DB.Close()

	DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "DB_error",
			"results": err.Error(),
		})
	}

	result,err := DB.Exec("UPDATE daily_2019_2020_up SET b_result=?,b_refuse_reason=? WHERE id=?",BResult,BRefuseReason,id)

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


func TransferApproveApplication(c *gin.Context){
	id := c.PostForm("id")
	LName := c.PostForm("l_name")
	LOpenid := c.PostForm("l_openid")
	LPhone := c.PostForm("l_phone")
	LResult := c.PostForm("l_result")
	LRefuseReason := c.PostForm("l_refuse_reason")

	DB,err := sqlx.Open("mysql",database)
	defer DB.Close()

	DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "DB_error",
			"results": err.Error(),
		})
	}

	result,err := DB.Exec("UPDATE daily_2019_2020_up SET l_name=?,l_openid=?,l_phone=?,l_result=?,l_refuse_reason=? WHERE id=?",LName,LOpenid,LPhone,LResult,LRefuseReason,id)

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

func OnekeyApproveApplication(c *gin.Context){
	askType := c.PostForm("type")
	BOpenid := c.PostForm("b_openid")
	BResult := c.PostForm("b_result")
	BRefuseReason := c.PostForm("b_refuse_reason")

	DB,err := sqlx.Open("mysql",database)
	defer DB.Close()

	DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "DB_error",
			"results": err.Error(),
		})
	}

	result,err := DB.Exec("UPDATE holiday_2019_2020_up SET b_result=?,b_refuse_reason=? WHERE b_openid=? AND type=? AND b_result='none'",BResult,BRefuseReason,BOpenid,askType)

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

func UploadCertificate(c *gin.Context) {
	id := c.Query("id")

	file, err := c.FormFile("certificate")

	filename := id + "_cer.jpg"
	savepath := "/enjfun/www.enjfun.com/img/cloudholiday/certificate/" + filename

	err = c.SaveUploadedFile(file, savepath)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"status": "error",
			"results": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"status": "success",
		})
	}
}

func UpdateIsHoliday(c *gin.Context) {
	isholiday := c.PostForm("isholiday")
	remark := c.PostForm("remark")


	DB, err := sqlx.Open("mysql", database)
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "DB_error",
			"results": err.Error(),
		})
	}

	result, err := DB.Exec("UPDATE isholiday SET isholiday=?,remark=? WHERE id=1",isholiday,remark)

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
