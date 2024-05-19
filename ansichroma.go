package ansichroma

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/quick"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/charmbracelet/lipgloss"
)

type NotTextFile struct {
	FileName string
}

// test
func (e *NotTextFile) Error() string {
	return fmt.Sprintf("'%s' is not a text file", e.FileName)
}

// Hightlight text file. When linesToRead is 0. the complete file is read.
func HighlightFromFile(path string, linesToRead int, style, background string) (resultString string, err error) {
	var buf bytes.Buffer
	var fileContent string
	var codeHighlight []chroma.Token

	fileFullPath, err := filepath.Abs(path)

	if err != nil {
		return "", errors.New("this file path is invalid")
	}

	fileName := filepath.Base(fileFullPath)

	lexer := lexers.Match(fileName)
	if lexer == nil {
		return "", &NotTextFile{FileName: fileName}
	}

	if linesToRead == 0 {
		file, err := os.ReadFile(fileFullPath)
		if err != nil {
			return "", err
		}
		fileContent = string(file)
	} else {
		file, err := os.Open(fileFullPath)
		if err != nil {
			return "", err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lineCount := 0

		for scanner.Scan() {
			fileContent += scanner.Text() + "\n"
			lineCount++
			if linesToRead > 0 && lineCount >= linesToRead {
				break
			}
		}

		if err := scanner.Err(); err != nil {
			return "", err
		}
	}

	if err := quick.Highlight(&buf, fileContent, lexer.Config().Name, "json", ""); err != nil {
		return "", err
	}

	if err := json.Unmarshal(buf.Bytes(), &codeHighlight); err != nil {
		return "", err
	}

	s := styles.Get(style)

	for _, data := range codeHighlight {
		color := s.Get(data.Type)

		renderString, countLineBreaks := trimTrailingNewlines(data.Value)

		colorStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color(color.Colour.String())).
			Background(lipgloss.Color(background)).
			Bold(getTrileanToBool(color.Bold)).
			Italic(getTrileanToBool(color.Italic)).
			Underline(getTrileanToBool(color.Underline))

		resultString += colorStyle.Render(renderString)
		resultString += strings.Repeat("\n", countLineBreaks)
	}

	return resultString, nil
}

// Hightlight text string.
func HightlightString(fileContent, format string, style, background string) (resultString string, err error) {
	var buf bytes.Buffer
	var codeHighlight []chroma.Token

	if err := quick.Highlight(&buf, fileContent, format, "json", ""); err != nil {
		return "", err
	}

	if err := json.Unmarshal(buf.Bytes(), &codeHighlight); err != nil {
		return "", err
	}

	s := styles.Get(style)

	for _, data := range codeHighlight {
		color := s.Get(data.Type)

		renderString, countLineBreaks := trimTrailingNewlines(data.Value)

		colorStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color(color.Colour.String())).
			Background(lipgloss.Color(background)).
			Bold(getTrileanToBool(color.Bold)).
			Italic(getTrileanToBool(color.Italic)).
			Underline(getTrileanToBool(color.Underline))

		resultString += colorStyle.Render(renderString)
		resultString += strings.Repeat("\n", countLineBreaks)
	}

	return resultString, nil
}
