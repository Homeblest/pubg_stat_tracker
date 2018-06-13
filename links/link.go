package links

// Link describes the properties of a link object
type Link struct {
	Self   string `json:"self"`
	Schema string `json:"schema,omitempty"`
}
