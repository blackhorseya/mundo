package model

// Collection is an entity for word collection.
type Collection struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	OwnerID string   `json:"owner_id"`
	Words   []string `json:"words"`
}
