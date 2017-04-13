package templates

import (
	"math/rand"
	"time"

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

	location := world.Locations[rand.Intn(len(world.Locations))]
	agent := world.Agents[rand.Intn(len(world.Agents))]
	goal := world.Goals[rand.Intn(len(world.Goals))]
	causation := world.Goals[rand.Intn(len(world.Goals))]

	protagonist := false
	if rand.Intn(2) == 1 {
		protagonist = true
	}

	event := models.Event{
		Name:        "Template 1",
		Description: "Random template",
		Protagonist: protagonist,
		Agents:      []models.Agent{agent},
		Location:    location,
		Goal:        goal,
		Propagation: rand.Intn(2) + 1,
		Cause:       causation,
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
		Protagonist: false,
		Agents:      []models.Agent{world.Agents[possibleAgents[agentID]]},
		Location:    world.Locations[3],
		Goal:        world.Goals[1], // Find the treasure
		Propagation: 1,
		Cause:       world.Goals[2], // Collect coins
	}

	return event
}

func getThirdTemplate(world models.World) models.Event {
	event := models.Event{
		Name:        "Template 3",
		Description: "City 2, The king, Talk to the king",
		Protagonist: false,
		Agents:      []models.Agent{world.Agents[0], world.Agents[1]}, // The king
		Location:    world.Locations[2],
		Goal:        world.Goals[0], // Talk to the king
		Propagation: 1,
		Cause:       world.Goals[4], // Rescue the princess
	}

	return event
}

func getFourthTemplate(world models.World) models.Event {
	event := models.Event{
		Name:        "Template 4",
		Description: "City 5, Protagonist, Fight the dragon",
		Protagonist: true,
		Agents:      []models.Agent{world.Agents[1]},
		Location:    world.Locations[4],
		Goal:        world.Goals[3], // Fight the dragon
		Propagation: 1,
	}

	return event
}

func getFifthTemplate(world models.World) models.Event {
	time := time.Now().Unix()
	rand.Seed(time)

	possibleCities := []int{3, 4}
	cityID := rand.Intn(len(possibleCities))

	event := models.Event{
		Name:        "Template 5",
		Description: "City 3 or 4, Wizard, Rescue the princess",
		Agents:      []models.Agent{world.Agents[2]}, // Wizard
		Location:    world.Locations[possibleCities[cityID]],
		Goal:        world.Goals[4], // Rescue the princess
		Propagation: 2,
		Cause:       world.Goals[3], // Fight the dragon
	}

	return event
}
