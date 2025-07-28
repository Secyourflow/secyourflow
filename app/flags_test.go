package app_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/lagzenthakuri/secyourflow/app"
	secyourflowTesting "github.com/lagzenthakuri/secyourflow/testing"
)

func TestHasFlag(t *testing.T) {
	t.Parallel()

	secyourflowApp, _, cleanup := secyourflowTesting.App(t)
	defer cleanup()

	// stage 1
	assert.False(t, app.HasFlag(secyourflowApp, "test"))

	// stage 2
	require.NoError(t, app.SetFlags(secyourflowApp, []string{"test"}))
	assert.True(t, app.HasFlag(secyourflowApp, "test"))
}

func Test_flags(t *testing.T) {
	t.Parallel()

	secyourflowApp, _, cleanup := secyourflowTesting.App(t)
	defer cleanup()

	got, err := app.Flags(secyourflowApp)
	require.NoError(t, err)

	want := []string{}
	assert.ElementsMatch(t, want, got)
}

func Test_setFlags(t *testing.T) {
	t.Parallel()

	secyourflowApp, _, cleanup := secyourflowTesting.App(t)
	defer cleanup()

	// stage 1
	require.NoError(t, app.SetFlags(secyourflowApp, []string{"test"}))

	got, err := app.Flags(secyourflowApp)
	require.NoError(t, err)

	assert.ElementsMatch(t, []string{"test"}, got)

	// stage 2
	require.NoError(t, app.SetFlags(secyourflowApp, []string{"test2"}))

	got, err = app.Flags(secyourflowApp)
	require.NoError(t, err)

	assert.ElementsMatch(t, []string{"test2"}, got)

	// stage 3
	require.NoError(t, app.SetFlags(secyourflowApp, []string{"test", "test2"}))

	got, err = app.Flags(secyourflowApp)
	require.NoError(t, err)

	assert.ElementsMatch(t, []string{"test", "test2"}, got)
}
