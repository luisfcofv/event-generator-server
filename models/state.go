package models

import (
	"github.com/graphql-go/graphql"
)

type PlayerState struct {
	Location int `json:"location"`
	Goal     int `json:"goal"`
}

var playerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Player",
		Fields: graphql.Fields{
			"location": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

type State struct {
	Player PlayerState `json:"player"`
}

var StateType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "State",
		Fields: graphql.Fields{
			"player": &graphql.Field{
				Type: playerType,
			},
		},
	},
)
