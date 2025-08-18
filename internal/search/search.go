package search

import (
	"bytes"
	"os/exec"
	"runtime"
	"strings"

	"github.com/imrancluster/sshmama/internal/model"
)

func Filter(entries []model.Entry, q string) []model.Entry {
	q = strings.ToLower(strings.TrimSpace(q))
	if q == "" {
		return entries
	}
	var out []model.Entry
	for _, e := range entries {
		hay := strings.ToLower(strings.Join([]string{
			e.Name, e.Host, e.User, strings.Join(e.Tags, ","), e.Notes,
		}, " "))
		if strings.Contains(hay, q) {
			out = append(out, e)
		}
	}
	return out
}

func HasFZF() bool {
	_, err := exec.LookPath("fzf")
	return err == nil
}

func RunFZF(lines []string) (string, error) {
	if runtime.GOOS == "windows" {
		return "", exec.ErrNotFound
	}
	cmd := exec.Command("fzf", "--ansi", "--prompt", "entry> ")
	cmd.Stdin = bytes.NewBufferString(strings.Join(lines, "\n"))
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
