package service

import (
	"PersonBlog/config"
	"PersonBlog/model"
	"PersonBlog/request"
)

func RegisterUser(user *model.User) (*model.User, error) {
	if err := config.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func LoginUser(user *model.User) (*model.User, error) {
	if err := config.DB.Where("username = ? ", user.Username).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func QuertyUser(user *model.User) (*model.User, error) {
	if err := config.DB.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func CreatePost(req request.PostRequest) (*model.Post, error) {
	user := model.User{}
	if useErr := config.DB.Model(&model.User{}).Where("username = ?", req.Username).First(&user).Error; useErr != nil {
		return nil, useErr
	}

	post := &model.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  user.ID,
	}
	if err := config.DB.Model(&model.Post{}).Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func GetPostList(req request.PostRequest) (*[]model.Post, error) {
	var list *[]model.Post
	postErr := config.DB.Model(&model.Post{}).Find(list).Error
	if postErr != nil {
		return nil, postErr
	}
	return list, nil
}

func GetPostInfo(req request.PostRequest) (*model.Post, error) {
	post := model.Post{}
	if err := config.DB.Model(&model.Post{}).Where("title = ?", req.Title).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func UpdatePost(req request.PostRequest) error {
	user := model.User{}
	if useErr := config.DB.Model(&model.User{}).Where("username = ?", req.Username).First(&user).Error; useErr != nil {
		return useErr
	}
	err := config.DB.Model(&model.Post{}).Where("user_id = ?", req.Username).Updates(model.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  user.ID,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func DeletePost(req request.PostRequest) error {
	user := model.User{}
	if useErr := config.DB.Model(&model.User{}).Where("username = ?", req.Username).First(&user).Error; useErr != nil {
		return useErr
	}

	err := config.DB.Model(&model.Post{}).Where("user_id = ? and title = ?", user.ID, req.Title).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateComment(req request.PostRequest) error {
	user := model.User{}
	if useErr := config.DB.Model(&model.User{}).Where("username = ?", req.Username).First(&user).Error; useErr != nil {
		return useErr
	}
	post := model.Post{}
	if postErr := config.DB.Model(&model.Post{}).Where("title = ?", req.Title).First(&post).Error; postErr != nil {
		return postErr
	}
	comment := model.Comment{
		UserID:  user.ID,
		PostID:  post.ID,
		Content: req.Content,
	}
	if err := config.DB.Create(&comment).Error; err != nil {
		return err
	}
	return nil
}

func GetComment(req request.PostRequest) (comment *[]model.Comment, err error) {
	post := model.Post{}
	if postErr := config.DB.Model(&model.Post{}).Where("title = ?", req.Title).First(&post).Error; postErr != nil {
		return nil, postErr
	}
	var res []model.Comment
	if err := config.DB.Where("post_id = ?", post.ID).Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}
