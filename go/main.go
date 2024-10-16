package main

import (
	"html/template"
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

func getUsers() *[]User {
	listUsers := &[]User{}

	yamlFile, err := os.ReadFile("../config_users.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, listUsers)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return listUsers
}

func main() {
	users := getUsers()

	usersTFFile, _ := os.Create("../terraform/users.tf")

	tmpl, _ := template.ParseFiles("templates/users.tf.tmpl")

	// Execute the template for each recipient.
	for _, user := range *users {
		err := tmpl.Execute(usersTFFile, user)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

	defer usersTFFile.Close()
}
