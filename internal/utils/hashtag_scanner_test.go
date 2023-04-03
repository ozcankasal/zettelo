package utils_test

import (
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
