package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetClientConfig(t *testing.T) {
	got := GetClientConfig()
	assert.NotNil(t, got, "TestGetClientConfig")
}
