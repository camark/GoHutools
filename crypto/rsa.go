package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// RSAPrivateKey is RSA private key wrapper
type RSAPrivateKey struct {
	Key *rsa.PrivateKey
}

// RSAPublicKey is RSA public key wrapper
type RSAPublicKey struct {
	Key *rsa.PublicKey
}

// GenerateRSAKeyPair generates RSA key pair
func GenerateRSAKeyPair(bits int) (*RSAPrivateKey, *RSAPublicKey, error) {
	if bits < 2048 {
		return nil, nil, errors.New("RSA key size must be at least 2048 bits")
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}

	return &RSAPrivateKey{Key: privateKey}, &RSAPublicKey{Key: &privateKey.PublicKey}, nil
}

// RSAPrivateKeyFromPEM parses PEM private key
func RSAPrivateKeyFromPEM(pemBytes []byte) (*RSAPrivateKey, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		// Try PKCS8 format
		keyInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		var ok bool
		key, ok = keyInterface.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("not an RSA private key")
		}
	}

	return &RSAPrivateKey{Key: key}, nil
}

// RSAPublicKeyFromPEM parses PEM public key
func RSAPublicKeyFromPEM(pemBytes []byte) (*RSAPublicKey, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}

	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaKey, ok := key.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not an RSA public key")
	}

	return &RSAPublicKey{Key: rsaKey}, nil
}

// RSAPrivateKeyToPEM exports private key to PEM
func RSAPrivateKeyToPEM(key *RSAPrivateKey) ([]byte, error) {
	if key == nil || key.Key == nil {
		return nil, errors.New("private key is nil")
	}

	keyBytes := x509.MarshalPKCS1PrivateKey(key.Key)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: keyBytes,
	}
	return pem.EncodeToMemory(block), nil
}

// RSAPublicKeyToPEM exports public key to PEM
func RSAPublicKeyToPEM(key *RSAPublicKey) ([]byte, error) {
	if key == nil || key.Key == nil {
		return nil, errors.New("public key is nil")
	}

	keyBytes, err := x509.MarshalPKIXPublicKey(key.Key)
	if err != nil {
		return nil, err
	}
	block := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: keyBytes,
	}
	return pem.EncodeToMemory(block), nil
}

// RSAEncrypt encrypts with public key
func RSAEncrypt(pub *RSAPublicKey, data []byte) ([]byte, error) {
	if pub == nil || pub.Key == nil {
		return nil, errors.New("public key is nil")
	}
	return rsa.EncryptPKCS1v15(rand.Reader, pub.Key, data)
}

// RSADecrypt decrypts with private key
func RSADecrypt(priv *RSAPrivateKey, encrypted []byte) ([]byte, error) {
	if priv == nil || priv.Key == nil {
		return nil, errors.New("private key is nil")
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv.Key, encrypted)
}

// RSASign signs data with private key
func RSASign(priv *RSAPrivateKey, data []byte) ([]byte, error) {
	if priv == nil || priv.Key == nil {
		return nil, errors.New("private key is nil")
	}

	hashed := sha256.Sum256(data)
	signature, err := rsa.SignPKCS1v15(rand.Reader, priv.Key, crypto.SHA256, hashed[:])
	if err != nil {
		return nil, err
	}
	return signature, nil
}

// RSAVerify verifies signature with public key
func RSAVerify(pub *RSAPublicKey, data, signature []byte) bool {
	if pub == nil || pub.Key == nil {
		return false
	}

	hashed := sha256.Sum256(data)
	err := rsa.VerifyPKCS1v15(pub.Key, crypto.SHA256, hashed[:], signature)
	return err == nil
}

// RSAEncryptOAEP encrypts with OAEP padding
func RSAEncryptOAEP(pub *RSAPublicKey, data []byte) ([]byte, error) {
	if pub == nil || pub.Key == nil {
		return nil, errors.New("public key is nil")
	}
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, pub.Key, data, nil)
}

// RSADecryptOAEP decrypts with OAEP padding
func RSADecryptOAEP(priv *RSAPrivateKey, encrypted []byte) ([]byte, error) {
	if priv == nil || priv.Key == nil {
		return nil, errors.New("private key is nil")
	}
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, priv.Key, encrypted, nil)
}
