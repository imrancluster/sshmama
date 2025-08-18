package cli

import (
	"encoding/json"
	"os"

	"github.com/imrancluster/sshmama/internal/db"
	"github.com/spf13/cobra"
)

func cmdExport() *cobra.Command {
	var file string
	c := &cobra.Command{
		Use:   "export --file <backup.json>",
		Short: "Export all entries and attachments to a JSON file",
		RunE: func(cmd *cobra.Command, args []string) error {
			if file == "" {
				return cmd.Usage()
			}

			entries, err := db.ListEntries(appState)
			if err != nil {
				return err
			}

			data := make(map[string]interface{})
			data["entries"] = entries

			// Export attachments
			attachments := map[string][]byte{}
			for _, e := range entries {
				if e.AttachmentID != "" {
					b, err := db.GetAttachment(appState, e.AttachmentID)
					if err != nil {
						return err
					}
					attachments[e.AttachmentID] = b
				}
			}
			data["attachments"] = attachments

			buf, _ := json.MarshalIndent(data, "", "  ")
			return os.WriteFile(file, buf, 0o600)
		},
	}

	c.Flags().StringVar(&file, "file", "", "Output JSON file path (required)")
	return c
}
