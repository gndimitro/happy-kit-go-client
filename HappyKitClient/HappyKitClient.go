package HappyKitClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/bluele/gcache"
)

type FeatureFlags map[string]interface{}
type RequestResponseBody struct {
	Flags FeatureFlags
}

// A user struct used to pass to the API
//
// key (string) (required): Unique key for this user
//
// email (string): Email-Address
//
// name (string): Full name or nickname
//
// avatar (string): URL to users profile picture
//
// country (string): Two-letter uppercase country-code, see ISO 3166-1
type User struct {
	Key     string `json:"key"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`
	Country string `json:"country"`
}

var flagsCache gcache.Cache

const flagsCacheKey = "flags"

// You can find the environment key for each stage in your project's settings on happykit.dev
var environmentKey string

// Initialze function to set and store the environment key for all future calls to HappyKit
func Initialize(envKey string) {
	InitializeCustomCacheExpiry(envKey, time.Minute/2)
}

func InitializeCustomCacheExpiry(envKey string, expirationTime time.Duration) {
	environmentKey = envKey
	flagsCache = gcache.New(1).Expiration(expirationTime).Build()
}

// Returns the boolean value of the provided feature flag key
//
// featureFlagKey (string) Key to be used when fetching the feature flag
//
// defaultValueOptional (bool) Backup value to be used in case the feature flag isn't found in the current dataset
func IsEnabled(featureFlagKey string, defaultValueOptional ...bool) bool {
	var isEnabled interface{}

	if len(defaultValueOptional) > 0 {
		isEnabled = GetFlagValue(featureFlagKey, defaultValueOptional[0])
	} else {
		isEnabled = GetFlagValue(featureFlagKey, false)
	}

	return isEnabled.(bool)
}

// Fetches the flag value using the provided key
//
// featureFlagKey (string) Key to be used when fetching the feature flag
//
// defaultValueOptional (interface{}) Backup value to be used in case the feature flag isn't found in the current dataset, by default this is nil
func GetFlagValue(featureFlagKey string, defaultValueOptional ...interface{}) interface{} {
	var defaultValue interface{} = nil

	if len(defaultValueOptional) > 0 {
		defaultValue = defaultValueOptional[0]
	}

	storedFlags, err := flagsCache.Get(flagsCacheKey)
	if err != nil {
		// Cache miss
		flags, success := FetchFeatureFlags()
		if !success {
			return defaultValue
		} else {
			if val, ok := flags[featureFlagKey]; ok {
				return val
			}
		}
	} else {
		// Cache hit
		flags := storedFlags.(FeatureFlags)
		if val, ok := flags[featureFlagKey]; ok {
			return val
		}
	}

	return defaultValue
}

// Checks if the flag is enabled for a specified user. Use only when the flag expected is a boolean type
//
// featureFlagKey (string) Key to be used when fetching the feature flag
//
// user (User) User object to be used in the fetch
//
// defaultValueOptional (bool) Backup value to be used in case the feature flag isn't found in the current dataset
func IsEnabledForUser(featureFlagKey string, user User, defaultValueOptional ...bool) bool {
	var isEnabled interface{}

	if len(defaultValueOptional) > 0 {
		isEnabled = GetFlagValueForUser(featureFlagKey, user, defaultValueOptional[0])
	} else {
		isEnabled = GetFlagValueForUser(featureFlagKey, user, false)
	}

	return isEnabled.(bool)
}

// Fetches the flag value using the provided key for the specified user. Use when your flag value is of any other type supported by HappyKit. This function does not use a cache.
//
// featureFlagKey (string) Key to be used when fetching the feature flag
//
// defaultValueOptional (interface{}) Backup value to be used in case the feature flag isn't found in the current dataset, by default this is nil
func GetFlagValueForUser(featureFlagKey string, user User, defaultValueOptional ...interface{}) interface{} {
	var defaultValue interface{} = nil

	if len(defaultValueOptional) > 0 {
		defaultValue = defaultValueOptional[0]
	}

	flags, success := FetchFeatureFlagsForUser(user)
	if !success {
		return defaultValue
	} else {
		if val, ok := flags[featureFlagKey]; ok {
			return val
		}
	}

	return defaultValue
}

// Fetches the feature flags from the api without any extra paramaters
func FetchFeatureFlags() (FeatureFlags, bool) {
	response, err := http.Post(fmt.Sprintf("https://happykit.dev/api/flags/%s", environmentKey), "application/json", bytes.NewReader([]byte("")))

	if err != nil {
		fmt.Print(err.Error())
		return nil, false
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return nil, false
	}

	var requestResponseBody RequestResponseBody

	err = json.Unmarshal(responseData, &requestResponseBody)
	if err != nil {
		log.Fatal(err)
	}

	err = flagsCache.Set(flagsCacheKey, requestResponseBody.Flags)
	if err != nil {
		fmt.Println("Debug: Failure saving the feature flags to the cache")
	}

	return requestResponseBody.Flags, true
}

// Fetches the feature flags from the api for the specified user
func FetchFeatureFlagsForUser(user User) (FeatureFlags, bool) {
	postBody, err := json.Marshal(user)
	if err != nil {
		return nil, false
	}
	response, err := http.Post(fmt.Sprintf("https://happykit.dev/api/flags/%s", environmentKey), "application/json", bytes.NewBuffer(postBody))

	if err != nil {
		fmt.Print(err.Error())
		return nil, false
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return nil, false
	}

	var requestResponseBody RequestResponseBody

	err = json.Unmarshal(responseData, &requestResponseBody)
	if err != nil {
		log.Fatal(err)
	}

	return requestResponseBody.Flags, true
}
