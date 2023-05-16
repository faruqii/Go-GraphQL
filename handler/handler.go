package handler

import (
	"github.com/faruqii/Go-GraphQL/config"
	"github.com/faruqii/Go-GraphQL/models"
	"github.com/graphql-go/graphql"
)

var BookType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Book",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: graphql.String,
			},
			"year": &graphql.Field{
				Type: graphql.Int,
			},
			"publisher": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func CreateBook(params graphql.ResolveParams) (interface{}, error) {
	title, _ := params.Args["title"].(string)
	author, _ := params.Args["author"].(string)
	year, _ := params.Args["year"].(int)
	publisher, _ := params.Args["publisher"].(string)

	book := &models.Book{
		Title:     title,
		Author:    author,
		Year:      year,
		Publisher: publisher,
	}

	config.DB.Create(book)

	return book, nil
}

func GetBooks(params graphql.ResolveParams) (interface{}, error) {
	var books []models.Book
	config.DB.Find(&books)
	return books, nil
}