package types

type PostBody struct {
	Calculation string `json:"calculation"`
	CreatedAt   string `json:"createdAt"`
}

type ResponsStruct struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseHistory struct {
	ID          int    `json:"id"`
	CreatedAt   string `json:"createdAt"`
	Calculation string `json:"calculation"`
}
