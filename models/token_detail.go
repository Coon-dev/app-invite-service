package models

type TokenDetailRequest struct {
	Token string `json:"token"`
}

type TokenDetailResponse struct {
	Status string `json:"status"`
}
