package env_test

import (
	"github.com/cksidharthan/env-test/env"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewOne(t *testing.T) {
	t.Parallel()
	_, err := env.New()
	assert.NoError(t, err)
}

func TestNewTwo(t *testing.T) {
	t.Parallel()
	_, err := env.New()
	assert.NoError(t, err)
}
