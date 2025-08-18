package cli

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/imrancluster/sshmama/internal/db"
	"github.com/spf13/cobra"
)

func cmdList() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List saved entries",
		RunE: func(cmd *cobra.Command, args []string) error {
			entries, err := db.ListEntries(appState)
			if err != nil {
				return err
			}
			tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			fmt.Fprintln(tw, "NAME\tUSER@HOST:PORT\tTAGS\tATTACHED")
			for _, e := range entries {
				attached := "no"
				if e.AttachmentID != "" {
					attached = "yes"
				}
				fmt.Fprintf(tw, "%s\t%s@%s:%d\t%v\t%s\n",
					e.Name, e.User, e.Host, e.Port, e.Tags, attached)
			}
			return tw.Flush()
		},
	}
}
