package models

import "github.com/graphql-go/graphql"

type Agent struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Connections []int  `json:"connections"`
}

var AgentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Location",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"connections": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)
