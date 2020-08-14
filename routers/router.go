package routers

import (
	"CloudHoliday/models"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 获取唯一标识符OpenID
	r.POST("/getoid", models.GetOpenid)
	// 获取AccessToken
	r.GET("/newatk", models.GetAccessToken)
	// 更新AccessToken
	r.POST("/updateatk",models.UpdateAccessToken)
	// 查询AccessToken
	r.GET("/getatk",models.QueryAccessToken)
	// 模板消息推送
	r.POST("/pushnfc", models.PushNotification)
	// 更新FormID
	r.POST("/updatefid",models.UpdateFormId)

	// 发送注册验证码
	r.POST("/smscode",models.SendVerificationCode)
	// 发送更换手机号验证码
	r.POST("/smschangecode",models.SendChangePhoneCode)
	// 发送请假单审批
	r.POST("/smsapprove",models.SendApplication)
	// 发送请假成功通知To教师
	r.POST("/smssuccesstot",models.SendSuccessTApplication)
	// 发送请假成功通知To学生
	r.POST("/smssuccesstos",models.SendSuccessSApplication)

	// 用户注册
	r.POST("/adminreg", models.AdminRegister)
	r.POST("/teacherreg", models.TeacherRegister)
	r.POST("/stdreg", models.StdRegister)

	// 用户登录
	r.POST("/adminlgn", models.AdminLogin)
	r.POST("/tslgn", models.TSLogin)

	// 检索获取用户信息
	r.POST("/admininfo", models.QueryAdminInfo)
	r.POST("/teacherinfo", models.QueryTeacherInfo)
	r.POST("/stdinfo", models.QueryStdInfo)
	// 更改学生信息
	r.POST("/editstdphone",models.EditStdPhone)
	r.POST("/editstdclass",models.EditStdClass)
	// 更改教师及领导手机号信息
	r.POST("/editteacherphone",models.EditTeacherPhone)

	// 搜索教师
	r.POST("/searchusers", models.SearchTeacher)
	// 搜索领导
	r.POST("/searchleader", models.SearchLeader)

	// 检索获取用户请假信息
	r.POST("/queryask", models.QueryApplication)
	// 检索获取用户假期往返信息
	r.POST("/queryaskh", models.QueryHApplication)

	// 获取滚动通知栏信息
	r.GET("/queryrollnotice",models.QueryRollNotice)

	// 申请请假
	r.POST("/ask",models.AskForLeave)
	// 申请假期往返
	r.POST("/askh",models.AskForHoliday)
	// 取消日常请假
	r.POST("/cancelask",models.CancelApplication)
	// 取消假期往返请假
	r.POST("/cancelaskh",models.CancelHApplication)
	// 审批请假
	r.POST("/lapprove",models.LApproveApplication)
	r.POST("/bapprove",models.BApproveApplication)
	// 审批假期往返
	r.POST("/approveh",models.ApproveHApplication)
	// 转达审批
	r.POST("/transferapprove",models.TransferApproveApplication)
	// 一键审批请假
	r.POST("/onekeyapprove",models.OnekeyApproveApplication)
	// 查询假期往返属性
	r.GET("/queryisholiday",models.QueryIsHoliday)
	// 更新假期往返属性
	r.POST("/updateisholiday",models.UpdateIsHoliday)
	// 请假凭证图片上传
	r.POST("/uploadcer",models.UploadCertificate)

	return r
}
