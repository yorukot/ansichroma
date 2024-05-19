# ansichroma
ansichroma is a Go package for syntax highlighting text files and strings using ANSI escape codes. It supports various programming languages and styles, providing a customizable way to highlight code for terminal output.

ansichrom Highly based on (chroma)[https://github.com/alecthomas/chroma]

## Installation
To install ansichroma, use go get:

```bash
go get github.com/yorukot/ansichroma
```

## Usage
```go
go
Copy code
package main

import (
	"fmt"
	"github.com/yorukot/ansichroma"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	// Highlight text file
	path := "path/to/your/file.txt"
	linesToRead := 10
	style := "monokai"
	background := "#1e1e2e"
	resultString, err := ansichroma.HighlightFromFile(path, linesToRead, style, background)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(resultString)

	// Highlight text string
	fileContent := "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, World!\")\n}"
	format := "go"
	style := "monokai"
	background := "#1e1e2e"
	resultString, err := ansichroma.HightlightString(fileContent, format, style, background)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(resultString)
}
```
## Documentation
For detailed documentation and examples, please refer to the [GoDoc page](https://pkg.go.dev/github.com/yorukot/ansichroma).

## Contributing
Contributions are welcome! Feel free to open issues or submit pull requests on the GitHub repository.

## License
This project is licensed under the MIT License - see the [LICENSE](/LICENSE) file for details.