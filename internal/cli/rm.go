package cli

import (
	"github.com/imrancluster/sshmama/internal/db"
	"github.com/spf13/cobra"
)

func cmdRM() *cobra.Command {
	c := &cobra.Command{
		Use:   "rm <name>",
		Short: "Remove an entry (and its attachment if present)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return db.DeleteEntry(appState, args[0])
		},
	}
	return c
}
