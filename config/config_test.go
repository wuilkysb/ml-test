package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"sync"
	"testing"
)

func TestEnvironments_WhenAllVarsAreSet(t *testing.T) {
	Environments()

	assert.Equal(t, c.XApplicationId, "local")
}

func TestEnvironments_WhenSomeVarIsMissed(t *testing.T) {
	expectedError := "Error parsing environment vars &errors.errorString{s:\"required key SERVER_HOST missing value\"}"
	once = sync.Once{}
	os.Clearenv()

	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, expectedError, r)
		}
	}()

	Environments()
}
