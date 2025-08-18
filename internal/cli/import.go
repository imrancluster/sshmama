package cli

import (
	"encoding/json"
	"os"

	"github.com/imrancluster/sshmama/internal/db"
	"github.com/imrancluster/sshmama/internal/model"
	"github.com/spf13/cobra"
)

func cmdImport() *cobra.Command {
	var file string
	c := &cobra.Command{
		Use:   "import --file <backup.json>",
		Short: "Import entries and attachments from a JSON file",
		RunE: func(cmd *cobra.Command, args []string) error {
			if file == "" {
				return cmd.Usage()
			}

			buf, err := os.ReadFile(file)
			if err != nil {
				return err
			}

			var data map[string]json.RawMessage
			if err := json.Unmarshal(buf, &data); err != nil {
				return err
			}

			// Import entries
			var entries []model.Entry
			if err := json.Unmarshal(data["entries"], &entries); err != nil {
				return err
			}

			for _, e := range entries {
				if err := db.PutEntry(appState, &e); err != nil {
					return err
				}
			}

			// Import attachments
			var attachments map[string][]byte
			if err := json.Unmarshal(data["attachments"], &attachments); err != nil {
				return err
			}
			for id, b := range attachments {
				if err := db.PutAttachment(appState, id, b); err != nil {
					return err
				}
			}
			return nil
		},
	}

	c.Flags().StringVar(&file, "file", "", "Input JSON file path (required)")
	return c
}
