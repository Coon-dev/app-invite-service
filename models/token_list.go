package models

import "time"

type TokenListResponse struct {
	TokenList []TokenList `json:"token_list"`
}

type TokenList struct {
	Token     string    `json:"token"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`
}
