package request

type NewRequest struct {
	Type    string `json:"type" validate:"required"`
	Content string `json:"content" validate:"required"`
}
