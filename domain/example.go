package domain

// Example the response of an example resource
type Example struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}
