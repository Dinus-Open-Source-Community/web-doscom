package dto

import "encoding/json"

type HistoryTimelineResponse struct {
	ID           int             `json:"id"`
	IDAuthor     int             `json:"id_author"`
	Title        string          `json:"title"`
	Year         string          `json:"year"`
	Description  string          `json:"description"`
	DisplayOrder int             `json:"display_order"`
	Photos       json.RawMessage `gorm:"column:photos" json:"photos"`
}

type HistoryPayload struct {
	AuthorID     int    `json:"author_id"`
	Title        string `json:"title"`
	Year         string `json:"year"`
	Description  string `json:"description"`
	DisplayOrder int    `json:"display_order"`
}
