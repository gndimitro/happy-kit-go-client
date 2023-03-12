package HappyKitClient

import (
	"fmt"
)

func main() {
	Initialize("flags_pub_development_XXXXXXX")

	if IsEnabled("bool") {
		fmt.Println("Bool is enabled")
	} else {
		fmt.Println("Bool is disabled")
	}

	if IsEnabled("string") {
		fmt.Println("String is enabled")
	}

	user := User{Key: "userKey"}
	if IsEnabledForUser("user_bool", user) {
		fmt.Println("User bool is enabled")
	}

	traits := struct {
		Trait string `json:"trait"`
	}{Trait: "testing"}
	if IsEnabledForTraits("trait_flag", traits) {
		fmt.Println("Traits bool is enabled")
	}
}
