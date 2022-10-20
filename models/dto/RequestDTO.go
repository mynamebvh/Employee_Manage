package dto

type QueryGetRequest struct {
	Type       string `json:"type"`
	Content    string `json:"content"`
	FullName   string `json:"full_name"`
	Status     string `json:"status"`
	ApprovedBy string `json:"approved_by"`
}
