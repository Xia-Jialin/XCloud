package es

type Metadata struct {
	Name    string
	Version int
	Size    int64
	Hash    string
}

func SearchAllVersions(name string, from, size int) ([]Metadata, error) {
	return nil, nil
}
