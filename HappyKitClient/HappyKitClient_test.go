package HappyKitClient

import (
	"testing"
)

func TestIsEnabled(t *testing.T) {
	Initialize("123")
	featureFlagKey := "testKey"
	featureFlags := FeatureFlags{
		featureFlagKey: true,
	}
	flagsCache.Set(flagsCacheKey, featureFlags)

	got := IsEnabled(featureFlagKey)
	want := true

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}

	flagsCache.Purge()
}

func TestIsEnabled_WhenDefaultValueIsSet_FlagExists(t *testing.T) {
	Initialize("123")
	featureFlagKey := "testKey"
	defaultValue := true
	featureFlags := FeatureFlags{
		featureFlagKey: false,
	}
	flagsCache.Set(flagsCacheKey, featureFlags)

	got := IsEnabled(featureFlagKey, defaultValue)
	want := false

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}

	featureFlags = nil
}

func TestIsEnabled_WhenDefaultValueIsSet_FlagDoesNotExist(t *testing.T) {
	// BROKEN, NEEDS FIXING
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
