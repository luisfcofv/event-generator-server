package world

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/luisfcofv/indexter/aws"
	"github.com/luisfcofv/indexter/models"
	"github.com/luisfcofv/indexter/player"
)

func CreateWorld() {
	player := player.Player{
		Name: "Luis",
	}

	locations := createLocations()

	myWorld := world{
		Name:      "My world",
		Player:    player,
		Locations: locations,
	}

	newWorldAttributes, err := dynamodbattribute.MarshalMap(myWorld)
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
}

func createLocations() []models.Location {
	location1 := models.Location{1, "City 1", "Description 1", nil}
	location2 := models.Location{2, "City 2", "Description 2", nil}
	location3 := models.Location{3, "City 3", "Description 3", nil}
	location4 := models.Location{4, "City 4", "Description 4", nil}
	location5 := models.Location{5, "City 5", "Description 5", nil}
	location6 := models.Location{6, "City 6", "Description 6", nil}
	location7 := models.Location{7, "City 7", "Description 7", nil}
	location8 := models.Location{8, "City 8", "Description 8", nil}

	location1.Connect(&location5, 0)
	location1.Connect(&location8, 0)
	location2.Connect(&location3, 0)
	location2.Connect(&location8, 0)
	location3.Connect(&location4, 0)
	location3.Connect(&location7, 0)
	location4.Connect(&location7, 0)
	location6.Connect(&location7, 0)

	return []models.Location{
		location1,
		location2,
		location3,
		location4,
		location5,
		location6,
		location7,
		location8,
	}
}
