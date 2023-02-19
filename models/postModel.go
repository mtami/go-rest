package models

import (
	"fmt"

	"github.com/mtami/go-crud/config"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string
	Body  string
}

func (b *Post) TableName() string {
	return "post"
}

func GetAllPost(p *[]Post) (err error) {
	if err = config.GetDB().Find(p).Error; err != nil {
		return err
	}
	return nil
}

func AddNewPost(p *Post) (err error) {
	if err = config.GetDB().Create(p).Error; err != nil {
		return err
	}
	return nil
}

func GetOnePost(p *Post, id string) (err error) {
	if err := config.GetDB().Where("id = ?", id).First(p).Error; err != nil {
		return err
	}
	return nil
}

func PutOnePost(p *Post, id string) (err error) {
	fmt.Println(p)
	config.GetDB().Save(p)
	return nil
}

func DeletePost(p *Post, id string) (err error) {
	config.GetDB().Where("id = ?", id).Delete(p)
	return nil
}
