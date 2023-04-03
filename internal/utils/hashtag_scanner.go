package utils

import "github.com/ozcankasal/zettelo/internal"

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
