package controller

import (
	"PersonBlog/model"
	"PersonBlog/request"
	"PersonBlog/response"
	"PersonBlog/service"
	"github.com/gin-gonic/gin"
)

func QuertyUserInfo(c *gin.Context) {
	user := model.User{}
	quertyUser, err := service.QuertyUser(&user)
	if err != nil {
		return
	}
	response.Success(c, *quertyUser)
}

// CreatPost 实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
func CreatPost(c *gin.Context) {
	req := request.PostRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return
	}

	post, err1 := service.CreatePost(req)
	if err1 != nil {
		return
	}
	response.Success(c, *post)
}

// GetPostList 实现文章的读取功能，支持获取所有文章列表
func GetPostList(c *gin.Context) {
	req := request.PostRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return
	}
	post, err1 := service.GetPostList(req)
	if err1 != nil {
		return
	}
	response.Success(c, *post)
}

// GetPostInfo 实现文章的读取功能，支持获取单个文章的详细信息。
func GetPostInfo(c *gin.Context) {
	req := request.PostRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return
	}
	post, err1 := service.GetPostInfo(req)
	if err1 != nil {
		return
	}
	response.Success(c, *post)
}

// UpdatePost 实现文章的更新功能，只有文章的作者才能更新自己的文章。
func UpdatePost(c *gin.Context) {
	req := request.PostRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return
	}
	err1 := service.UpdatePost(req)
	if err1 != nil {
		return
	}
	response.Success(c, req)
}

// DeletePost 实现文章的删除功能，只有文章的作者才能删除自己的文章。
func DeletePost(c *gin.Context) {
	req := request.PostRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return
	}
	err1 := service.DeletePost(req)
	if err1 != nil {
		return
	}
	response.Success(c, req)
}

// CreateComment 实现评论的创建功能，已认证的用户可以对文章发表评论。
func CreateComment(c *gin.Context) {
	req := request.PostRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return
	}
	err1 := service.CreateComment(req)
	if err1 != nil {
		return
	}
	response.Success(c, req)
}

// GetComment 实现评论的读取功能，支持获取某篇文章的所有评论列表。
func GetComment(c *gin.Context) {
	req := request.PostRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return
	}
	comments, err1 := service.GetComment(req)
	if err1 != nil {
		return
	}
	response.Success(c, comments)
}
