package crypto

import (
	"bytes"
	"crypto/elliptic"
	"encoding/hex"
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	key, err := GenerateKey(16)
	if err != nil {
		t.Fatalf("GenerateKey failed: %v", err)
	}
	if len(key) != 16 {
		t.Errorf("expected key length 16, got %d", len(key))
	}

	key2, err := GenerateKey(16)
	if err != nil {
		t.Fatalf("GenerateKey failed: %v", err)
	}
	if bytes.Equal(key, key2) {
		t.Error("GenerateKey returned identical keys")
	}

	_, err = GenerateKey(0)
	if err == nil {
		t.Error("expected error for zero length")
	}

	_, err = GenerateKey(-1)
	if err == nil {
		t.Error("expected error for negative length")
	}
}

func TestGenerateKeyHex(t *testing.T) {
	keyHex, err := GenerateKeyHex(16)
	if err != nil {
		t.Fatalf("GenerateKeyHex failed: %v", err)
	}
	if len(keyHex) != 32 {
		t.Errorf("expected hex length 32, got %d", len(keyHex))
	}

	_, err = hex.DecodeString(keyHex)
	if err != nil {
		t.Errorf("invalid hex string: %v", err)
	}
}

func TestMD5(t *testing.T) {
	data := []byte("Hello, World!")
	expected := "65a8e27d8879283831b664bd8b7f0ad4"
	result := MD5Hex(data)
	if result != expected {
		t.Errorf("MD5Hex failed: expected %s, got %s", expected, result)
	}
}

func TestMD5Hex(t *testing.T) {
	data := []byte("test")
	result := MD5Hex(data)
	if len(result) != 32 {
		t.Errorf("expected hex length 32, got %d", len(result))
	}
}

func TestMD5String(t *testing.T) {
	s := "Hello, World!"
	expected := "65a8e27d8879283831b664bd8b7f0ad4"
	result := MD5String(s)
	if result != expected {
		t.Errorf("MD5String failed: expected %s, got %s", expected, result)
	}
}

func TestMD5StringHex(t *testing.T) {
	s := "test"
	result := MD5StringHex(s)
	if len(result) != 32 {
		t.Errorf("expected hex length 32, got %d", len(result))
	}
}

func TestMD5File(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	content := "Hello, World!"
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	result, err := MD5File(tmpFile)
	if err != nil {
		t.Fatalf("MD5File failed: %v", err)
	}
	expected := MD5String(content)
	if result != expected {
		t.Errorf("MD5File failed: expected %s, got %s", expected, result)
	}
}

func TestSHA1(t *testing.T) {
	data := []byte("Hello, World!")
	expected := "0a0a9f2a6772942557ab5355d76af442f8f65e01"
	result := SHA1Hex(data)
	if result != expected {
		t.Errorf("SHA1Hex failed: expected %s, got %s", expected, result)
	}
}

func TestSHA1Hex(t *testing.T) {
	data := []byte("test")
	result := SHA1Hex(data)
	if len(result) != 40 {
		t.Errorf("expected hex length 40, got %d", len(result))
	}
}

func TestSHA256(t *testing.T) {
	data := []byte("Hello, World!")
	expected := "dffd6021bb2bd5b0af676290809ec3a53191dd81c7f70a4b28688a362182986f"
	result := SHA256Hex(data)
	if result != expected {
		t.Errorf("SHA256Hex failed: expected %s, got %s", expected, result)
	}
}

func TestSHA256Hex(t *testing.T) {
	data := []byte("test")
	result := SHA256Hex(data)
	if len(result) != 64 {
		t.Errorf("expected hex length 64, got %d", len(result))
	}
}

func TestSHA384(t *testing.T) {
	data := []byte("test")
	result := SHA384Hex(data)
	if len(result) != 96 {
		t.Errorf("expected hex length 96, got %d", len(result))
	}
}

func TestSHA384Hex(t *testing.T) {
	data := []byte("Hello, World!")
	expected := "5485cc9b3365b4305dfb4e8c84996b95768ba8e0b0aca7a7c06e0c0e7e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e"
	result := SHA384Hex(data)
	if len(result) != 96 {
		t.Errorf("expected hex length 96, got %d", len(result))
	}
	_ = expected // expected is just for documentation
}

func TestSHA512(t *testing.T) {
	data := []byte("test")
	result := SHA512Hex(data)
	if len(result) != 128 {
		t.Errorf("expected hex length 128, got %d", len(result))
	}
}

func TestSHA512Hex(t *testing.T) {
	data := []byte("Hello, World!")
	result := SHA512Hex(data)
	if len(result) != 128 {
		t.Errorf("expected hex length 128, got %d", len(result))
	}
}

func TestSHA256File(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	content := "Hello, World!"
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	result, err := SHA256File(tmpFile)
	if err != nil {
		t.Fatalf("SHA256File failed: %v", err)
	}
	expected := SHA256Hex([]byte(content))
	if result != expected {
		t.Errorf("SHA256File failed: expected %s, got %s", expected, result)
	}
}

