package inventory

type Image struct {
	ID  string `json:"id" toml:"id" yaml:"id"`
	URL string `json:"url" toml:"url" yaml:"url"`
}

type Images []Image
