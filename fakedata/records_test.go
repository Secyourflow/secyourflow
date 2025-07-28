package fakedata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/lagzenthakuri/secyourflow/fakedata"
	secyourflowTesting "github.com/lagzenthakuri/secyourflow/testing"
)

func Test_records(t *testing.T) {
	t.Parallel()

	app, _, cleanup := secyourflowTesting.App(t)
	defer cleanup()

	got, err := fakedata.Records(app, 2, 2)
	require.NoError(t, err)

	assert.Greater(t, len(got), 2)
}

func TestGenerate(t *testing.T) {
	t.Parallel()

	app, _, cleanup := secyourflowTesting.App(t)
	defer cleanup()

	err := fakedata.Generate(app, 0, 0)
	require.NoError(t, err)
}
