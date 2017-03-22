package models

import "github.com/graphql-go/graphql"

type Location struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Neighbors   []int  `json:"neighbors"`
}

func (location *Location) Connect(otherLocation *Location) {
	location.Neighbors = append(location.Neighbors, otherLocation.ID)
	otherLocation.Neighbors = append(otherLocation.Neighbors, location.ID)
}

var LocationType = graphql.NewObject(
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
			"neighbors": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)
