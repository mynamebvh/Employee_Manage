package request

type MessageResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PaginationResponse struct {
	Success  bool        `json:"success"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
	Current  int         `json:"page_current"`
	Total    int         `json:"page_total"`
	PageSize int         `json:"page_size"`
}
