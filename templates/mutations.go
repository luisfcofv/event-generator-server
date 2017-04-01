package templates

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/graphql-go/graphql"

	"github.com/luisfcofv/indexter/aws"
	"github.com/luisfcofv/indexter/generator"
	"github.com/luisfcofv/indexter/models"
	"github.com/luisfcofv/indexter/player"
)

var GenerateEventsField = &graphql.Field{
	Type: graphql.NewList(models.EventType),
	Args: graphql.FieldConfigArgument{
		"world": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"knowledge": &graphql.ArgumentConfig{
			Type: player.KnowledgeInputType,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		worldName, _ := params.Args["world"].(string)
		world := models.GetWorld(worldName)

		currentKnowledge, _ := params.Args["knowledge"]
		if currentKnowledge != nil {
			fmt.Println(currentKnowledge)

		}

		eventTemplates := GetEventTemplates(world)
		world.LatestEvents = eventTemplates
		generator.Compute(&world)

		newWorldAttributes, err := dynamodbattribute.MarshalMap(world)
		if err != nil {
			fmt.Println(err)
		}

		item := &dynamodb.PutItemInput{
			TableName: aws.String(db.AppConfig.Table),
			Item:      newWorldAttributes,
		}

		_, err = db.DynamodbClient.PutItem(item)
		if err != nil {
			fmt.Println(err)
		}

		return eventTemplates, nil
	},
}

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"generateEvents": GenerateEventsField,
	},
})
