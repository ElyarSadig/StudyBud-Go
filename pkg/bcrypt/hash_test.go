package bcrypt

import "testing"

func TestHashPassword(t *testing.T) {
	password := "test1234"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatal("unexpected error happened:", err)
	}
	ok := CheckPasswordHash(password, hashedPassword)
	if !ok {
		t.Fatal("expected hashedPassword to be equal to password")
	}
}
