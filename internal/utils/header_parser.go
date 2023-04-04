package utils

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

func SyncHeader(filePath string) {
	header := readHeader(filePath)
	if header != "" {
		id := getId(header)

		if id != "" {

		} else {
			fmt.Println("No ID found in header of", filePath)
			newHeaderText := addId(header, getUUID())
			updateHeader(filePath, newHeaderText)
		}
	} else {
		fmt.Println("No header found in", filePath)
		addHeader(filePath)
		newHeaderText := addId(readHeader(filePath), getUUID())
		updateHeader(filePath, newHeaderText)
	}
}

func checkHeader(filePath string) bool {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	headerRegex := regexp.MustCompile(`(?s)\A---\n(.+)\n---\n`)
	headerMatches := headerRegex.FindStringSubmatch(string(content))

	if len(headerMatches) > 1 {
		return true
	} else {
		return false
	}
}

func readHeader(filePath string) string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	headerRegex := regexp.MustCompile(`(?s)\A---\n(.+)\n---\n`)
	headerMatches := headerRegex.FindStringSubmatch(string(content))
	if len(headerMatches) > 1 {
		return headerMatches[1]
	} else {
		return ""
	}
}

func getId(headerText string) string {
	idRegex := regexp.MustCompile(`(\w+):\s*(.*)\n`)
	idMatches := idRegex.FindStringSubmatch(headerText)
	if len(idMatches) > 2 {
		return idMatches[2]
	} else {
		return ""
	}
}

// this function adds an ID to the header text already given in an extra line
func addId(headerText string, uuid string) string {
	newText := headerText + "\n" + "id: " + uuid + "\n"
	return newText
}

// this function generates new uuid
func getUUID() string {
	uuid := uuid.New()
	return uuid.String()
}

// this function writes the new header text to the file
func updateHeader(filePath string, newHeaderText string) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	headerRegex := regexp.MustCompile(`(?s)\A---\n(.+)\n---\n`)
	headerMatches := headerRegex.FindStringSubmatch(string(content))
	if len(headerMatches) > 1 {
		headerText := headerMatches[1]
		newContent := strings.Replace(string(content), headerText, newHeaderText, 1)

		// write the new content to the file
		ioutil.WriteFile(filePath, []byte(newContent), 0)
	}
}

func addHeader(filePath string) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	headerText := "---\n \n---\n"
	newContent := headerText + string(content)

	// write the new content to the file
	ioutil.WriteFile(filePath, []byte(newContent), 0)
}
