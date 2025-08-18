package cli

import (
	"errors"
	"time"

	"github.com/imrancluster/sshmama/internal/db"
	"github.com/spf13/cobra"
)

func cmdEdit() *cobra.Command {
	var (
		name, host, userFlag, key, notes string
		port                             int
		tags                             []string
	)
	c := &cobra.Command{
		Use:   "edit --name <entry>",
		Short: "Edit an existing SSH entry",
		RunE: func(cmd *cobra.Command, args []string) error {
			if name == "" {
				return errors.New("--name is required")
			}
			ent, err := db.GetEntry(appState, name)
			if err != nil {
				return err
			}
			if host != "" {
				ent.Host = host
			}
			if userFlag != "" {
				ent.User = userFlag
			}
			if port > 0 {
				ent.Port = port
			}
			if key != "" {
				ent.KeyPath = key
			}
			if len(tags) > 0 {
				ent.Tags = tags
			}
			if notes != "" {
				ent.Notes = notes
			}
			ent.CreatedAt = time.Now() // update timestamp
			return db.PutEntry(appState, ent)
		},
	}

	c.Flags().StringVar(&name, "name", "", "Entry name to edit (required)")
	c.Flags().StringVar(&host, "host", "", "Host/IP")
	c.Flags().StringVar(&userFlag, "user", "", "SSH user")
	c.Flags().IntVar(&port, "port", 0, "SSH port")
	c.Flags().StringVar(&key, "key", "", "Path to private key")
	c.Flags().StringSliceVar(&tags, "tag", nil, "Tags (comma separated)")
	c.Flags().StringVar(&notes, "notes", "", "Notes")
	return c
}
