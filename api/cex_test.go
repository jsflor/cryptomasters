package api_test

import (
	"jsflor/cryptomasters/api"
	"testing"
)

func TestGetRate(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Error("Error should not be nil")
	}
}
