package models

import "github.com/graphql-go/graphql"

type Goal struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var GoalType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Goal",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
