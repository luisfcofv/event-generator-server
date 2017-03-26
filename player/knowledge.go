package player

import "github.com/graphql-go/graphql"

type TimeRange struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

var timeRangeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TimeRange",
		Fields: graphql.Fields{
			"start": &graphql.Field{
				Type: graphql.Int,
			},
			"end": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

type Knowledge struct {
	Locations []int       `json:"locations"`
	Goals     []int       `json:"goals"`
	Social    []int       `json:"social"`
	Times     []TimeRange `json:"times"`
}

var knowledgeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Knowledge",
		Fields: graphql.Fields{
			"locations": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"goals": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"social": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"times": &graphql.Field{
				Type: graphql.NewList(timeRangeType),
			},
		},
	},
)
