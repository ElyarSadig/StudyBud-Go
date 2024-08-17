package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"errors"
	"io"
)

type AES[T any] struct {
	block cipher.Block
	gcm   cipher.AEAD
}

func NewAES[T any](key []byte) (*AES[T], error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, errors.New("key length must be 16, 24, or 32 bytes")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return &AES[T]{
		block: block,
		gcm:   gcm,
	}, nil
}

func (a *AES[T]) Encrypt(data T) ([]byte, error) {
	plaintext, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, a.gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	ciphertext := a.gcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

func (a *AES[T]) Decrypt(data []byte) (T, error) {
	var result T
	nonceSize := a.gcm.NonceSize()
	if len(data) < nonceSize {
		return result, errors.New("ciphertext too short")
	}
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := a.gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(plaintext, &result); err != nil {
		return result, err
	}
	return result, nil
}
