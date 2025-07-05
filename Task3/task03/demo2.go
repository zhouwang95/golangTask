package task03

import (
	"fmt"
	"gorm.io/gorm"
)

type UserPost struct {
	Name         string
	PostNum      int
	Title        string
	CommentNum   int
	CommentStaus string
	CommText     string
}

func Querty1(db *gorm.DB) {
	str := `
		select u.name,u.post_num,p.title,p.comment_num,p.comment_staus,c.comm_text 
		from users u 
		left join posts p on p.user_id = u.id
		left join comments c on c.post_id = p.id
		where u.name = ?;`
	var userPost []UserPost
	db.Raw(str, "杜甫").Find(&userPost)
	fmt.Println(userPost)
}

func Querty2(db *gorm.DB) {
	str := `
		select u.name,u.post_num,p.title, p.comment_num,p.comment_staus
		from posts p 
		left join comments c on c.post_id = p.id 
		left join users u on u.id = p.user_id
		order by p.comment_num desc 
		limit 1;`
	var userPost UserPost
	db.Raw(str).Find(&userPost)
	fmt.Println(userPost)
}
