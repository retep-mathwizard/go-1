package load

import (
	"encoding/json"
	"io/ioutil"
)

type Race struct {
	Name    string `json:"Name"`
	About   string `json:"About"`
	Health  int    `json:"Health"`
	Attacks []Move `json:"Attacks"`
}

type Move struct {
	Name     string `json:"Name"`
	Info     string `json:"Info"`
	Damage   int    `json:"Damage"`
	Accuracy int    `json:"Accuracy"`
}

func LoadClass(path string) *Race {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	jsonClass := &Race{}
	//err = json.NewDecoder(bytes.Body).Decode(jsonClass)
	err = json.Unmarshal(bytes, jsonClass)
	//decodes it
	if err != nil {
		panic(err)
	}
	return jsonClass
}
