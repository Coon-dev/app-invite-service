package test

import (
	"server/app-invite-service/utils"
	"testing"
)

func TestRandomToken(t *testing.T) {
	alphanumeric := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	length := 7
	token := utils.RandomToken(length)
	if len(token) != length {
		t.Errorf("case 1 length: expected %d, found: %d", length, len(token))
	}

	token_rune := []rune(token)
	for i := range token_rune {
		found := false
		for _, al := range alphanumeric {
			if token_rune[i] == al {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("case 2 alphanumeric: expected all alphanumberic, found: %s, index %d", token, i)
		}
	}

}
