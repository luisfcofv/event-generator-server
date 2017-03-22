package models

import "github.com/graphql-go/graphql"

type Agent struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    int    `json:"location"`
	Connections []int  `json:"connections"`
}

func (agent *Agent) Connect(otherAgent *Agent) {
	agent.Connections = append(agent.Connections, otherAgent.ID)
	otherAgent.Connections = append(otherAgent.Connections, agent.ID)
}

var AgentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Agent",
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
			"connections": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"location": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
