package main

import (
	"fmt"

	"github.com/gndimitro/happykit-go-client/HappyKitClient"
)

func main() {
	HappyKitClient.Initialize("flags_pub_development_XXXXXXX")

	if HappyKitClient.IsEnabled("bool") {
		fmt.Println("Bool is enabled")
	} else {
		fmt.Println("Bool is disabled")
	}

	if HappyKitClient.IsEnabled("string") {
		fmt.Println("String is enabled")
	}

	user := HappyKitClient.User{Key: "userKey"}
	if HappyKitClient.IsEnabledForUser("user_bool", user) {
		fmt.Println("User bool is enabled")
	}

	traits := struct {
		Trait string `json:"trait"`
	}{Trait: "testing"}
	if HappyKitClient.IsEnabledForTraits("trait_flag", traits) {
		fmt.Println("Traits bool is enabled")
	}
}
