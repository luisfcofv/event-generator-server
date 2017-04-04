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

func getValues(object interface{}) []int {
	if object == nil {
		return []int{}
	}

	values := object.([]interface{})
	var list []int

	for index := range values {
		list = append(list, values[index].(int))
	}

	return list
}

var GenerateEventsField = &graphql.Field{
	Type: models.WorldType,
	Args: graphql.FieldConfigArgument{
		"world": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"knowledge": &graphql.ArgumentConfig{
			Type: player.KnowledgeInputType,
		},
		"location": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		worldName, _ := params.Args["world"].(string)
		world := models.GetWorld(worldName)

		interfaceKnowledge, ok := params.Args["knowledge"]
		if ok {
			mapKnowledge, ok := interfaceKnowledge.(map[string]interface{})

			if ok {
				locations := getValues(mapKnowledge["locations"])
				social := getValues(mapKnowledge["social"])
				goals := getValues(mapKnowledge["goals"])

				world.Player.Knowledge.Locations = locations
				world.Player.Knowledge.Social = social
				world.Player.Knowledge.Goals = goals
			}
		}

		playerLocation, ok := params.Args["location"].(int)

		if ok {
			world.State.Player.Location = playerLocation
			fmt.Println(world.State.Player.Location)
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

		return world, nil
	},
}
