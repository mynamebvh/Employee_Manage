package request

type NewRequest struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}
