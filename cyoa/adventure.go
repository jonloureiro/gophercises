package cyoa

// Adventure doc
type Adventure struct {
	Slug    string   `json:"slug"`
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

// Option doc
type Option struct {
	Text string `json:"text"`
	Slug string `json:"arc"`
}
