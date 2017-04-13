package models

import "github.com/graphql-go/graphql"

type Event struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Location    Location `json:"location"`
	Protagonist bool     `json:"protagonist"`
	Agents      []Agent  `json:"agents"`
	Goal        Goal     `json:"goal"`
	Cause       Goal     `json:"cause"`
	Propagation int      `json:"propagation"`
	Salience    Salience `json:"salience"`
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
				Type: LocationType,
			},
			"protagonist": &graphql.Field{
				Type: graphql.Boolean,
			},
			"agents": &graphql.Field{
				Type: graphql.NewList(AgentType),
			},
			"goal": &graphql.Field{
				Type: GoalType,
			},
			"cause": &graphql.Field{
				Type: GoalType,
			},
			"propagation": &graphql.Field{
				Type: graphql.Int,
			},
			"salience": &graphql.Field{
				Type: SalienceType,
			},
		},
	},
)
