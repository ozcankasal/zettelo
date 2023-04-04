// Main file for the zettelo command line tool.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ozcankasal/zettelo/internal"
	"github.com/ozcankasal/zettelo/internal/utils"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var hashtags []byte

func main() {
	updates := make(chan []string)
	// Initialize the file system watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err)
	}
	defer watcher.Close()

	// Read configuration from environment variable
	config, err := getConfigFromEnv()
	if err != nil {
		fmt.Printf("Failed to read configuration: %v\n", err)
		os.Exit(1)
	}

	// Read folder from command line argument
	foldername := getFolderNameFromArgs()
	hashtags = getHashtags(foldername, config)

	scanFolder(foldername)

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					hashtags = getHashtags(foldername, config)
					updates <- []string{"update"}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = filepath.Walk(foldername, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			err = watcher.Add(path)
			if err != nil {
				fmt.Println(err)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	err = watcher.Add(foldername)
	if err != nil {
		fmt.Println(err)
	}

	http.Handle("/", http.FileServer(http.Dir("./static")))

	go handle(updates)
	fmt.Println(http.ListenAndServe(":8080", nil))
}

func handle(updates chan []string) {
	http.HandleFunc("/hashtags", func(w http.ResponseWriter, r *http.Request) {

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Error upgrading connection:", err)
			return
		}
		defer conn.Close()
		conn.WriteMessage(websocket.TextMessage, hashtags)

		for {
			lines := <-updates
			fmt.Println(lines)
			conn.WriteMessage(websocket.TextMessage, hashtags)
		}

	})
}

func getHashtags(folderName string, config *internal.Config) []byte {

	files, err := scanFolder(folderName)
	if err != nil {
		fmt.Printf("Failed to scan folder %s: %v\n", folderName, err)
		os.Exit(1)
	}

	// Combine tagged lines from all files
	taggedLinesByFile := extractTaggedLinesFromFiles(files, *config)

	// Group tagged lines by canonical type and file path
	groupedLines := groupTaggedLines(taggedLinesByFile)

	// Convert to internal.TaggedLine slice
	outputLines := convertToTaggedLines(groupedLines)

	// Write JSON output to stdout
	b, err := utils.WriteJSON(os.Stdout, outputLines)
	if err != nil {
		fmt.Printf("Failed to write JSON output: %v\n", err)
		os.Exit(1)
	}

	return b
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
			utils.SyncHeader(path)
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
