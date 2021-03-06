package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	var mapa = map[string]map[string]string{
		"Category": {
			"Products []*Product `json:\"products\"`": "Products []*Product `json:\"products\" gorm:\"many2many:product_categories;\"`",
		},
		"Country": {
			"ID string `json:\"id\"`":                 "ID string `json:\"id\" gorm:\"primaryKey\"`",
			"Products []*Product `json:\"products\"`": "Products []*Product `json:\"products\" gorm:\"many2many:product_countries;\"`",
		},
		"Product": {
			"Countries []*Country `json:\"countries\"`":    "Countries []*Country `json:\"countries\" gorm:\"many2many:product_countries;\"`",
			"Categories []*Category `json:\"categories\"`": "Categories []*Category `json:\"categories\" gorm:\"many2many:product_categories;\"`",
		},
		"ProductAdmin": {
			"Countries []*Country `json:\"countries\"`":    "Countries []*Country `json:\"countries\" gorm:\"many2many:product_countries;\"`",
			"Categories []*Category `json:\"categories\"`": "Categories []*Category `json:\"categories\" gorm:\"many2many:product_categories;\"`",
		},
	}

	data, err := ioutil.ReadFile("dbmodels/graph/model/models_gen.go")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}

	var newFile string

	lines := strings.Split(string(data), "\n")
	var structName string
	for _, line := range lines {

		line = strings.ReplaceAll(line, "\r", "")
		line = strings.ReplaceAll(line, "\n", "")
		line = strings.ReplaceAll(line, "\t", "")

		if line == "}" {
			structName = ""
			newFile += line + "\n"
			continue
		}

		words := strings.Split(line, " ")

		if structName == "" && len(words) > 0 && words[0] == "type" {
			structName = words[1]
		} else if val, ok := mapa[structName][line]; ok {
			line = val
		}

		newFile += line + "\n"
	}

	err = ioutil.WriteFile("dbmodels/graph/model/models_gen.go", []byte(newFile), 0777)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
}
