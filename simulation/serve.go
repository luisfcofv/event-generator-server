package simulation

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/rs/cors"
	"github.com/sogko/graphql-go-handler"

	"github.com/luisfcofv/indexter/aws"
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there!")
}

func Serve() {
	db.Setup()
	createWorld()

	mux := http.NewServeMux()
	graphqlHandler := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	mux.HandleFunc("/", hello)
	mux.Handle("/graphql", graphqlHandler)

	corsHandler := cors.Default().Handler(mux)
	fmt.Println("Indexter is running on port 3000")
	http.ListenAndServe(":3000", corsHandler)
}
