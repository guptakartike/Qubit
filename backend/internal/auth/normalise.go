package auth

import "strings"

func NormaliseEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}
