package world

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/graphql-go/graphql"
	"github.com/luisfcofv/indexter/player"
)

type world struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	State     interface{}   `json:"state"`
	Player    player.Player `json:"player"`
	Locations []Location    `json:"locations"`
}

var data map[string]world

var worldType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "World",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"player": &graphql.Field{
				Type: player.PlayerType,
			},
			"state": &graphql.Field{
				Type: graphql.String,
			},
			"locations": &graphql.Field{
				Type: graphql.NewList(LocationType),
			},
		},
	},
)

var WorlField = &graphql.Field{
	Type: worldType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		_ = importJSONDataFromFile("data.json", &data)
		idQuery, isOK := p.Args["id"].(string)
		if isOK {
			return data[idQuery], nil
		}
		return nil, nil
	},
}

//Helper function to import json from file to map
func importJSONDataFromFile(fileName string, result interface{}) (isOK bool) {
	isOK = true
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print("Error:", err)
		isOK = false
	}
	err = json.Unmarshal(content, result)
	if err != nil {
		isOK = false
		fmt.Print("Error:", err)
	}
	return
}