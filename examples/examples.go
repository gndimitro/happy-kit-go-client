package main

import (
	"fmt"

	client "github.com/gndimitro/happykit-go-client"
)

func main() {
	client.Initialize("flags_pub_development_XXXXXXX")

	if client.IsEnabled("bool") {
		fmt.Println("Bool is enabled")
	} else {
		fmt.Println("Bool is disabled")
	}

	if client.IsEnabled("string") {
		fmt.Println("String is enabled")
	}

	user := client.User{Key: "userKey"}
	if client.IsEnabledForUser("user_bool", user) {
		fmt.Println("User bool is enabled")
	}

	traits := struct {
		Trait string `json:"trait"`
	}{Trait: "testing"}
	if client.IsEnabledForTraits("trait_flag", traits) {
		fmt.Println("Traits bool is enabled")
	}
}
