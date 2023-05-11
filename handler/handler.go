package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/graphql-go/graphql"
	"github.com/faruqii/Go-GraphQL/graph"
)

func GraphqlHandler(w http.ResponseWriter, r *http.Request) {
    // Read the request body
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error reading request body", http.StatusBadRequest)
        return
    }

    // Execute the GraphQL query
    params := graphql.Params{
        Schema:        graph.Schema,
        RequestString: string(body),
    }
    result := graphql.Do(params)

    // Write the response
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    json.NewEncoder(w).Encode(result)
}
