package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	//	"log"

	helper "github.com/dselans/goroq/helper"
)

type Project struct {
	Name string `json:"ProjectName"`
	Dir  string `json:"ProjectDir"`
	Log  string `json:"ProjectLog"`
}

func Read(file string) ([]Project, error) {
	if !helper.FileExists(file) {
		return nil, errors.New(fmt.Sprintf("Config file '%v' does not exist!", file))
	}

	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	projects := make([]Project, 1)
	if err := json.Unmarshal(contents, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}
