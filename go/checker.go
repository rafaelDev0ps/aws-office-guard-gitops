// go:build checker
package main

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"strings"

	user "office-guard/user"
)

func checkDuplicateNames(users []user.User) error {
	listNames := []string{}
	for _, user := range users {
		listNames = append(listNames, user.Name)
	}

	slices.Sort(listNames)
	dupNames := []string{}
	for i := 0; i < len(listNames)-1; i++ {
		if listNames[i] == listNames[i+1] {
			dupNames = append(dupNames, listNames[i])
		}
	}

	if len(dupNames) > 0 {
		formatedDups := strings.Trim(fmt.Sprint(dupNames), "[]")
		return errors.New("The following names are duplicated: " + formatedDups)
	}

	return nil
}

func checkEffectSintax(users []user.User) error {
	for _, user := range users {
		for _, perm := range user.Permissions {
			if perm.Effect != "Allow" && perm.Effect != "Deny" {
				return errors.New("The Effect sintax is wrong. Only 'Allow' or 'Deny' values are accepted.")
			}
		}
	}

	return nil
}

func main() {
	users := user.GetUsers()

	err := checkDuplicateNames(*users)
	if err != nil {
		log.Fatal(err)
	}

	err = checkEffectSintax(*users)
	if err != nil {
		log.Fatal(err)
	}
}
