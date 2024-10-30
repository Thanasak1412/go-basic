package main

import (
	"bytes"
	"fmt"
	"github.com/Thanasak1412/read-files/fileutils"
	"os"
	"text/template"
)

func main() {
	// Define a template string
	tmpl := `Content => {{ .Content }}`

	// Create a new template and parse the template string
	t, _ := template.New("test").Parse(tmpl)

	rootPath, _ := os.Getwd()

	content, err := fileutils.ReadTextFile(rootPath + "/data/text.txt")

	if err != nil {
		panic(err)
	} else {
		data := struct {
			Content string
		}{
			Content: content,
		}

		// Create a buffer to capture the output
		var result bytes.Buffer

		// Execute the template and write the output to stdout
		err = t.Execute(&result, data)

		if err != nil {
			panic(err)
		}

		fmt.Println(result.String())
	}

	newContent := fmt.Sprintf("old content %v + new content", content)

	_ = fileutils.WriteFile(rootPath+"/data/text.txt", []byte(newContent))
}
