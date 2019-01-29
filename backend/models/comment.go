package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//Comment is a Comment
type Comment struct {
	gorm.Model
	UserName string `db:"name" json:"name"`
	Text     string `db:"text" json:"text"`
}

func AllComments() []Comment {

	var allComments []Comment

	db.Find(&allComments)
	fmt.Println("{}", allComments)

	return allComments
}

func FindCommentByUser(name string) Comment {

	var Comment Comment

	db.Where("name = ?", name).Find(&Comment)
	fmt.Println("{}", Comment)

	return Comment
}

func CreateComment(username string, message string) {

	fmt.Println("{}", username)

	db.Create(&Comment{UserName: username, Text: message})
}

func LastAddedComment() Comment {

	var comment Comment

	db.Last(&comment)
	return comment
}

func UpdateComment(name string, text string) {
	Comment := FindCommentByUser(name)
	Comment.Text = text

	db.Save(&Comment)
}

func DeleteComment(name string) {
	Comment := FindCommentByUser(name)
	db.Delete(&Comment)
}
