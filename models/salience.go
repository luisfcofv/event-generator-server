package models

import (
	"github.com/graphql-go/graphql"
)

type Salience struct {
	Space     float64 `json:"space"`
	Social    float64 `json:"social"`
	Causation float64 `json:"causation"`
	Intention float64 `json:"intention"`
	Time      float64 `json:"time"`
	Total     float64 `json:"total"`
}

var SalienceType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Salience",
		Fields: graphql.Fields{
			"space": &graphql.Field{
				Type: graphql.Float,
			},
			"social": &graphql.Field{
				Type: graphql.Float,
			},
			"causation": &graphql.Field{
				Type: graphql.Float,
			},
			"intention": &graphql.Field{
				Type: graphql.Float,
			},
			"time": &graphql.Field{
				Type: graphql.Float,
			},
			"total": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)
