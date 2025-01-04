package models

// APIResponse is the standard response structure for the API.
type APIResponse struct {
	Data       interface{} `json:"data,omitempty"`       // The main response data (can be list or single resource).
	Pagination *Pagination `json:"pagination,omitempty"` // Pagination details (optional, for paginated responses).
	Message    string      `json:"message,omitempty"`    // Message for additional context (optional).
	Error      string      `json:"error,omitempty"`      // Error message if any error occurs.
}
