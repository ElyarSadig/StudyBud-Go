package encryption

import "testing"

func TestEncryption(t *testing.T) {
	secretKey := "MYSECRETKEYFORTE"
	aes, err := NewAES[int]([]byte(secretKey))
	if err != nil {
		t.Fatal("unexpected error happened:", err)
	}
	data := 123456
	encryptOutPut, err := aes.Encrypt(data)
	if err != nil {
		t.Fatal("unexpected error happened:", err)
	}
	decryptOutput, err := aes.Decrypt(encryptOutPut)
	if err != nil {
		t.Fatal("unexpected error happened:", err)
	}
	if decryptOutput != data {
		t.Errorf("expected decryptOutput to be equal to %d, but got %d", data, decryptOutput)
	}
}
