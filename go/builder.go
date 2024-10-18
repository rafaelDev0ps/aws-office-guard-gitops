// go:build builder
package main

import (
	"html/template"
	"log"
	user "office-guard/user"
	"os"
)

func main() {
	users := user.GetUsers()

	usersTFFile, _ := os.Create("../terraform/users.tf")

	tmpl, _ := template.ParseFiles("templates/users.tf.tmpl")

	for _, user := range *users {
		err := tmpl.Execute(usersTFFile, user)
		if err != nil {
			log.Println("Executing template:", err)
		}
	}

	defer usersTFFile.Close()
}
