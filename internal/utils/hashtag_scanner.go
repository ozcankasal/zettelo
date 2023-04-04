package utils

import (
	"bufio"
	"bytes"
	"regexp"
	"strings"

	"github.com/ozcankasal/zettelo/internal"
)

/*
MapTagToCanonicalType maps a tag to its canonical type.

Usage:

	tag := "foo"
	canonicalType := MapTagToCanonicalType(tag, config)

Parameters:

	tag (string): the tag to map
	config (internal.Config): the configuration to use

Returns:

	(string): the canonical type of the tag, or an empty string if the tag is not mapped to a canonical type
*/
func MapTagToCanonicalType(tag string, config internal.Config) string {
	for tagKey, tagValue := range config.TagMappings {
		if tag == tagKey {
			return tagValue
		}
	}

	return ""
}

/*
removeHashtagsFromLine removes hashtags from a line.

Usage:

	line := "This is a #line with #hashtags"
	lineWithoutHashtags := removeHashtagsFromLine(line)

Parameters:

	line (string): the line to remove hashtags from

Returns:

	(string): the line without hashtags
*/
func RemoveHashtagsFromLine(line string) string {
	return regexp.MustCompile(`(?:^|\s)(#[^\s]+)`).ReplaceAllString(line, "")
}

func extractTagsFromLine(line string) []string {
	re := regexp.MustCompile(`(?:^|\s)(#[^\s#][^\s]*)`)
	return re.FindAllString(line, -1)
}

/*
ExtractTaggedLines extracts tagged lines from a file.

Usage:

	fileName := "test.md"
	data := []byte("#tag1: value1	#tag2: value2")
	taggedLines := ExtractTaggedLines(fileName, data, config)

Parameters:

	fileName (string): the name of the file
	data ([]byte): the file contents
	config (internal.Config): the configuration to use

Returns:

	([]internal.TaggedLine): the tagged lines in the file contents with their values and file paths (if any)
*/
func ExtractTaggedLines(fileName string, data []byte, config internal.Config) []internal.TaggedLine {
	var result []internal.TaggedLine
	var currentTag string

	reader := bytes.NewReader(data)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()

		tags := extractTagsFromLine(line)

		if len(tags) > 0 {
			for _, t := range tags {
				tag := strings.TrimSpace(t)
				value := strings.TrimSpace(strings.Replace(line, t, "", -1))
				value = strings.TrimSpace(RemoveHashtagsFromLine(value))

				// Map the tag to its canonical type
				canonicalType := MapTagToCanonicalType(tag, config)
				if canonicalType == "" {
					canonicalType = tag
				}

				if len(value) > 0 {
					// This line has a value
					found := false
					for i := range result {
						if result[i].Tag == canonicalType {
							result[i].Tag = canonicalType
							canonicalValue := internal.ResultValue{Line: value, FilePath: fileName}
							result[i].Values = append(result[i].Values, canonicalValue)
							found = true
							break
						}
					}
					if !found {
						// This is the first line with this tag
						currentValue := internal.ResultValue{Line: value, FilePath: fileName}
						result = append(result, internal.TaggedLine{Tag: canonicalType, Values: []internal.ResultValue{currentValue}})

					}
				} else {
					// This line is a tag with an empty value
					found := false
					for i := range result {
						if result[i].Tag == canonicalType {
							found = true
							break
						}
					}
					if !found {
						// This is the first line with this tag
						result = append(result, internal.TaggedLine{Tag: canonicalType, Values: []internal.ResultValue{}})
					}
				}

				currentTag = canonicalType
			}
		} else {
			// This line is not a tag line, ignore it
			if len(result) > 0 && currentTag == result[len(result)-1].Tag {
				// This line is a continuation of the previous line
			} else {
				currentTag = ""
			}
		}
	}

	return result
}
