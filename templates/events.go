package templates

import (
	"math/rand"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/luisfcofv/indexter/models"
)

func getEventTemplates(world models.World) []models.Event {
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

	location := rand.Intn(len(world.Locations))
	agent := rand.Intn(len(world.Agents))
	goal := rand.Intn(len(world.Goals))

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

var EventTemplatesField = &graphql.Field{
	Type: graphql.NewList(models.EventType),
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		name, isOK := p.Args["name"].(string)
		if isOK {
			world := models.GetWorld(name)
			return getEventTemplates(world), nil
		}
		return nil, nil
	},
}
