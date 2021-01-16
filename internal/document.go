package internal

// Document schema
type Document struct {
	Content   string `json:"content"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Topic     string `json:"topic"`
	Watermark string `json:"watermark,omitempty"`
}

// Filter for search
type Filter struct {
	Key string `json:"key"`

	// If value is empty, just return everything but sorted with the key
	Value string `json:"value,omitempty"`
}

// Status of query
type Status string

// Status
const (
	Pending    Status = "Pending"
	Started    Status = "Started"
	InProgress Status = "InProgress"
	Finished   Status = "Finished"
	Failed     Status = "Failed"
)
