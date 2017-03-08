package models

import "github.com/graphql-go/graphql"

type Neighbor struct {
	ID       int32 `json:"id"`
	Distance int32 `json:"distance"`
}

var neighborType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Neighbor",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"distance": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

type Location struct {
	ID          int32      `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Neighbors   []Neighbor `json:"neighbors"`
}

func (location *Location) Connect(otherLocation *Location, distance int32) {
	location.Neighbors = append(location.Neighbors, Neighbor{otherLocation.ID, distance})
	otherLocation.Neighbors = append(otherLocation.Neighbors, Neighbor{location.ID, distance})
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
