package model

// Profile is a value object for user profile.
type Profile struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}
