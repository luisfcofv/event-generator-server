package simulation

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/rs/cors"
	"github.com/sogko/graphql-go-handler"

	"github.com/luisfcofv/indexter/aws"
	"github.com/luisfcofv/indexter/models"
)

/*
   Create Query object type with fields "user" has type [userType] by using GraphQLObjectTypeConfig:
       - Name: name of object type
       - Fields: a map of fields by using GraphQLFields
   Setup type of field use GraphQLFieldConfig to define:
       - Type: type of field
       - Args: arguments to query with current field
       - Resolve: function to query data using params from [Args] and return value with current type
*/
var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"world": models.WorldField,
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
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

func Serve() {
	db.Setup()
	CreateWorld()

	mux := http.NewServeMux()
	graphqlHandler := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})
	mux.Handle("/graphql", graphqlHandler)

	corsHandler := cors.Default().Handler(mux)
	fmt.Println("Indexter is running on port 8080")
	http.ListenAndServe(":8080", corsHandler)
}
