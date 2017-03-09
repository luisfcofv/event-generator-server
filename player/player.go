package player

import "github.com/graphql-go/graphql"

type Player struct {
	Name      string    `json:"name"`
	Knowledge knowledge `json:"knowledge"`
}

var PlayerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Player",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"knowledge": &graphql.Field{
				Type: knowledgeType,
			},
		},
	},
)
