package graph

import (
	"github.com/graphql-go/graphql"
	"github.com/faruqii/Go-GraphQL/models"
)

var schemaString = `
type Query {
	books: [Book]
}

type Book {
	title: String
	author: String
}
`
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"books": &graphql.Field{
				Type: graphql.NewList(bookType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Here you would normally query your database or other data source
					// to get the list of books. For this example, we'll just return
					// some hard-coded data.
					return []models.Book{
						{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
						{Title: "To Kill a Mockingbird", Author: "Harper Lee"},
						{Title: "1984", Author: "George Orwell"},
					}, nil
				},
			},
		},
	}),
})

var bookType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Book",
	Fields: graphql.Fields{
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"author": &graphql.Field{
			Type: graphql.String,
		},
	},
})
