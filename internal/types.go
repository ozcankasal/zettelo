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
	TagMappings map[string]string `yaml:"tag_mappings"`
}

type TagList []TaggedLine

func (t TagList) Len() int {
	return len(t)
}

func (t TagList) Less(i, j int) bool {
	return t[i].Tag < t[j].Tag
}

func (t TagList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
