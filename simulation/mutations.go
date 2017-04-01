package simulation

import (
	"github.com/graphql-go/graphql"

	"github.com/luisfcofv/indexter/templates"
)

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"generateEvents": templates.GenerateEventsField,
	},
})
