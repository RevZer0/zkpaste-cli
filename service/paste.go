package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hkdf"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
)

func EncryptPaste(paste string, password string) (ciphertext, iv, key, signature []byte) {
	iv = make([]byte, 12)
	key = make([]byte, 32)

	rand.Read(iv)
	rand.Read(key)
	encryptionKey := key

	if len(password) > 0 {
		encryptionKey = deriveFromPassword(key, password)
	}

	aes, _ := aes.NewCipher(encryptionKey)
	gcm, _ := cipher.NewGCM(aes)

	ciphertext = gcm.Seal(nil, iv, []byte(paste), nil)
	signature = generateSignature(encryptionKey, paste)

	return
}

func DecryptPaste(ciphertext, iv, key []byte, password string) (string, error) {
	if len(password) > 0 {
		key = deriveFromPassword(key, password)
	}
	aes, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		return "", err
	}

	plaintext, err := gcm.Open(nil, iv, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func ProofOfKnowlege(key []byte, plaintext string, password string) []byte {
	if len(password) > 0 {
		key = deriveFromPassword(key, password)
	}
	return generateSignature(key, plaintext)
}

func generateSignature(key []byte, plaintext string) []byte {
	sign := hmac.New(sha256.New, key)
	sign.Write([]byte(plaintext))
	return sign.Sum(nil)
}

func deriveFromPassword(ikm []byte, password string) []byte {
	key, _ := hkdf.Key(sha256.New, ikm, []byte(password), "", sha256.New().Size())
	return key
}
