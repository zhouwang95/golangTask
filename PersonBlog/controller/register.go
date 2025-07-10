package controller

import (
	"PersonBlog/logger"
	"PersonBlog/model"
	"PersonBlog/request"
	"PersonBlog/response"
	"PersonBlog/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	req := request.RegisterRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logger.AddLog(response.ParamError, "参数错误", err)
		response.Fail(c, response.ParamError, "参数错误")
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.AddLog(response.PasswordErr, "加密密码失败", err)
		response.Fail(c, response.PasswordErr, "加密密码失败"+err.Error())
		return
	}
	user, err1 := service.RegisterUser(&model.User{
		Username: req.Name,
		Password: string(hashedPassword),
	})
	if err1 != nil {
		logger.AddLog(response.RegisterUserErr, "注册用户失败", err)
		response.Fail(c, response.RegisterUserErr, "注册用户失败"+err1.Error())
	}
	response.Success(c, *user)
}
