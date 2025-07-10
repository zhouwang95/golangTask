package controller

import (
	"PersonBlog/logger"
	"PersonBlog/model"
	"PersonBlog/request"
	"PersonBlog/response"
	"PersonBlog/service"
	"PersonBlog/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	req := request.LoginRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logger.AddLog(response.ParamError, "参数错误", err)
		response.Fail(c, response.ParamError, "参数错误")
		return
	}

	user, err := service.LoginUser(&model.User{
		Username: req.Username,
	})

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		logger.AddLog(response.NameOrPasswordErr, "用户名或密码错误", err)
		response.Fail(c, response.NameOrPasswordErr, "用户名或密码错误"+err.Error())
		return
	}
	// 调用封装的 token 生成方法
	token, _ := utils.GenerateToken(user.Username)

	response.Success(c, gin.H{
		"token":     token,
		"user_id":   user.ID,
		"user_name": user.Username,
	})
}
