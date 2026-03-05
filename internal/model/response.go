package model

type ResponseMessage struct {
	Message string `json:"message"`
	Meta    Meta   `json:"meta"`
}

type ResponseBody struct {
	Response any  `json:"result"`
	Meta     Meta `json:"meta"`
}

type PaginationResponse struct {
	Result any            `json:"result"`
	Meta   PaginationMeta `json:"meta"`
}

type PaginationMeta struct {
	TotalData int    `json:"total_data"`
	Page      int    `json:"page"`
	Size      int    `json:"size"`
	Previous  string `json:"previous"`
	Next      string `json:"next"`
}

type Meta struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}
