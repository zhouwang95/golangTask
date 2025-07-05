package task03

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Password string
	Age      uint
	PostNum  int
	Post     []Post `gorm:"foreignKey:UserId"`
}

type Post struct {
	gorm.Model
	Title        string
	UserId       uint
	CommentNum   int
	CommentStaus string
	Comment      []Comment `gorm:"foreignKey:PostId"`
}

type Comment struct {
	gorm.Model
	PostId   uint
	CommText string
}

/*func Run(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})

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
}*/
