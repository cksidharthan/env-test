package env_test

import (
	"github.com/cksidharthan/env-test/env"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWorkingNewOne(t *testing.T) {
	t.Parallel()
	_, err := env.WorkingNew()
	assert.NoError(t, err)
}

func TestWorkingNewTwo(t *testing.T) {
	t.Parallel()
	_, err := env.WorkingNew()
	assert.NoError(t, err)
}
