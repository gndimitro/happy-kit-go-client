package HappyKitClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type FeatureFlags map[string]bool
type RequestResponseBody struct {
	Flags FeatureFlags
}

var featureFlags FeatureFlags

func IsEnabled(featureFlagKey string, defaultValueOptional ...bool) bool {
	defaultValue := false

	if len(defaultValueOptional) > 0 {
		defaultValue = defaultValueOptional[0]
	}

	if val, ok := featureFlags[featureFlagKey]; ok {
		return val
	}

	return defaultValue
}

func FetchFeatureFlags(flagsKey string) bool {
	response, err := http.Post(fmt.Sprintf("https://happykit.dev/api/flags/%s", flagsKey), "application/json", bytes.NewReader([]byte("")))

	if err != nil {
		fmt.Print(err.Error())
		return false
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return false
	}

	var requestResponseBody RequestResponseBody

	err = json.Unmarshal(responseData, &requestResponseBody)
	if err != nil {
		log.Fatal(err)
	}

	featureFlags = requestResponseBody.Flags

	return true
}
