package player

import "github.com/graphql-go/graphql"

type knowledge struct {
	Space       string `json:"space"`
	Causation   string `json:"causation"`
	Time        string `json:"time"`
	Intention   string `json:"intention"`
	Protagonist string `json:"protagonist"`
}

var knowledgeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Knowledge",
		Fields: graphql.Fields{
			"space": &graphql.Field{
				Type: graphql.String,
			},
			"causation": &graphql.Field{
				Type: graphql.String,
			},
			"time": &graphql.Field{
				Type: graphql.String,
			},
			"intention": &graphql.Field{
				Type: graphql.String,
			},
			"protagonist": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
