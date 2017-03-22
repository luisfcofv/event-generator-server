package player

import "github.com/graphql-go/graphql"

type knowledge struct {
	Space  []int  `json:"space"`
	Goals  []int  `json:"goals"`
	Social []int  `json:"social"`
	Time   string `json:"time"`
}

var knowledgeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Knowledge",
		Fields: graphql.Fields{
			"space": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"goals": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"social": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"time": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
