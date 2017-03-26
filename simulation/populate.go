package simulation

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/luisfcofv/indexter/aws"
	"github.com/luisfcofv/indexter/models"
	"github.com/luisfcofv/indexter/player"
	"github.com/luisfcofv/indexter/templates"
)

func CreateWorld() {
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
	timeStart := time.Now()
	timeEnd := timeStart.AddDate(0, 0, 1)
	timeRange := player.TimeRange{timeStart.Unix(), timeEnd.Unix()}
	knowledge := player.Knowledge{
		Social:    []int{1},
		Locations: []int{1},
		Goals:     []int{1},
		Times:     []player.TimeRange{timeRange},
	}

	return player.Player{
		Name:      "Protagonist",
		Knowledge: knowledge,
	}
}

func createGoals() []models.Goal {
	goal1 := models.Goal{1, "Goal 1", "Talk to the king"}
	goal2 := models.Goal{2, "Goal 2", "Find the treasure"}
	goal3 := models.Goal{3, "Goal 3", "Collect coins"}
	goal4 := models.Goal{4, "Goal 4", "Fight the dragon"}
	goal5 := models.Goal{5, "Goal 5", "Rescue the princess"}

	return []models.Goal{
		goal1,
		goal2,
		goal3,
		goal4,
		goal5,
	}
}

func createLocations() []models.Location {
	location1 := models.Location{1, "City 1", "Description 1", nil}
	location2 := models.Location{2, "City 2", "Description 2", nil}
	location3 := models.Location{3, "City 3", "Description 3", nil}
	location4 := models.Location{4, "City 4", "Description 4", nil}
	location5 := models.Location{5, "City 5", "Description 5", nil}

	location1.Connect(&location2)
	location2.Connect(&location4)
	location2.Connect(&location5)
	location3.Connect(&location5)
	location4.Connect(&location5)

	return []models.Location{
		location1,
		location2,
		location3,
		location4,
		location5,
	}
}

func createAgents() []models.Agent {
	agent1 := models.Agent{1, "Agent 1", "The king", nil}
	agent2 := models.Agent{2, "Agent 2", "Witness", nil}
	agent3 := models.Agent{3, "Agent 3", "Wizard", nil}
	agent4 := models.Agent{4, "Agent 4", "The Queen", nil}
	agent5 := models.Agent{5, "Agent 5", "Traveler ", nil}

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
