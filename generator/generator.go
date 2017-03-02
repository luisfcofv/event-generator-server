package generator

import (
	"encoding/json"
	"fmt"
	"github.com/luisfcofv/indexter/generator/salience"
	"github.com/luisfcofv/indexter/models"
	"io/ioutil"
)

type generator struct {
	Locations []models.Location `json:"locations"`
	Social    []models.Agent    `json:"social"`
}

var data generator

func Compute() {
	_ = importJSONDataFromFile("data.json", &data)

	activeNodes := []string{"2", "7"}
	experienceNodes := []string{"1", "4"}
	salience.SpaceSalience(data.Locations, activeNodes, experienceNodes)
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
