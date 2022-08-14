package testing

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunInTest(t *testing.T) {
	assert.True(t, RunInTest())
}

func TestCreateEnv(t *testing.T) {
	err := CreateEnv()
	assert.Nil(t, err)
	assert.FileExists(t, ".env")

	err = os.Remove(".env")
	assert.Nil(t, err)
}
