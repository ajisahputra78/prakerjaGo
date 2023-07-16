package controller

type APIResponse struct {
	Kode    int16       `json:"kode"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
