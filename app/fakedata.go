package app

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/spf13/cobra"

	"github.com/lagzenthakuri/secyourflow/fakedata"
)

func fakeDataCmd(app core.App) *cobra.Command {
	var userCount, ticketCount int

	cmd := &cobra.Command{
		Use: "fake-data",
		RunE: func(_ *cobra.Command, _ []string) error {
			return fakedata.Generate(app, userCount, ticketCount)
		},
	}

	cmd.PersistentFlags().IntVar(&userCount, "users", 10, "Number of users to generate")

	cmd.PersistentFlags().IntVar(&ticketCount, "tickets", 100, "Number of tickets to generate")

	return cmd
}

func defaultDataCmd(app core.App) *cobra.Command {
	cmd := &cobra.Command{
		Use: "default-data",
		RunE: func(_ *cobra.Command, _ []string) error {
			return fakedata.GenerateDefaultData(app)
		},
	}

	return cmd
}
