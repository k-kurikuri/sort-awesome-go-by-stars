package model

// Repository struct
type Repository struct {
	URL         string
	Star        int
	Description string
}

// Repositories Repository list
type Repositories []Repository
