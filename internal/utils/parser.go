/*
Package utils provides utility functions for working with JSON data and other common tasks.
Author: Özcan Kasal <ozcankasal@gmail.com>
Date:   April 3, 2023
*/

package utils

import (
	"bufio"
	"encoding/json"
	"io"
	"sort"

	"github.com/ozcankasal/zettelo/internal"
	"gopkg.in/yaml.v2"
)

/*
WriteJSON writes a JSON representation of a slice of TaggedLines to an io.Writer.

Usage:

	err := WriteJSON(w, lines)

Parameters:

	w (io.Writer): the io.Writer to write to
	lines (internal.TagList): the slice of TaggedLines to write

Returns:

	(error): if an error occurred during marshaling or writing, returns the error; otherwise, returns nil.
*/
func WriteJSON(lines internal.TagList) ([]byte, error) {
	sort.Sort(lines)
	// Convert []TaggedLine to []byte
	b, err := json.Marshal(lines)
	if err != nil {
		return nil, err
	}
	return b, nil
}

/*
ReadJSON reads a JSON input from an io.Reader and returns an array of TaggedLines and an error if one occurs.

It reads from the given io.Reader into a []byte buffer and unmarshals the []byte into a []TaggedLine.

Usage:

Call ReadJSON with an io.Reader containing a single line of valid JSON to parse the data into an array of TaggedLines.

Parameters:

r: An io.Reader that holds the JSON input to be read.

Returns:

internal.TagList: An array of TaggedLines that holds the parsed JSON data.

error: An error, if any, that occurred during the process of reading and unmarshaling the JSON input.

Note:

This function expects the JSON input to be a single line of valid JSON.
*/
func ReadJSON(r io.Reader) (internal.TagList, error) {
	// Read from io.Reader into a []byte buffer
	buf := bufio.NewReader(r)
	b, err := buf.ReadBytes('\n')
	if err != nil && err != io.EOF {
		return nil, err
	}

	// Unmarshal []byte into []TaggedLine
	var lines internal.TagList
	err = json.Unmarshal(b, &lines)
	if err != nil {
		return nil, err
	}
	return lines, err
}

func ParseConfig(configData []byte) (*internal.Config, error) {

	var config internal.Config
	err := yaml.Unmarshal(configData, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
