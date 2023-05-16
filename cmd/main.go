package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/faruqii/Go-GraphQL/config"
	"github.com/faruqii/Go-GraphQL/handler"
	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
)

func main() {

	// load .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database
	db, err := config.Connect()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal("Failed to get the database connection:", err)
		}
		err = sqlDB.Close()
		if err != nil {
			log.Fatal("Failed to close the database connection:", err)
		}
	}()

	// GraphQL Schema
	fields := graphql.Fields{
		"books": &graphql.Field{
			Type:        graphql.NewList(handler.BookType),
			Description: "Get all books",
			Resolve:     handler.GetBooks,
		},
		"createBook": &graphql.Field{
			Type:        handler.BookType,
			Description: "Create a new book",
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"author": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"year": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"publisher": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: handler.CreateBook,
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createBook": &graphql.Field{
				Type:        handler.BookType,
				Description: "Create a new book",
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"author": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"year": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"publisher": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: handler.CreateBook,
			},
		},
	})

	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: rootMutation,
	}
	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatal("Failed to create the GraphQL schema:", err)
	}

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			result := graphql.Do(graphql.Params{
				Schema:        schema,
				RequestString: r.URL.Query().Get("query"),
			})

			if len(result.Errors) > 0 {
				log.Printf("Failed to execute GraphQL operation, errors: %+v", result.Errors)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(result)
		} else if r.Method == "POST" {
			var request struct {
				Query     string                 `json:"query"`
				Variables map[string]interface{} `json:"variables"`
			}
			err := json.NewDecoder(r.Body).Decode(&request)
			if err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}

			// Execute the GraphQL query
			result := graphql.Do(graphql.Params{
				Schema:         schema,
				RequestString:  request.Query,
				VariableValues: request.Variables,
			})

			if len(result.Errors) > 0 {
				log.Printf("Failed to execute GraphQL operation, errors: %+v", result.Errors)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(result)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))

}
