package HappyKitClient

import (
	"testing"
)

func TestIsEnabled(t *testing.T) {
	featureFlagKey := "testKey"
	featureFlags = map[string]bool {
		featureFlagKey: true,
	}
	
	got := IsEnabled(featureFlagKey)
	want := true

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}

	featureFlags = nil
}

func TestIsEnabled_WhenDefaultValueIsSet_FlagExists(t *testing.T) {
	featureFlagKey := "testKey"
	defaultValue := true
	featureFlags = map[string]bool {
		featureFlagKey: false,
	}
	
	got := IsEnabled(featureFlagKey, defaultValue)
	want := false

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}

	featureFlags = nil
}

func TestIsEnabled_WhenDefaultValueIsSet_FlagDoesNotExist(t *testing.T) {
	featureFlagKey := "testKey"
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