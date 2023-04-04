package utils_test

import (
	"reflect"
	"testing"

	"github.com/ozcankasal/zettelo/internal"
	"github.com/ozcankasal/zettelo/internal/utils"
)

func TestMapTagToCanonicalType(t *testing.T) {
	// Define test cases
	testCases := []struct {
		tag      string
		config   internal.Config
		expected string
	}{
		{
			tag: "tag1",
			config: internal.Config{
				TagMappings: map[string]string{
					"tag1": "canonicalType1",
					"tag2": "canonicalType2",
				},
			},
			expected: "canonicalType1",
		},
		{
			tag: "tag2",
			config: internal.Config{
				TagMappings: map[string]string{
					"tag1": "canonicalType1",
					"tag2": "canonicalType2",
				},
			},
			expected: "canonicalType2",
		},
		{
			tag: "tag3",
			config: internal.Config{
				TagMappings: map[string]string{
					"tag1": "canonicalType1",
					"tag2": "canonicalType2",
				},
			},
			expected: "",
		},
	}

	// Run test cases
	for _, testCase := range testCases {
		actual := utils.MapTagToCanonicalType(testCase.tag, testCase.config)

		if actual != testCase.expected {
			t.Errorf("Expected %s, got %s", testCase.expected, actual)
		}
	}
}

func TestRemoveHashtagsFromLine(t *testing.T) {
	// Define test cases
	testCases := []struct {
		line     string
		expected string
	}{
		{
			line:     "This is a #line with #hashtags",
			expected: "This is a with",
		},
		{
			line:     "This is a #line with #hashtags and #more hashtags",
			expected: "This is a with and hashtags",
		},
		{
			line:     "This is a line with no hashtags",
			expected: "This is a line with no hashtags",
		},
	}

	// Run test cases
	for _, testCase := range testCases {
		actual := utils.RemoveHashtagsFromLine(testCase.line)
		if actual != testCase.expected {
			t.Errorf("Expected %s, got %s", testCase.expected, actual)
		}
	}
}

func TestExtractTaggedLines(t *testing.T) {
	// Define test data
	fileName := "test.txt"
	data := []byte("#tag1 value1\n#tag2 value2\nline 1\n#tag1 value3\n")
	config := internal.Config{TagMappings: map[string]string{"#tag1": "#canonicalTag1", "#tag2": "#canonicalTag2"}}

	// Expected output
	expected := internal.TagList{
		{
			Tag: "#canonicalTag1",
			Values: []internal.ResultValue{
				{FilePath: fileName, Line: "value1"},
				{FilePath: fileName, Line: "value3"},
			},
		},
		{
			Tag: "#canonicalTag2",
			Values: []internal.ResultValue{
				{FilePath: fileName, Line: "value2"},
			},
		},
	}

	// Call the function
	result := utils.ExtractTaggedLines(fileName, data, config)

	// Check the result
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ExtractTaggedLines() returned unexpected result:\nExpected: %v\nGot: %v", expected, result)
	}
}
