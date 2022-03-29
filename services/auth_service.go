package services

import "server/app-invite-service/configs"

func AuthService(auth string) bool {
	if auth != configs.AuthKey {
		return false
	}
	return true
}
