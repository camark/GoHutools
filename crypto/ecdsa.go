package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"math/big"
)

// ECDSAPrivateKey is ECDSA private key wrapper
type ECDSAPrivateKey struct {
	Key *ecdsa.PrivateKey
}

// ECDSAPublicKey is ECDSA public key wrapper
type ECDSAPublicKey struct {
	Key *ecdsa.PublicKey
}

// GenerateECDSAKeyPair generates ECDSA key pair
func GenerateECDSAKeyPair(curve elliptic.Curve) (*ECDSAPrivateKey, *ECDSAPublicKey, error) {
	if curve == nil {
		return nil, nil, errors.New("curve is nil")
	}

	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	return &ECDSAPrivateKey{Key: privateKey}, &ECDSAPublicKey{Key: &privateKey.PublicKey}, nil
}

// ECDSASign signs data with ECDSA private key
func ECDSASign(priv *ECDSAPrivateKey, data []byte) ([]byte, error) {
	if priv == nil || priv.Key == nil {
		return nil, errors.New("private key is nil")
	}

	hashed := sha256.Sum256(data)
	r, s, err := ecdsa.Sign(rand.Reader, priv.Key, hashed[:])
	if err != nil {
		return nil, err
	}

	// Encode r and s as fixed-length byte slices
	byteLen := (priv.Key.Curve.Params().BitSize + 7) / 8
	signature := make([]byte, 2*byteLen)
	r.FillBytes(signature[:byteLen])
	s.FillBytes(signature[byteLen:])

	return signature, nil
}

// ECDSAVerify verifies ECDSA signature
func ECDSAVerify(pub *ECDSAPublicKey, data, signature []byte) bool {
	if pub == nil || pub.Key == nil {
		return false
	}

	byteLen := (pub.Key.Curve.Params().BitSize + 7) / 8
	if len(signature) != 2*byteLen {
		return false
	}

	r := new(big.Int).SetBytes(signature[:byteLen])
	s := new(big.Int).SetBytes(signature[byteLen:])

	hashed := sha256.Sum256(data)
	return ecdsa.Verify(pub.Key, hashed[:], r, s)
}
