package player

import "github.com/graphql-go/graphql"

type Player struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Knowledge knowledge `json:"knowledge"`
}

var PlayerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Player",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"knowledge": &graphql.Field{
				Type: knowledgeType,
			},
		},
	},
)

func ResolvePlayer(id string) Player {
	player := Player{"id: " + id, "luis", knowledge{"0", "0", "0", "0", "0"}}
	return player
}
