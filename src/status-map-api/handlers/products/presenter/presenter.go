package presenter

import "time"

type Product struct {
	ID           string    `json:"id"`
	Etag         string    `json:"etag,omitempty"`
	Org          string    `json:"org"`
	Platform     string    `json:"platform"`
	Infras       []string  `json:"infras"`
	Stages       []string  `json:"stages"`
	StatusInfras StatusMap `json:"status_infras"`
	StatusStages StatusMap `json:"status_stages"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type StatusMap map[string]Status

type Status struct {
	CurrentStatus string    `json:"current_status"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Products struct {
	Kind  string    `json:"kind"`
	Items []Product `json:"items"`
}

type Platform struct {
	Kind  string    `json:"kind"`
	Name  string    `json:"name"`
	Items []Product `json:"items"`
}

type ProductsGroupedByPlatforms struct {
	Kind  string     `json:"kind"`
	Items []Platform `json:"items"`
}
