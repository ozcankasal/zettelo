package utils_test

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ozcankasal/zettelo/internal"
	"github.com/ozcankasal/zettelo/internal/utils"
)

func TestWriteJSON(t *testing.T) {
	// Define some sample data
	lines := internal.TagList{
		{
			Tag: "tag1",
			Values: []internal.ResultValue{
				{
					FilePath: "file1",
					Line:     "line1",
				},
				{
					FilePath: "file2",
					Line:     "line2",
				},
			},
		},
		{
			Tag: "tag2",
			Values: []internal.ResultValue{
				{
					FilePath: "file3",
					Line:     "line3",
				},
				{
					FilePath: "file4",
					Line:     "line4",
				},
			},
		},
	}

	// Write the data to a buffer
	var buf bytes.Buffer
	_, err := utils.WriteJSON(&buf, lines)
	if err != nil {
		t.Errorf("WriteJSON failed: %v", err)
	}
}

func TestReadJSON(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expected    internal.TagList
		expectedErr error
	}{
		{
			name:  "valid input",
			input: `[{"tag": "tag1", "values": [{"file_path": "file1", "line": "line1"}]},{"tag": "tag2", "values": [{"file_path": "file2", "line": "line2"}]}]`,
			expected: internal.TagList{
				{
					Tag: "tag1",
					Values: []internal.ResultValue{
						{
							FilePath: "file1",
							Line:     "line1",
						},
					},
				},
				{
					Tag: "tag2",
					Values: []internal.ResultValue{
						{
							FilePath: "file2",
							Line:     "line2",
						},
					},
				},
			},
			expectedErr: nil,
		},
		{
			name:        "invalid input",
			input:       `[{"tag": 25}]]`,
			expected:    nil,
			expectedErr: &json.SyntaxError{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := bytes.NewBufferString(tc.input)
			result, err := utils.ReadJSON(r)

			if err != nil {
				if tc.expectedErr == nil {
					t.Errorf("Expected no error but got %v", err)
				} else if reflect.TypeOf(err) != reflect.TypeOf(tc.expectedErr) {
					t.Errorf("Expected error type %T but got %T", tc.expectedErr, err)
				}
			} else if tc.expectedErr != nil {
				t.Errorf("Expected error %v but got no error", tc.expectedErr)
			}

			if len(result) != len(tc.expected) {
				t.Errorf("Expected length %d but got %d", len(tc.expected), len(result))
			} else {
				for i := range result {
					if result[i].Tag != tc.expected[i].Tag {
						t.Errorf("Expected tag %s but got %s", tc.expected[i].Tag, result[i].Tag)
					}

					if len(result[i].Values) != len(tc.expected[i].Values) {
						t.Errorf("Expected values length %d but got %d", len(tc.expected[i].Values), len(result[i].Values))
					} else {
						for j := range result[i].Values {
							if result[i].Values[j].FilePath != tc.expected[i].Values[j].FilePath {
								t.Errorf("Expected file path %s but got %s", tc.expected[i].Values[j].FilePath, result[i].Values[j].FilePath)
							}

							if result[i].Values[j].Line != tc.expected[i].Values[j].Line {
								t.Errorf("Expected line %s but got %s", tc.expected[i].Values[j].Line, result[i].Values[j].Line)
							}
						}
					}
				}
			}
		})
	}
}
