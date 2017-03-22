package models

import (
	"github.com/graphql-go/graphql"
)

type Salience struct {
	Space     float32 `json:"space"`
	Social    float32 `json:"social"`
	Causation float32 `json:"causation"`
	Intention float32 `json:"intention"`
	Time      float32 `json:"time"`
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
		},
	},
)
