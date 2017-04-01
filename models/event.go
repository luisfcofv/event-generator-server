package models

import "github.com/graphql-go/graphql"

type Event struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    int    `json:"location"`
	Agents      []int  `json:"agents"`
	// https://github.com/galaxykate/tracery
	// Generate the events with this tool
	// Fill out the missing properties with randomly properties
	Goal     int      `json:"goal"`
	Cause    int      `json:"cause"`
	Time     int64    `json:"time"`
	Salience Salience `json:"salience"`
}

var EventType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Event",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"location": &graphql.Field{
				Type: graphql.Int,
			},
			"agents": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"goal": &graphql.Field{
				Type: graphql.Int,
			},
			"time": &graphql.Field{
				Type: graphql.Int,
			},
			"salience": &graphql.Field{
				Type: SalienceType,
			},
		},
	},
)
