package simulation

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/luisfcofv/indexter/world"
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
			"world": world.WorlField,
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
	fmt.Println("Serve")
	svc := dynamodb.New(session.New(&aws.Config{Region: aws.String("eu-west-2")}))
	fmt.Println("dynamodb")
	result, err := svc.ListTables(&dynamodb.ListTablesInput{})
	fmt.Println("result", result)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Tables:")
	for _, table := range result.TableNames {
		fmt.Println(*table)
	}

	paramsItem := &dynamodb.GetItemInput{
		TableName: aws.String("World"),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				N: aws.String("1"),
			},
		},
	}

	// Make the DynamoDB Query API call
	pResult, pRrr := svc.GetItem(paramsItem)
	fmt.Println(pResult)
	if err != nil {
		fmt.Println("failed to make GetItem API call", pRrr)
		return
	}

	world := world.World{}

	// Unmarshal the Items field in the result value to the Item Go type.
	err = dynamodbattribute.UnmarshalMap(pResult.Item, &world)

	if err != nil {
		fmt.Println(err)

	}

	// Print out the items returned
	fmt.Println(world.Locations)

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query()["query"][0], schema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Indexter is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
