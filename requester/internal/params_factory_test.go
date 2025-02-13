package internal

import (
	"flag"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupFlagTest(args []string) func() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	oldArgs := os.Args
	os.Args = args
	return func() {
		os.Args = oldArgs
	}
}

func TestDefaultValuesNewParamsFromFlags(t *testing.T) {
	cleanup := setupFlagTest([]string{"requester"})
	defer cleanup()

	params := NewParamsFromFlags()

	assert.Equal(t, 10, params.NumRequests)
	assert.Equal(t, "http://localhost:3010", params.Host)
}

func TestCustomValuesNewParamsFromFlags(t *testing.T) {
	randomParams := ParamsMotherRamdon()
	cleanup := setupFlagTest([]string{"requester", "-n", strconv.Itoa(randomParams.NumRequests), "-h", randomParams.Host})
	defer cleanup()

	params := NewParamsFromFlags()

	assert.Equal(t, randomParams.NumRequests, params.NumRequests)
	assert.Equal(t, randomParams.Host, params.Host)
}
