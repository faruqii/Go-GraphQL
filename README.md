# Install Dependencies
```bash
$ go mod tidy
```

# How To run
```bash
$ go run ./cmd/main.go
```

# How To test
Url: http://localhost:3000/graphql 

### POST
```json
{
  "query": "mutation createBook($title: String!, $author: String!, $year: Int!, $publisher: String!) {\n  createBook(title: $title, author: $author, year: $year, publisher: $publisher) {\n    id\n    title\n    author\n    year\n    publisher\n  }\n}",
  "variables": {
    "title": "Book 1",
    "author": "Author 1",
    "year": 2021,
    "publisher": "Publisher 1"
  }
}

```

### GET
`http://localhost:3000/graphql?query={books{title, author, year, publisher}}`