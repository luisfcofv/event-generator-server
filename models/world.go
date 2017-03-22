package models

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/graphql-go/graphql"
	"github.com/luisfcofv/indexter/aws"
	"github.com/luisfcofv/indexter/player"
)

type World struct {
	Name         string        `json:"name"`
	State        interface{}   `json:"state"`
	Player       player.Player `json:"player"`
	Locations    []Location    `json:"locations"`
	Agents       []Agent       `json:"agents"`
	Goals        []Goal        `json:"goals"`
	LatestEvents []Event       `json:"latestEvents"`
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
				Type: graphql.NewList(LocationType),
			},
			"agents": &graphql.Field{
				Type: graphql.NewList(AgentType),
			},
			"goals": &graphql.Field{
				Type: graphql.NewList(GoalType),
			},
			"latestEvents": &graphql.Field{
				Type: graphql.NewList(EventType),
			},
		},
	},
)

var WorldField = &graphql.Field{
	Type: worldType,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		name, isOK := params.Args["name"].(string)
		if isOK {
			return GetWorld(name), nil
		}
		return nil, nil
	},
}

func GetWorld(name string) World {
	paramsItem := &dynamodb.GetItemInput{
		TableName: aws.String(db.AppConfig.Table),
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(name),
			},
		},
	}

	worldInstance := World{}

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

	return worldInstance
}
