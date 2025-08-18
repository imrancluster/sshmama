package cli

import (
	"errors"
	"os/user"
	"time"

	"github.com/imrancluster/sshmama/internal/db"
	"github.com/imrancluster/sshmama/internal/model"
	"github.com/spf13/cobra"
)

func cmdAdd() *cobra.Command {
	var (
		name, host, userFlag, key, notes string
		port                             int
		tags                             []string
	)
	c := &cobra.Command{
		Use:   "add",
		Short: "Add or update an SSH entry",
		RunE: func(cmd *cobra.Command, args []string) error {
			if name == "" || host == "" {
				return errors.New("--name and --host are required")
			}
			if userFlag == "" {
				if u, err := user.Current(); err == nil {
					userFlag = u.Username
				} else {
					userFlag = "root"
				}
			}
			if port == 0 {
				port = 22
			}
			e := &model.Entry{
				Name:      name,
				Host:      host,
				User:      userFlag,
				Port:      port,
				KeyPath:   key,
				Tags:      tags,
				Notes:     notes,
				CreatedAt: time.Now(),
			}
			return db.PutEntry(appState, e)
		},
	}
	c.Flags().StringVar(&name, "name", "", "Unique entry name (e.g., prod)")
	c.Flags().StringVar(&host, "host", "", "Host or IP (e.g., 203.0.113.10)")
	c.Flags().StringVar(&userFlag, "user", "", "SSH user (default: current user)")
	c.Flags().IntVar(&port, "port", 22, "SSH port")
	c.Flags().StringVar(&key, "key", "", "Path to private key (optional)")
	c.Flags().StringSliceVar(&tags, "tag", nil, "Tag(s), repeatable")
	c.Flags().StringVar(&notes, "notes", "", "Freeform notes")
	c.Flags().SortFlags = false
	c.Example = `  sshmama add --name prod --host 203.0.113.10 --user ubuntu --port 22 --tag aws --tag critical`
	return c
}
