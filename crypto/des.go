package crypto

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"errors"
	"io"
)

// DESEncrypt encrypts data with DES-CBC
func DESEncrypt(key, data []byte) ([]byte, error) {
	if len(key) != 8 {
		return nil, errors.New("DES key must be 8 bytes")
	}

	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// PKCS7 padding
	padding := des.BlockSize - len(data)%des.BlockSize
	padtext := make([]byte, len(data)+padding)
	copy(padtext, data)
	for i := len(data); i < len(padtext); i++ {
		padtext[i] = byte(padding)
	}

	ciphertext := make([]byte, des.BlockSize+len(padtext))
	iv := ciphertext[:des.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[des.BlockSize:], padtext)

	return ciphertext, nil
}

// DESDecrypt decrypts data with DES-CBC
func DESDecrypt(key, encrypted []byte) ([]byte, error) {
	if len(key) != 8 {
		return nil, errors.New("DES key must be 8 bytes")
	}

	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(encrypted) < des.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := encrypted[:des.BlockSize]
	encrypted = encrypted[des.BlockSize:]

	if len(encrypted)%des.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(encrypted))
	mode.CryptBlocks(decrypted, encrypted)

	// Remove PKCS7 padding
	if len(decrypted) == 0 {
		return nil, errors.New("decrypted data is empty")
	}
	padding := int(decrypted[len(decrypted)-1])
	if padding > des.BlockSize || padding == 0 {
		return nil, errors.New("invalid padding")
	}
	for i := len(decrypted) - padding; i < len(decrypted); i++ {
		if decrypted[i] != byte(padding) {
			return nil, errors.New("invalid padding")
		}
	}
	decrypted = decrypted[:len(decrypted)-padding]

	return decrypted, nil
}

// TripleDESEncrypt encrypts data with 3DES-CBC
func TripleDESEncrypt(key, data []byte) ([]byte, error) {
	if len(key) != 24 {
		return nil, errors.New("triple DES key must be 24 bytes")
	}

	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}

	// PKCS7 padding
	padding := des.BlockSize - len(data)%des.BlockSize
	padtext := make([]byte, len(data)+padding)
	copy(padtext, data)
	for i := len(data); i < len(padtext); i++ {
		padtext[i] = byte(padding)
	}

	ciphertext := make([]byte, des.BlockSize+len(padtext))
	iv := ciphertext[:des.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[des.BlockSize:], padtext)

	return ciphertext, nil
}

// TripleDESDecrypt decrypts data with 3DES-CBC
func TripleDESDecrypt(key, encrypted []byte) ([]byte, error) {
	if len(key) != 24 {
		return nil, errors.New("triple DES key must be 24 bytes")
	}

	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}

	if len(encrypted) < des.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := encrypted[:des.BlockSize]
	encrypted = encrypted[des.BlockSize:]

	if len(encrypted)%des.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(encrypted))
	mode.CryptBlocks(decrypted, encrypted)

	// Remove PKCS7 padding
	if len(decrypted) == 0 {
		return nil, errors.New("decrypted data is empty")
	}
	padding := int(decrypted[len(decrypted)-1])
	if padding > des.BlockSize || padding == 0 {
		return nil, errors.New("invalid padding")
	}
	for i := len(decrypted) - padding; i < len(decrypted); i++ {
		if decrypted[i] != byte(padding) {
			return nil, errors.New("invalid padding")
		}
	}
	decrypted = decrypted[:len(decrypted)-padding]

	return decrypted, nil
}
