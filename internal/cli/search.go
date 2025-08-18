package cli

import (
	"fmt"
	"strings"

	"github.com/imrancluster/sshmama/internal/db"
	"github.com/imrancluster/sshmama/internal/search"
	"github.com/spf13/cobra"
)

func cmdSearch() *cobra.Command {
	var (
		query   string
		connect bool
	)
	c := &cobra.Command{
		Use:   "search",
		Short: "Search entries (fzf if available)",
		RunE: func(cmd *cobra.Command, args []string) error {
			entries, err := db.ListEntries(appState)
			if err != nil {
				return err
			}
			// If no query and fzf exists, run an interactive picker
			if query == "" && search.HasFZF() {
				lines := make([]string, 0, len(entries))
				for _, e := range entries {
					lines = append(lines, fmt.Sprintf("%s\t%s@%s:%d\t[%s]",
						e.Name, e.User, e.Host, e.Port, strings.Join(e.Tags, ",")))
				}
				sel, err := search.RunFZF(lines)
				if err != nil {
					return err
				}
				name := strings.SplitN(sel, "\t", 2)[0]
				if connect {
					// delegate to connect command
					return runConnect(name, false)
				}
				fmt.Println(name)
				return nil
			}

			// substring filter
			matches := search.Filter(entries, query)
			for _, e := range matches {
				fmt.Printf("%s\t%s@%s:%d\t[%s]\n", e.Name, e.User, e.Host, e.Port, strings.Join(e.Tags, ","))
			}
			return nil
		},
	}
	c.Flags().StringVarP(&query, "query", "q", "", "Substring query")
	c.Flags().BoolVar(&connect, "connect", false, "Connect to selected entry")
	return c
}
