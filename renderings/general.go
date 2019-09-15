package renderings

type GeneralResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	// Token   string `json:"token,omitempty"`
}
