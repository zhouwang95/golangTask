package task03

import "gorm.io/gorm"

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	user := User{}
	tx.Where("id = ?", p.UserId).Find(&user)
	tx.Model(&User{}).Where("id = ?", p.UserId).Update("post_num", user.PostNum+1)
	return
}

func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	post := Post{}
	tx.Where("id = ?", c.PostId).Find(&post)
	tx.Model(&Post{}).Where("id = ?", c.PostId).Update("comment_num", post.CommentNum+1)
	return
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	post := Post{}
	tx.Unscoped().Where("id = ?", c.PostId).Find(&post)

	num := post.CommentNum - 1
	tx.Model(&Post{}).Where("id = ?", post.ID).Update("comment_num", num)
	if num == 0 {
		tx.Model(&Post{}).Where("id = ?", post.ID).Update("comment_staus", "无评论")
	}
	return
}
func Run(db *gorm.DB) {
	db.AutoMigrate(&User{})
	//db.AutoMigrate(&Post{})
	//db.AutoMigrate(&Comment{})

	user := User{Name: "杜甫"}
	res1 := db.Create(&user)
	if res1 != nil {
		post := Post{Title: "石壕吏", UserId: user.ID}
		res2 := db.Create(&post)
		if res2 != nil {
			comment := Comment{PostId: post.ID, CommText: "石壕吏评论1"}
			db.Create(&comment)
			comment = Comment{PostId: post.ID, CommText: "石壕吏评论2"}
			db.Create(&comment)
		}

		post = Post{Title: "潼关吏", UserId: user.ID}
		res2 = db.Create(&post)
		if res2 != nil {
			comment := Comment{PostId: post.ID, CommText: "潼关吏评论1"}
			db.Create(&comment)
			comment = Comment{PostId: post.ID, CommText: "潼关吏评论2"}
			db.Create(&comment)
		}

		post = Post{Title: "新安吏", UserId: user.ID}
		res2 = db.Create(&post)
		if res2 != nil {
			comment := Comment{PostId: post.ID, CommText: "新安吏评论1"}
			db.Create(&comment)
			comment = Comment{PostId: post.ID, CommText: "新安吏评论2"}
			db.Create(&comment)
		}
	}

	user = User{Name: "李白"}
	res1 = db.Create(&user)
	if res1 != nil {
		post := Post{Title: "蜀道难", UserId: user.ID}
		res2 := db.Create(&post)
		if res2 != nil {
			comment := Comment{PostId: post.ID, CommText: "蜀道难评论1"}
			db.Create(&comment)
			comment = Comment{PostId: post.ID, CommText: "蜀道难评论2"}
			db.Create(&comment)
		}

		post = Post{Title: "将进酒", UserId: user.ID}
		res2 = db.Create(&post)
		if res2 != nil {
			comment := Comment{PostId: post.ID, CommText: "将进酒评论1"}
			db.Create(&comment)
			comment = Comment{PostId: post.ID, CommText: "将进酒评论2"}
			db.Create(&comment)
		}

	}
}

func DeleteComm(db *gorm.DB) {
	var comm Comment
	db.Unscoped().Where("id = 1").Find(&comm).Delete(&comm)
}
