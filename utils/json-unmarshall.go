package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/3n0ugh/JSON-to-MongoDB/config"
	"github.com/3n0ugh/JSON-to-MongoDB/models"
)

var BookList *models.Books

func GetDataFromJSON() error {
	config, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	dat, err := os.ReadFile(config.JsonPath)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	err = json.Unmarshal(dat, &BookList)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}
