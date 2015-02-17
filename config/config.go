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

func Validate(projects []Project) error {
	for _, project := range projects {
		if !helper.FileExists(project.Dir) {
			return errors.New(fmt.Sprintf("Project directory %v does not exist!", project.Dir))
		}

		if !helper.IsWritable(project.Log) {
			return errors.New(fmt.Sprintf("Project %v log file %v is not writable", project.Log))
		}
	}

	return nil
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

	if err := Validate(projects); err != nil {
		return nil, err
	}

	return projects, nil
}
