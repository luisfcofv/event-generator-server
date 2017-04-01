package player

import "github.com/graphql-go/graphql"

type Knowledge struct {
	Locations []int `json:"locations"`
	Goals     []int `json:"goals"`
	Social    []int `json:"social"`
}

var KnowledgeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Knowledge",
		Fields: graphql.Fields{
			"locations": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"goals": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"social": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var KnowledgeInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "KnowledgeInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"locations": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.Int),
		},
		"goals": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.Int),
		},
		"social": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.Int),
		},
	},
},
)
