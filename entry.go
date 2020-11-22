package rest

func NewEmbeddedEntry(key, name, path, description string) *EmbeddedEntry {
	return &EmbeddedEntry{name, key, path, description}
}

type EmbeddedEntry struct {
	Name string
	// Anyway, an identifier for routers.
	Key string

	Path string

	Description string
}

type IRestEntry interface {
	GetPath() string
	GetTag() string
}
