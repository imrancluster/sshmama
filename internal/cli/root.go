package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/imrancluster/sshmama/internal/app"
	"github.com/imrancluster/sshmama/pkg/version"
)

var (
	rootCmd  *cobra.Command
	appState *app.App
	dataDir  string
)

func init() {
	rootCmd = &cobra.Command{
		Use:   "sshmama",
		Short: "Manage named SSH connections with optional encrypted keys.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if appState != nil {
				return nil
			}
			if dataDir == "" {
				var err error
				dataDir, err = app.DefaultDataDir()
				if err != nil {
					return err
				}
			}
			a, err := app.New(dataDir)
			if err != nil {
				return err
			}
			appState = a
			return nil
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			// no-op; keep DB open across subcommands run
		},
	}
	rootCmd.PersistentFlags().StringVar(&dataDir, "data-dir", "", "Custom data directory (default: OS config dir)")
	rootCmd.Version = version.Version

	// subcommands
	rootCmd.AddCommand(cmdAdd(), cmdList(), cmdSearch(), cmdConnect(), cmdAttach(), cmdRM(), cmdCompletion(), cmdEdit(), cmdExport(), cmdImport())
}

func Execute() {
	defer func() {
		if appState != nil {
			appState.Close()
		}
	}()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
