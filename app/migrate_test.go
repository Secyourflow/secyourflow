package app_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lagzenthakuri/secyourflow/app"
	"github.com/lagzenthakuri/secyourflow/migrations"
	secyourflowTesting "github.com/lagzenthakuri/secyourflow/testing"
)

func Test_MigrateDBsDown(t *testing.T) {
	t.Parallel()

	secyourflowApp, _, cleanup := secyourflowTesting.App(t)
	defer cleanup()

	_, err := secyourflowApp.Dao().FindCollectionByNameOrId(migrations.ReactionCollectionName)
	require.NoError(t, err)

	require.NoError(t, app.MigrateDBsDown(secyourflowApp))
}
