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

const configContent = `
# This is the configuration file for Zettelo.
# It is written in YAML format.
# For more information about the YAML format, see http://yaml.org/

# Web server configuration
web:
  port: 8080
  host: localhost

# Application-specific settings
app:
  tag_mappings:
    #todo: #todo
    #to-do: #todo
    #todo: #todo
  folders:
    - /path/to/folder1
    - /path/to/folder2
`

func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	configDir := filepath.Join(homeDir, ".zettelo")
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		if err := os.Mkdir(configDir, 0755); err != nil {
			return "", err
		}
	}

	configPath := filepath.Join(configDir, "config.yaml")
	return configPath, nil
}

func createConfigFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(configContent)
	if err != nil {
		return err
	}
	return nil
}

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
	configPath, _ := getConfigPath()

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := createConfigFile(configPath); err != nil {
			fmt.Printf("Failed to create configuration file: %v\n", err)
			os.Exit(1)
		}
	}

	// read the configuration file
	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	config, err := utils.ParseConfig(configData)
	if err != nil {
		fmt.Printf("Failed to read configuration: %v\n", err)
		os.Exit(1)
	}

	// Read folder from command line argument
	folderList := config.App.Folders
	if len(folderList) == 0 {
		fmt.Println("No folders specified in configuration file.")
		os.Exit(1)
	}
	for _, folder := range folderList {
		scanFolder(folder)
	}

	hashtags = getHashtags(folderList, config)

	for folder := range folderList {
		foldername := folderList[folder]
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
	}

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					hashtags = getHashtags(folderList, config)
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

	http.Handle("/", http.FileServer(http.Dir("./static")))

	go handle(updates)

	url := fmt.Sprintf("%s:%d", config.Web.Host, config.Web.Port)
	fmt.Printf("Server is listening on %s. Click %s to open in browser.\n", url, url)

	fmt.Println(http.ListenAndServe(url, nil))
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

		for range updates {
			conn.WriteMessage(websocket.TextMessage, hashtags)
		}

	})
}

func getHashtags(folderList []string, config *internal.Config) []byte {
	tempHashtagList := internal.TagList{}
	for _, folderName := range folderList {
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

		tempHashtagList = append(tempHashtagList, outputLines...)

	}

	// Write JSON output to stdout
	b, err := utils.WriteJSON(tempHashtagList)

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

// Scan folder recursively for markdown files
func scanFolder(folderName string) ([]string, error) {
	fmt.Println(folderName)
	var files []string
	err := filepath.Walk(folderName, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".md" {
			fmt.Println(path)
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

func extractTaggedLinesFromFiles(files []string, config internal.Config) map[string]internal.TagList {
	taggedLinesByFile := make(map[string]internal.TagList)
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
func groupTaggedLines(taggedLinesByFile map[string]internal.TagList) map[string]map[string][]internal.ResultValue {
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
func convertToTaggedLines(groupedLines map[string]map[string][]internal.ResultValue) internal.TagList {
	var outputLines internal.TagList
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
