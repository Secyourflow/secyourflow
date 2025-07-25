package fakedata_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	secyourflowTesting "github.com/lagzenthakuri/secyourflow/app"
	"github.com/lagzenthakuri/secyourflow/fakedata"
)

func TestDefaultData(t *testing.T) {
	t.Parallel()

	app, err := secyourflowTesting.App("", true)
	require.NoError(t, err)

	require.NoError(t, fakedata.GenerateDefaultData(app))

	require.NoError(t, fakedata.ValidateDefaultData(app))
}
