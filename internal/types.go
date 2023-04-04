package internal

type TaggedLine struct {
	Tag    string        `json:"tag"`
	Values []ResultValue `json:"values"`
}

type ResultValue struct {
	FilePath string `json:"file_path"`
	Line     string `json:"line"`
}

type Config struct {
	TagMappings map[string]string `json:"tag_mappings"`
}
