package entities

import "time"

type MstURL struct {
	ID          int        `json:"id"`
	OriginalURL string     `json:"original_url"`
	ShorterURL  string     `json:"short_url"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func(MstURL) TableName() string {
	return "mst_url"
}
