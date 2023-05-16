package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID        string `json:"id" gorm:"primary_key, type:uuid, default:uuid_generate_v4()" graphql:"id" description:"Book's ID" example:"1"`
	Title     string `json:"title" gorm:"type:varchar(100)" graphql:"title" description:"Book's title" example:"The Great Gatsby"`
	Author    string `json:"author" gorm:"type:varchar(100)" graphql:"author" description:"Book's author" example:"F. Scott Fitzgerald"`
	Year      int    `json:"year" gorm:"type:int" graphql:"year" description:"Book's year" example:"1925"`
	Publisher string `json:"publisher" gorm:"type:varchar(100)" graphql:"publisher" description:"Book's publisher" example:"Charles Scribner's Sons"`
}

func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.NewString()
	return
}
