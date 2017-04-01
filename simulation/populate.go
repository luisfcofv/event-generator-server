package simulation

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/luisfcofv/indexter/aws"
	"github.com/luisfcofv/indexter/models"
	"github.com/luisfcofv/indexter/player"
	"github.com/luisfcofv/indexter/templates"
)

func createWorld() {
	player := createPlayer()
	goals := createGoals()
	locations := createLocations()
	state := createInitialState()
	agents := createAgents()

	myWorld := models.World{
		Name:      "My world",
		Player:    player,
		Locations: locations,
		Agents:    agents,
		State:     state,
		Goals:     goals,
	}

	latestEvents := templates.GetEventTemplates(myWorld)
	myWorld.LatestEvents = latestEvents

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

func createPlayer() player.Player {
	knowledge := player.Knowledge{
		Social:    []int{1},
		Locations: []int{1},
		Goals:     []int{1},
	}

	return player.Player{
		Name:      "Hero",
		Knowledge: knowledge,
	}
}

func createGoals() []models.Goal {
	goal1 := models.Goal{1, "Talk to the king", "Talk to the king"}
	goal2 := models.Goal{2, "Find the treasure", "Find the treasure"}
	goal3 := models.Goal{3, "Collect coins", "Collect coins"}
	goal4 := models.Goal{4, "Fight the dragon", "Fight the dragon"}
	goal5 := models.Goal{5, "Rescue the princess", "Rescue the princess"}

	return []models.Goal{
		goal1,
		goal2,
		goal3,
		goal4,
		goal5,
	}
}

func createLocations() []models.Location {
	location1 := models.Location{1, "Castle", "Description 1", nil}
	location2 := models.Location{2, "Main Town", "Description 2", nil}
	location3 := models.Location{3, "Small Town", "Description 3", nil}
	location4 := models.Location{4, "Valley", "Description 4", nil}
	location5 := models.Location{5, "Volcano", "Description 5", nil}

	location1.Connect(&location2, 2)
	location2.Connect(&location4, 1)
	location2.Connect(&location5, 4)
	location3.Connect(&location5, 3)
	location4.Connect(&location5, 7)

	return []models.Location{
		location1,
		location2,
		location3,
		location4,
		location5,
	}
}

func createAgents() []models.Agent {
	agent1 := models.Agent{1, "King", "The king", nil}
	agent2 := models.Agent{2, "Witness", "Witness", nil}
	agent3 := models.Agent{3, "Wizard", "Wizard", nil}
	agent4 := models.Agent{4, "Queen", "The Queen", nil}
	agent5 := models.Agent{5, "Traveler", "Traveler ", nil}

	agent1.Connect(&agent2)
	agent2.Connect(&agent3)
	agent2.Connect(&agent5)
	agent4.Connect(&agent5)

	return []models.Agent{
		agent1,
		agent2,
		agent3,
		agent4,
		agent5,
	}
}

func createInitialState() models.State {
	return models.State{
		Player: models.PlayerState{
			Location: 1,
		},
	}
}
