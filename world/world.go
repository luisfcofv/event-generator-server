package world

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/graphql-go/graphql"
	"github.com/luisfcofv/indexter/aws"
	"github.com/luisfcofv/indexter/models"
	"github.com/luisfcofv/indexter/player"
)

type world struct {
	Name      string            `json:"name"`
	State     interface{}       `json:"state"`
	Player    player.Player     `json:"player"`
	Locations []models.Location `json:"locations"`
	Social    []models.Agent    `json:"social"`
}

var worldType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "World",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"player": &graphql.Field{
				Type: player.PlayerType,
			},
			"state": &graphql.Field{
				Type: graphql.String,
			},
			"locations": &graphql.Field{
				Type: graphql.NewList(models.LocationType),
			},
			"social": &graphql.Field{
				Type: graphql.NewList(models.AgentType),
			},
		},
	},
)

var WorlField = &graphql.Field{
	Type: worldType,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		name, isOK := p.Args["name"].(string)
		if isOK {
			return getWorld(name), nil
		}
		return nil, nil
	},
}

func getWorld(name string) world {
	paramsItem := &dynamodb.GetItemInput{
		TableName: aws.String(db.AppConfig.Table),
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(name),
			},
		},
	}

	worldInstance := world{}

	// Make the DynamoDB Query API call
	result, err := db.DynamodbClient.GetItem(paramsItem)
	if err != nil {
		fmt.Errorf("failed to make GetItem API call", err)
		return worldInstance
	}

	// Unmarshal the Items field in the result value to the Item Go type.
	errUnmarshal := dynamodbattribute.UnmarshalMap(result.Item, &worldInstance)
	if errUnmarshal != nil {
		fmt.Errorf("failed to unmarshal map", errUnmarshal)
		return worldInstance
	}

	fmt.Println(worldInstance)

	return worldInstance
}
