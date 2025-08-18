package cli

import (
	"errors"
	"os"

	"github.com/spf13/cobra"

	"github.com/imrancluster/sshmama/internal/crypto"
	"github.com/imrancluster/sshmama/internal/db"
	"github.com/imrancluster/sshmama/internal/util"
)

func cmdAttach() *cobra.Command {
	var name, file string
	c := &cobra.Command{
		Use:   "attach --name <entry> --file </path/to/privatekey>",
		Short: "Encrypt & attach a private file (e.g., PEM) to an entry",
		RunE: func(cmd *cobra.Command, args []string) error {
			if name == "" || file == "" {
				return errors.New("--name and --file are required")
			}
			ent, err := db.GetEntry(appState, name)
			if err != nil {
				return err
			}
			content, err := os.ReadFile(file)
			if err != nil {
				return err
			}
			pass1, err := util.ReadPassword("Choose passphrase: ")
			if err != nil {
				return err
			}
			pass2, err := util.ReadPassword("Repeat passphrase: ")
			if err != nil {
				return err
			}
			if string(pass1) != string(pass2) {
				return errors.New("passphrases do not match")
			}
			ct, err := crypto.EncryptPassphrase(string(pass1), content)
			if err != nil {
				return err
			}
			id := crypto.NewID()

			if err := db.PutAttachment(appState, id, ct); err != nil {
				return err
			}
			ent.AttachmentID = id
			return db.PutEntry(appState, ent)
		},
	}
	c.Flags().StringVar(&name, "name", "", "Entry name")
	c.Flags().StringVar(&file, "file", "", "Path to a private key or other secret file")
	return c
}