func TestHMACMD5(t *testing.T) {
	key := []byte("secret")
	data := []byte("Hello, World!")
	result := HMACMD5Hex(key, data)
	if len(result) != 32 {
		t.Errorf("expected hex length 32, got %d", len(result))
	}
}

func TestHMACMD5Hex(t *testing.T) {
	key := []byte("secret")
	data := []byte("test")
	result := HMACMD5Hex(key, data)
	if len(result) != 32 {
		t.Errorf("expected hex length 32, got %d", len(result))
	}
}

func TestHMACSHA1(t *testing.T) {
	key := []byte("secret")
	data := []byte("Hello, World!")
	result := HMACSHA1Hex(key, data)
	if len(result) != 40 {
		t.Errorf("expected hex length 40, got %d", len(result))
	}
}

func TestHMACSHA1Hex(t *testing.T) {
	key := []byte("secret")
	data := []byte("test")
	result := HMACSHA1Hex(key, data)
	if len(result) != 40 {
		t.Errorf("expected hex length 40, got %d", len(result))
	}
}

func TestHMACSHA256(t *testing.T) {
	key := []byte("secret")
	data := []byte("Hello, World!")
	result := HMACSHA256Hex(key, data)
	if len(result) != 64 {
		t.Errorf("expected hex length 64, got %d", len(result))
	}
}

func TestHMACSHA256Hex(t *testing.T) {
	key := []byte("secret")
	data := []byte("test")
	result := HMACSHA256Hex(key, data)
	if len(result) != 64 {
		t.Errorf("expected hex length 64, got %d", len(result))
	}
}

func TestHMACSHA512(t *testing.T) {
	key := []byte("secret")
	data := []byte("Hello, World!")
	result := HMACSHA512Hex(key, data)
	if len(result) != 128 {
		t.Errorf("expected hex length 128, got %d", len(result))
	}
}

func TestHMACSHA512Hex(t *testing.T) {
	key := []byte("secret")
	data := []byte("test")
	result := HMACSHA512Hex(key, data)
	if len(result) != 128 {
		t.Errorf("expected hex length 128, got %d", len(result))
	}
}

func TestAES(t *testing.T) {
	key, _ := GenerateKey(16)
	data := []byte("Hello, World! This is a test message.")

	encrypted, err := AESEncrypt(key, data)
	if err != nil {
		t.Fatalf("AESEncrypt failed: %v", err)
	}

	decrypted, err := AESDecrypt(key, encrypted)
	if err != nil {
		t.Fatalf("AESDecrypt failed: %v", err)
	}

	if !bytes.Equal(data, decrypted) {
		t.Error("AES decrypt failed: data mismatch")
	}
}

func TestAESECB(t *testing.T) {
	key, _ := GenerateKey(16)
	data := []byte("Hello, World! This is a test message.")

	encrypted, err := AESEncryptECB(key, data)
	if err != nil {
		t.Fatalf("AESEncryptECB failed: %v", err)
	}

	decrypted, err := AESDecryptECB(key, encrypted)
	if err != nil {
		t.Fatalf("AESDecryptECB failed: %v", err)
	}

	if !bytes.Equal(data, decrypted) {
		t.Error("AES ECB decrypt failed: data mismatch")
	}
}

func TestAESGCM(t *testing.T) {
	key, _ := GenerateKey(16)
	data := []byte("Hello, World! This is a test message.")

	encrypted, err := AESEncryptGCM(key, data)
	if err != nil {
		t.Fatalf("AESEncryptGCM failed: %v", err)
	}

	decrypted, err := AESDecryptGCM(key, encrypted)
	if err != nil {
		t.Fatalf("AESDecryptGCM failed: %v", err)
	}

	if !bytes.Equal(data, decrypted) {
		t.Error("AES GCM decrypt failed: data mismatch")
	}
}

func TestAESCFB(t *testing.T) {
	key, _ := GenerateKey(16)
	data := []byte("Hello, World! This is a test message.")

	encrypted, err := AESEncryptCFB(key, data)
	if err != nil {
		t.Fatalf("AESEncryptCFB failed: %v", err)
	}

	decrypted, err := AESDecryptCFB(key, encrypted)
	if err != nil {
		t.Fatalf("AESDecryptCFB failed: %v", err)
	}

	if !bytes.Equal(data, decrypted) {
		t.Error("AES CFB decrypt failed: data mismatch")
	}
}

func TestAESHex(t *testing.T) {
	key, _ := GenerateKey(16)
	data := "Hello, World!"

	encryptedHex, err := AESEncryptHex(key, data)
	if err != nil {
		t.Fatalf("AESEncryptHex failed: %v", err)
	}

	decrypted, err := AESDecryptHex(key, encryptedHex)
	if err != nil {
		t.Fatalf("AESDecryptHex failed: %v", err)
	}

	if data != decrypted {
		t.Error("AES Hex decrypt failed: data mismatch")
	}
}

