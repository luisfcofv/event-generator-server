package world

import (
	"github.com/graphql-go/graphql"
	"github.com/luisfcofv/indexter/player"
)

type world struct {
	ID     string        `json:"id"`
	Name   string        `json:"name"`
	State  interface{}   `json:"state"`
	Player player.Player `json:"player"`
}

var worldType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "World",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"player": &graphql.Field{
				Type: player.PlayerType,
			},
			"state": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var WorlField = &graphql.Field{
	Type: worldType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idQuery := p.Args["id"].(string)
		state := make(map[string]string)
		state["location"] = "test"
		state["key"] = "value"
		// player := player.Player{"1", "luis", nil}
		player := player.ResolvePlayer(idQuery)
		println(idQuery)
		myworld := world{"1", "World id: " + idQuery, state, player}
		return myworld, nil
	},
}
