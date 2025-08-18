package crypto

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"filippo.io/age"
)

// Encrypt data with passphrase
func Encrypt(data []byte, pass string) ([]byte, error) {
	recipient, err := age.NewScryptRecipient(pass)
	if err != nil {
		return nil, fmt.Errorf("failed to create recipient: %w", err)
	}

	var out bytes.Buffer
	w, err := age.Encrypt(&out, recipient)
	if err != nil {
		return nil, fmt.Errorf("failed to create encryptor: %w", err)
	}

	if _, err := w.Write(data); err != nil {
		return nil, fmt.Errorf("failed to write data: %w", err)
	}
	if err := w.Close(); err != nil {
		return nil, fmt.Errorf("failed to close writer: %w", err)
	}

	return out.Bytes(), nil
}

// Decrypt data with passphrase
func Decrypt(data []byte, pass string) ([]byte, error) {
	identity, err := age.NewScryptIdentity(pass)
	if err != nil {
		return nil, fmt.Errorf("failed to create identity: %w", err)
	}

	r, err := age.Decrypt(bytes.NewReader(data), identity)
	if err != nil {
		return nil, fmt.Errorf("failed to create decryptor: %w", err)
	}

	out, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read decrypted data: %w", err)
	}

	return out, nil
}

// Utility function: save encrypted file
func SaveEncryptedFile(filename string, data []byte, pass string) error {
	enc, err := Encrypt(data, pass)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, enc, 0600)
}

// Utility function: load and decrypt file
func LoadDecryptedFile(filename string, pass string) ([]byte, error) {
	enc, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return Decrypt(enc, pass)
}

// EncryptPassphrase encrypts data with a passphrase.
func EncryptPassphrase(passphrase string, data []byte) ([]byte, error) {
	recipient, err := age.NewScryptRecipient(passphrase)
	if err != nil {
		return nil, fmt.Errorf("failed to create recipient: %w", err)
	}

	var buf bytes.Buffer
	w, err := age.Encrypt(&buf, recipient)
	if err != nil {
		return nil, fmt.Errorf("failed to start encryption: %w", err)
	}

	if _, err := w.Write(data); err != nil {
		return nil, fmt.Errorf("failed to write data: %w", err)
	}

	if err := w.Close(); err != nil {
		return nil, fmt.Errorf("failed to finalize encryption: %w", err)
	}

	return buf.Bytes(), nil
}

// DecryptPassphrase decrypts data with a passphrase.
func DecryptPassphrase(passphrase string, data []byte) ([]byte, error) {
	identity, err := age.NewScryptIdentity(passphrase)
	if err != nil {
		return nil, fmt.Errorf("failed to create identity: %w", err)
	}

	r, err := age.Decrypt(bytes.NewReader(data), identity)
	if err != nil {
		return nil, fmt.Errorf("failed to start decryption: %w", err)
	}

	out, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read decrypted data: %w", err)
	}

	return out, nil
}

// NewID generates a random ID for stored secrets/entries.
func NewID() string {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return "randomid"
	}
	return hex.EncodeToString(b)
}
