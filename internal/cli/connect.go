package cli

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"github.com/imrancluster/sshmama/internal/crypto"
	"github.com/imrancluster/sshmama/internal/db"
	"github.com/imrancluster/sshmama/internal/ssh"
	"github.com/imrancluster/sshmama/internal/util"
)

func cmdConnect() *cobra.Command {
	var dryRun bool
	c := &cobra.Command{
		Use:   "connect [name]",
		Short: "Connect to an entry via ssh",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var name string
			if len(args) > 0 {
				name = args[0]
			}
			return runConnect(name, dryRun)
		},
	}
	c.Flags().BoolVar(&dryRun, "dry-run", false, "Print the ssh command without running it")
	return c
}

func runConnect(name string, dryRun bool) error {
	if name == "" {
		// Fetch all entries from DB
		entries, err := db.ListEntries(appState)
		if err != nil {
			return fmt.Errorf("failed to list entries: %v", err)
		}
		if len(entries) == 0 {
			return fmt.Errorf("no SSH entries found. Add one with `sshmama add`")
		}

		// Show TUI selection
		prompt := promptui.Select{
			Label: "Select SSH host",
			Items: entries,
			Templates: &promptui.SelectTemplates{
				Label:    "{{ . }}?",
				Active:   "▶ {{ .Name | cyan }} ({{ .Host }})",
				Inactive: "  {{ .Name }} ({{ .Host }})",
				Selected: "✔ {{ .Name | green }} ({{ .Host }})",
			},
			Size: 10,
			Searcher: func(input string, index int) bool {
				entry := entries[index]
				return contains(entry.Name, input) || contains(entry.Host, input)
			},
		}

		i, _, err := prompt.Run()
		if err != nil {
			return fmt.Errorf("prompt failed: %v", err)
		}

		name = entries[i].Name
	}

	ent, err := db.GetEntry(appState, name)
	if err != nil {
		return err
	}

	var keyPath string
	if ent.AttachmentID != "" {
		pass, err := util.ReadPassword("Passphrase for attached file: ")
		if err != nil {
			return err
		}
		cipher, err := db.GetAttachment(appState, ent.AttachmentID)
		if err != nil {
			return err
		}
		plain, err := crypto.DecryptPassphrase(string(pass), cipher)
		if err != nil {
			return errors.New("decryption failed (wrong passphrase?)")
		}
		// write to temp with 0600
		tmpDir := filepath.Join(os.TempDir(), "sshmama")
		if err := os.MkdirAll(tmpDir, 0o700); err != nil {
			return err
		}
		tmpFile := filepath.Join(tmpDir, ent.Name+"-key")
		if err := os.WriteFile(tmpFile, plain, 0o600); err != nil {
			return err
		}
		defer func() { _ = os.Remove(tmpFile) }()
		keyPath = tmpFile
	} else if ent.KeyPath != "" {
		keyPath = ent.KeyPath
	}

	return ssh.RunSSH(ssh.CmdOptions{
		User: ent.User, Host: ent.Host, Port: ent.Port, Key: keyPath, DryRun: dryRun,
	})
}

func contains(s, substr string) bool {
	return len(substr) == 0 || (len(s) >= len(substr) &&
		(len(s) >= len(substr) && (stringContainsIgnoreCase(s, substr))))
}

func stringContainsIgnoreCase(str, substr string) bool {
	return (len(str) >= len(substr)) && (stringIndexIgnoreCase(str, substr) != -1)
}

// For now, let's leave a stub – you can later use strings.IndexFold (Go 1.20+)
func stringIndexIgnoreCase(str, substr string) int {
	return -1
}
