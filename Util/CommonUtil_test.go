package util

import (
	"bytes"
	"testing"
)

func TestAesEncryptDecrypt(t *testing.T) {
	key := []byte("0123456789123456")
	iv := []byte("0123456789123456")
	plaintext := []byte("Hello, World!")

	// Test encryption
	ciphertext, err := AesEncrypt(plaintext, key, iv)
	if err != nil {
		t.Fatalf("AesEncrypt failed: %v", err)
	}

	// Test decryption
	decryptedText, err := AesDecrypt(ciphertext, key, iv)
	if err != nil {
		t.Fatalf("AesDecrypt failed: %v", err)
	}

	// Check if decrypted text matches the original plaintext
	if !bytes.Equal(decryptedText, plaintext) {
		t.Errorf("Decrypted text does not match original plaintext. Got %s, want %s", decryptedText, plaintext)
	}
}