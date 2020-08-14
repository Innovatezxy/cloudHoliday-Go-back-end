package models

import(
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/qiniu/api.v7/auth"
	auth2 "github.com/qiniu/api.v7/auth"
	"github.com/qiniu/api.v7/sms"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var manager *sms.Manager // 七牛云短信

func init() {
	accessKey := "XXX"
	secretKey := "XXX"

	mac := auth.New(accessKey, secretKey)
	manager = sms.NewManager((*auth2.Credentials)(mac))
}

// 生成随机验证码
func VerificationCode(width int) string {
	numeric := [10]byte{0,1,2,3,4,5,6,7,8,9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[ rand.Intn(r) ])
	}
	return sb.String()
}

// 反转验证码字符串
func reverseString(r []rune) string{
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// 发送验证码
func SendVerificationCode(c *gin.Context) {
	var code = VerificationCode(6)
	var reverseCode = reverseString([]rune(code))
	phone := c.PostForm("phone")

	args := sms.MessagesRequest{
		SignatureID: "1166286721294278656",
		TemplateID:  "1166288585549156352",
		Mobile:      phone,
		Parameters: map[string]interface{}{
			"code": code,
		},
	}

	ret, err := manager.SendMessage(args)

	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"status": "error",
			"results": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"status": "success",
			"results": reverseCode,
		})
	}

	fmt.Print(ret)
}

func SendChangePhoneCode(c *gin.Context) {
	var code = VerificationCode(6)
	var reverseCode = reverseString([]rune(code))
	phone := c.PostForm("phone")

	args := sms.MessagesRequest{
		SignatureID: "1166286721294278656",
		TemplateID:  "1170228354318602240",
		Mobile:      phone,
		Parameters: map[string]interface{}{
			"code": code,
		},
	}

	ret, err := manager.SendMessage(args)

	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"status": "error",
			"results": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"status": "success",
			"results": reverseCode,
		})
	}

	fmt.Print(ret)
}

// 发送请假单审批
func SendApplication(c *gin.Context) {
	phone := c.PostForm("phone")
	class := c.PostForm("class")
	name := c.PostForm("name")

	args := sms.MessagesRequest{
		SignatureID: "1166286721294278656",
		TemplateID:  "1166288813534744576",
		Mobile:      phone,
		Parameters: map[string]interface{}{
			"class": class,
			"name": name,
		},
	}

	ret, err := manager.SendMessage(args)

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

	fmt.Print(ret)
}

// 发送请假成功通知To教师
func SendSuccessTApplication(c *gin.Context)  {
	phone := c.PostForm("phone")
	class := c.PostForm("class")
	name := c.PostForm("name")
	stuid := c.PostForm("stuid")
	course := c.PostForm("course")

	args := sms.MessagesRequest{
		SignatureID: "1166286721294278656",
		TemplateID:  "1166293662053703680",
		Mobile:      phone,
		Parameters: map[string]interface{}{
			"class": class,
			"name": name,
			"stuid": stuid,
			"course": course,
		},
	}

	ret, err := manager.SendMessage(args)

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

	fmt.Print(ret)
}

// 发送请假成功通知To学生
func SendSuccessSApplication(c *gin.Context)  {
	phone := c.PostForm("phone")
	id := c.PostForm("id")
	result := c.PostForm("result")

	args := sms.MessagesRequest{
		SignatureID: "1166286721294278656",
		TemplateID:  "1169626608491831296",
		Mobile:      phone,
		Parameters: map[string]interface{}{
			"id": id,
			"result": result,
		},
	}

	ret, err := manager.SendMessage(args)

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

	fmt.Print(ret)
}