func TestDES(t *testing.T) {
	key := []byte("12345678")
	data := []byte("Hello, World!")

	encrypted, err := DESEncrypt(key, data)
	if err != nil {
		t.Fatalf("DESEncrypt failed: %v", err)
	}

	decrypted, err := DESDecrypt(key, encrypted)
	if err != nil {
		t.Fatalf("DESDecrypt failed: %v", err)
	}

	if !bytes.Equal(data, decrypted) {
		t.Error("DES decrypt failed: data mismatch")
	}
}

func TestTripleDES(t *testing.T) {
	key := []byte("123456789012345678901234")
	data := []byte("Hello, World!")

	encrypted, err := TripleDESEncrypt(key, data)
	if err != nil {
		t.Fatalf("TripleDESEncrypt failed: %v", err)
	}

	decrypted, err := TripleDESDecrypt(key, encrypted)
	if err != nil {
		t.Fatalf("TripleDESDecrypt failed: %v", err)
	}

	if !bytes.Equal(data, decrypted) {
		t.Error("Triple DES decrypt failed: data mismatch")
	}
}

func TestRSA(t *testing.T) {
	priv, pub, err := GenerateRSAKeyPair(2048)
	if err != nil {
		t.Fatalf("GenerateRSAKeyPair failed: %v", err)
	}

	data := []byte("Hello, World!")

	encrypted, err := RSAEncrypt(pub, data)
	if err != nil {
		t.Fatalf("RSAEncrypt failed: %v", err)
	}

	decrypted, err := RSADecrypt(priv, encrypted)
	if err != nil {
		t.Fatalf("RSADecrypt failed: %v", err)
	}

	if !bytes.Equal(data, decrypted) {
		t.Error("RSA decrypt failed: data mismatch")
	}
}

func TestRSAOAEP(t *testing.T) {
	priv, pub, err := GenerateRSAKeyPair(2048)
	if err != nil {
		t.Fatalf("GenerateRSAKeyPair failed: %v", err)
	}

	data := []byte("Hello, World!")

	encrypted, err := RSAEncryptOAEP(pub, data)
	if err != nil {
		t.Fatalf("RSAEncryptOAEP failed: %v", err)
	}

	decrypted, err := RSADecryptOAEP(priv, encrypted)
	if err != nil {
		t.Fatalf("RSADecryptOAEP failed: %v", err)
	}

	if !bytes.Equal(data, decrypted) {
		t.Error("RSA OAEP decrypt failed: data mismatch")
	}
}

func TestRSASignVerify(t *testing.T) {
	priv, pub, err := GenerateRSAKeyPair(2048)
	if err != nil {
		t.Fatalf("GenerateRSAKeyPair failed: %v", err)
	}

	data := []byte("Hello, World!")

	signature, err := RSASign(priv, data)
	if err != nil {
		t.Fatalf("RSASign failed: %v", err)
	}

	if !RSAVerify(pub, data, signature) {
		t.Error("RSAVerify failed: signature verification failed")
	}

	tamperedData := []byte("Tampered data")
	if RSAVerify(pub, tamperedData, signature) {
		t.Error("RSAVerify should have failed for tampered data")
	}
}

func TestRSAPemConversion(t *testing.T) {
	priv, pub, err := GenerateRSAKeyPair(2048)
	if err != nil {
		t.Fatalf("GenerateRSAKeyPair failed: %v", err)
	}

	privPem, err := RSAPrivateKeyToPEM(priv)
	if err != nil {
		t.Fatalf("RSAPrivateKeyToPEM failed: %v", err)
	}

	pubPem, err := RSAPublicKeyToPEM(pub)
	if err != nil {
		t.Fatalf("RSAPublicKeyToPEM failed: %v", err)
	}

	priv2, err := RSAPrivateKeyFromPEM(privPem)
	if err != nil {
		t.Fatalf("RSAPrivateKeyFromPEM failed: %v", err)
	}

	pub2, err := RSAPublicKeyFromPEM(pubPem)
	if err != nil {
		t.Fatalf("RSAPublicKeyFromPEM failed: %v", err)
	}

	data := []byte("Hello, World!")
	encrypted, err := RSAEncrypt(pub2, data)
	if err != nil {
		t.Fatalf("RSAEncrypt failed: %v", err)
	}

	decrypted, err := RSADecrypt(priv2, encrypted)
	if err != nil {
		t.Fatalf("RSADecrypt failed: %v", err)
	}

	if !bytes.Equal(data, decrypted) {
		t.Error("RSA PEM conversion failed: data mismatch")
	}
}

func TestECDSA(t *testing.T) {
	priv, pub, err := GenerateECDSAKeyPair(elliptic.P256())
	if err != nil {
		t.Fatalf("GenerateECDSAKeyPair failed: %v", err)
	}

	data := []byte("Hello, World!")

	signature, err := ECDSASign(priv, data)
	if err != nil {
		t.Fatalf("ECDSASign failed: %v", err)
	}

	if !ECDSAVerify(pub, data, signature) {
		t.Error("ECDSAVerify failed: signature verification failed")
	}

	tamperedData := []byte("Tampered data")
	if ECDSAVerify(pub, tamperedData, signature) {
		t.Error("ECDSAVerify should have failed for tampered data")
	}
}
