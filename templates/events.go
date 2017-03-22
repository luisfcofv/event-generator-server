package templates

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/graphql-go/graphql"

	"github.com/luisfcofv/indexter/aws"
	"github.com/luisfcofv/indexter/generator"
	"github.com/luisfcofv/indexter/models"
)

func GetEventTemplates(world models.World) []models.Event {
	return []models.Event{
		getFirstTemplate(world),
		getSecondTemplate(world),
		getThirdTemplate(world),
		getFourthTemplate(world),
		getFifthTemplate(world),
	}
}

func getFirstTemplate(world models.World) models.Event {
	time := time.Now().Unix()
	rand.Seed(time)

	location := rand.Intn(len(world.Locations)) + 1
	agent := rand.Intn(len(world.Agents)) + 1
	goal := rand.Intn(len(world.Goals)) + 1

	event := models.Event{
		Name:        "Template 1",
		Description: "Random template",
		Agent:       agent,
		Location:    location,
		Goal:        goal,
		Time:        time,
	}

	return event
}

func getSecondTemplate(world models.World) models.Event {
	time := time.Now().Unix()

	rand.Seed(time)

	possibleAgents := []int{2, 4}
	agentID := rand.Intn(len(possibleAgents))

	event := models.Event{
		Name:        "Template 2",
		Description: "City 3, Witness or The Queen, Find the treasure",
		Agent:       possibleAgents[agentID],
		Location:    3,
		Goal:        2, // Find the treasure
		Time:        time,
	}

	return event
}

func getThirdTemplate(world models.World) models.Event {
	time := time.Now().Unix()

	event := models.Event{
		Name:        "Template 3",
		Description: "City 2, The king, Talk to the king",
		Agent:       1, // The king
		Location:    2,
		Goal:        1, // Talk to the king
		Time:        time,
	}

	return event
}

func getFourthTemplate(world models.World) models.Event {
	time := time.Now().Unix()

	event := models.Event{
		Name:        "Template 4",
		Description: "City 5, Protagonist, Fight the dragon",
		Agent:       0, // Protagonist
		Location:    5,
		Goal:        4, // Fight the dragon
		Time:        time,
	}

	return event
}

func getFifthTemplate(world models.World) models.Event {
	time := time.Now().Unix()

	possibleCities := []int{4, 5}
	cityID := rand.Intn(len(possibleCities))

	event := models.Event{
		Name:        "Template 5",
		Description: "City 4 or 5, Wizard, Rescue the princess",
		Agent:       3, // Wizard
		Location:    possibleCities[cityID],
		Goal:        5, // Rescue the princess
		Time:        time,
	}

	return event
}

var EventTemplatesMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"generateEvents": &graphql.Field{
			Type: graphql.NewList(models.EventType),
			Args: graphql.FieldConfigArgument{
				"world": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				worldName, _ := params.Args["world"].(string)
				world := models.GetWorld(worldName)

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
		},
	},
})
