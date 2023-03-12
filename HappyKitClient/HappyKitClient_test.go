package HappyKitClient

import (
	"testing"
)

func TestIsEnabled(t *testing.T) {
	Initialize("123")
	flagKey := "testKey"
	flagValue := true
	featureFlags := FeatureFlags{
		flagKey: flagValue,
	}
	flagsCache.Set(flagsCacheKey, featureFlags)

	got := IsEnabled(flagKey)
	want := flagValue

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}

	flagsCache.Purge()
}

func TestIsEnabled_WhenDefaultValueIsSet_FlagExists(t *testing.T) {
	Initialize("123")
	flagKey := "testKey"
	flagValue := false
	defaultValue := true
	featureFlags := FeatureFlags{
		flagKey: flagValue,
	}
	flagsCache.Set(flagsCacheKey, featureFlags)

	got := IsEnabled(flagKey, defaultValue)
	want := flagValue

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}

	flagsCache.Purge()
}

func TestIsEnabled_WhenDefaultValueIsSet_FlagDoesNotExist(t *testing.T) {
	featureFlagKey := "testKey123"
	defaultValue := true

	got := IsEnabled(featureFlagKey, defaultValue)
	want := defaultValue

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestIsEnabled_WhenDefaultValueNotSet_FlagDoesNotExist(t *testing.T) {
	featureFlagKey := "testKey"

	got := IsEnabled(featureFlagKey)
	want := false

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestIsEnabled_WhenBadNonBooleanFlagUsed_DefaultValueReturned(t *testing.T) {
	Initialize("123")
	flagKey := "testKey"
	flagValue := "true"
	featureFlags := FeatureFlags{
		flagKey: flagValue,
	}
	flagsCache.Set(flagsCacheKey, featureFlags)

	got := IsEnabled(flagKey)
	want := false

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}

	flagsCache.Purge()
}

func TestIsEnabled_WhenBadNonBooleanFlagUsed_GivenDefaultValueReturned(t *testing.T) {
	Initialize("123")
	flagKey := "testKey"
	flagValue := "true"
	featureFlags := FeatureFlags{
		flagKey: flagValue,
	}
	flagsCache.Set(flagsCacheKey, featureFlags)

	got := IsEnabled(flagKey, true)
	want := true

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}

	flagsCache.Purge()
}
