package simulation

import (
	"github.com/graphql-go/graphql"

	"github.com/luisfcofv/indexter/models"
)

var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"world": models.WorldField,
		},
	},
)
