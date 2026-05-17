package dto

type CreatePengurusSosmedRequest struct {
	Url string `json:"url" validate:"omitempty,socialurl"`
}

type UpdatePengurusSosmedRequest struct {
	ID  *int   `json:"id" validate:"omitempty,numeric"`
	Url string `json:"url"`
}

type CreatePengurusSosmedPayload struct {
	PengurusID int    `json:"pengurus_id"`
	Platform   string `json:"platform"`
	Username   string `json:"username"`
	Url        string `json:"url"`
	IsPrimary  bool   `json:"is_primary"`
}

type PengurusSosmedResponse struct {
	Platform  string `json:"platform"`
	Username  string `json:"username"`
	Url       string `json:"url"`
	IsPrimary bool   `json:"is_primary"`
}
