package test

import (
	"server/app-invite-service/configs"
	"server/app-invite-service/services"
	"testing"
)

func TestAuthService(t *testing.T) {
	if services.AuthService("") != false {
		t.Error("case 1 empty auth: expected false")
	}
	if services.AuthService("Hello World") != false {
		t.Error("case 2 wrong auth: expected false")
	}

	if services.AuthService(configs.AuthKey) != true {
		t.Error("case 3 correct auth: expected true")
	}
}
