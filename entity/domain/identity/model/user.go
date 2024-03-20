package model

// User is an entity for user.
type User struct {
	ID string `json:"id,omitempty"`

	Profile *Profile `json:"profile,omitempty"`
}
