package auth

import (
	"testing"
)

func TestHashAndVerifyPassword(t *testing.T) {
	password := "correct-password"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPasswod() error = %v", err)
	}

	if password == hash {
		t.Fatal("password hash must not equal plaintext password")
	}

	if !VerifyPassword(password, hash) {
		t.Fatalf("VerifyPassword() must accept correct password")
	}
	if VerifyPassword("wrong-password", hash) {

		t.Fatal("VerifyPassword() returned true for incorrect password")
	}

}

func TestVerifyPasswordRejectsInvalidHash(t *testing.T) {

	if VerifyPassword("password", "not-a-valid-argon2-hash") {
		t.Fatal("VerifyPassword() accepted an invalid hash")
	}

}
