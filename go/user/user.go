package user

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Permissions struct {
	Effect    string   `yaml:"effect"`
	Actions   []string `yaml:"actions"`
	Resources []string `yaml:"resources"`
}

type User struct {
	Name        string              `yaml:"name"`
	Description string              `yaml:"description"`
	Tags        []map[string]string `yaml:"tags"`
	Permissions []Permissions       `yaml:"permissions"`
}

func GetUsers() *[]User {
	listUsers := &[]User{}

	yamlFile, err := os.ReadFile("../config_users.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, listUsers)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return listUsers
}
