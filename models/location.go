package models

import "github.com/graphql-go/graphql"

type Neighbor struct {
	ID   int `json:"id"`
	Time int `json:"time"`
}

var NeighborType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Neighbor",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"time": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

type Location struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Neighbors   []Neighbor `json:"neighbors"`
}

func (location *Location) Connect(otherLocation *Location, time int) {
	location.Neighbors = append(location.Neighbors, Neighbor{otherLocation.ID, time})
	otherLocation.Neighbors = append(otherLocation.Neighbors, Neighbor{location.ID, time})
}

var LocationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Location",
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
			"neighbors": &graphql.Field{
				Type: graphql.NewList(NeighborType),
			},
		},
	},
)
