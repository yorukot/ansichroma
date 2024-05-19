package ansichroma_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/alecthomas/chroma/lexers"
	"github.com/charmbracelet/lipgloss"
	"github.com/yorukot/ansichroma"
)

func TestAllFilesWithPathInDirectory(t *testing.T) {
	testDir := "./testFile"

	err := filepath.Walk(testDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		format := lexers.Match(filepath.Base(path))
		
		if format == nil {
			fmt.Printf("\n=======================================\n")
			fmt.Print("PASS unsupported Language or not text file\n")

			return nil
		}
		
		result, err := ansichroma.HighlightFromFile(path, 0, "catppuccin-frappe", lipgloss.Color("#1e1e2e"))
		fmt.Printf("\n=======================================\n[%s Test] File: %s\n\n", format.Config().Name, path)
		fmt.Println(result)
		if err != nil {
			t.Errorf("Error while running %s: %s", path, err)
		}
		return nil
	})

	if err != nil {
		t.Errorf("Error while walking directory: %s", err)
	}
}

func TestAllFilesWithStringInDirectory(t *testing.T) {
	testDir := "./testFile"

	err := filepath.Walk(testDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		content, err := os.ReadFile(path)
		if err != nil {
			t.Errorf("Error while read file %s: %s", path, err)
		}
		format := lexers.Match(filepath.Base(path))
		
		if format == nil {
			t.Log("PASS unsupported Language or not text file")
			return nil
		}

		result, err := ansichroma.HightlightString(string(content),  format.Config().Name, "catppuccin-frappe", lipgloss.Color("#1e1e2e"))
		fmt.Printf("\n=======================================\n[%s Test] File: %s\n\n",format.Config().Name, path)
		fmt.Println(result)
		if err != nil {
			t.Errorf("Error while running %s: %s", path, err)
		}
		return nil
	})

	if err != nil {
		t.Errorf("Error while walking directory: %s", err)
	}
}
