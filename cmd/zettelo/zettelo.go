// Main file for the zettelo command line tool.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ozcankasal/zettelo/internal"
	"github.com/ozcankasal/zettelo/internal/utils"
)

func main() {
	// Read folder from command line argument
	foldername := getFolderNameFromArgs()

	// Read configuration from environment variable
	config, err := getConfigFromEnv()
	if err != nil {
		fmt.Printf("Failed to read configuration: %v\n", err)
		os.Exit(1)
	}

	// Scan all markdown files in folder
	files, err := scanFolder(foldername)
	if err != nil {
		fmt.Printf("Failed to scan folder %s: %v\n", foldername, err)
		os.Exit(1)
	}
	fmt.Println(files)

	// Combine tagged lines from all files
	taggedLinesByFile := extractTaggedLinesFromFiles(files, *config)

	// Group tagged lines by canonical type and file path
	groupedLines := groupTaggedLines(taggedLinesByFile)

	// Convert to internal.TaggedLine slice
	outputLines := convertToTaggedLines(groupedLines)

	// Write JSON output to stdout
	err = utils.WriteJSON(os.Stdout, outputLines)
	if err != nil {
		fmt.Printf("Failed to write JSON output: %v\n", err)
		os.Exit(1)
	}
}

func getFolderNameFromArgs() string {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a folder name.")
		os.Exit(1)
	}
	return os.Args[1]
}

func getConfigFromEnv() (*internal.Config, error) {
	configData, err := ioutil.ReadFile(os.Getenv("ZETTELO_CONFIG"))
	if err != nil {
		return nil, err
	}

	return utils.ParseConfig(configData)
}

// Scan folder recursively for markdown files
func scanFolder(folderName string) ([]string, error) {
	var files []string
	err := filepath.Walk(folderName, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".md" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func extractTaggedLinesFromFiles(files []string, config internal.Config) map[string][]internal.TaggedLine {
	taggedLinesByFile := make(map[string][]internal.TaggedLine)
	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("Failed to read file %s: %v\n", file, err)
			os.Exit(1)
		}

		// Extract tagged lines from markdown file
		lines := utils.ExtractTaggedLines(file, data, config)

		// Append to taggedLinesByFile
		taggedLinesByFile[file] = lines
	}
	return taggedLinesByFile
}

// groupTaggedLines groups the tagged lines by canonical type and file path
func groupTaggedLines(taggedLinesByFile map[string][]internal.TaggedLine) map[string]map[string][]internal.ResultValue {
	groupedLines := make(map[string]map[string][]internal.ResultValue)

	for file, lines := range taggedLinesByFile {
		for _, line := range lines {
			if len(line.Values) > 0 {
				if _, ok := groupedLines[line.Tag]; !ok {
					groupedLines[line.Tag] = make(map[string][]internal.ResultValue)
				}
				groupedLines[line.Tag][file] = append(groupedLines[line.Tag][file], line.Values...)
			} else {
				if _, ok := groupedLines[line.Tag]; !ok {
					groupedLines[line.Tag] = make(map[string][]internal.ResultValue)
				}
				groupedLines[line.Tag][file] = append(groupedLines[line.Tag][file], internal.ResultValue{})
			}
		}
	}

	return groupedLines
}

// convertGroupedLinesToTaggedLines converts the grouped lines to an internal.TaggedLine slice
func convertToTaggedLines(groupedLines map[string]map[string][]internal.ResultValue) []internal.TaggedLine {
	var outputLines []internal.TaggedLine
	outputMap := make(map[string][]internal.ResultValue)

	for tag, fileLines := range groupedLines {
		for _, value := range fileLines {
			outputMap[tag] = append(outputMap[tag], value...)
		}
	}

	for tag, values := range outputMap {
		outputLines = append(outputLines, internal.TaggedLine{Tag: tag, Values: values})
	}

	return outputLines
}
