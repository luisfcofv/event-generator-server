package models

import "github.com/graphql-go/graphql"

type Neighbor struct {
	ID       string  `json:"id"`
	Distance float32 `json:"distance"`
}

var neighborType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Neighbor",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"distance": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

type Location struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Neighbors   []Neighbor `json:"neighbors"`
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
				Type: graphql.NewList(neighborType),
			},
		},
	},
)
