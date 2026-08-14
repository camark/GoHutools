---
title: crypto
---

# crypto

加密工具包，提供 MD5、SHA、HMAC、AES、DES、RSA、ECDSA 等常用加密算法的封装。

## 导入

```go
import "github.com/gongm/gohutool/crypto"
```

## 通用函数

### GenerateKey

```go
func GenerateKey(length int) ([]byte, error)
```

生成指定长度的随机密钥。

**示例:**

```go
key, err := crypto.GenerateKey(32) // 256-bit key
```

### GenerateKeyHex

```go
func GenerateKeyHex(length int) (string, error)
```

生成指定长度的随机密钥并返回十六进制字符串。

**示例:**

```go
keyHex, err := crypto.GenerateKeyHex(32)
fmt.Println(keyHex)
```

## MD5 函数

### MD5

```go
func MD5(data []byte) []byte
```

计算字节数据的 MD5 哈希值。

**示例:**

```go
hash := crypto.MD5([]byte("Hello"))
```

### MD5Hex

```go
func MD5Hex(data []byte) string
```

计算字节数据的 MD5 哈希值，返回十六进制字符串。

**示例:**

```go
hash := crypto.MD5Hex([]byte("Hello"))
fmt.Println(hash) // 8b1a9953c4611296a827abf8c47804d7
```

### MD5String

```go
func MD5String(s string) string
```

计算字符串的 MD5 哈希值，返回十六进制字符串。

**示例:**

```go
hash := crypto.MD5String("Hello")
```

### MD5StringHex

```go
func MD5StringHex(s string) string
```

计算字符串的 MD5 哈希值，返回十六进制字符串（等同于 MD5String）。

**示例:**

```go
hash := crypto.MD5StringHex("Hello")
```

### MD5File

```go
func MD5File(path string) (string, error)
```

计算文件的 MD5 哈希值。

**示例:**

```go
hash, err := crypto.MD5File("/path/to/file")
fmt.Println(hash)
```

## SHA 函数

### SHA1

```go
func SHA1(data []byte) []byte
```

计算 SHA-1 哈希值。

**示例:**

```go
hash := crypto.SHA1([]byte("Hello"))
```

### SHA1Hex

```go
func SHA1Hex(data []byte) string
```

计算 SHA-1 哈希值，返回十六进制字符串。

**示例:**

```go
hash := crypto.SHA1Hex([]byte("Hello"))
```

### SHA256

```go
func SHA256(data []byte) []byte
```

计算 SHA-256 哈希值。

**示例:**

```go
hash := crypto.SHA256([]byte("Hello"))
```

### SHA256Hex

```go
func SHA256Hex(data []byte) string
```

计算 SHA-256 哈希值，返回十六进制字符串。

**示例:**

```go
hash := crypto.SHA256Hex([]byte("Hello"))
```

### SHA384

```go
func SHA384(data []byte) []byte
```

计算 SHA-384 哈希值。

**示例:**

```go
hash := crypto.SHA384([]byte("Hello"))
```

### SHA384Hex

```go
func SHA384Hex(data []byte) string
```

计算 SHA-384 哈希值，返回十六进制字符串。

### SHA512

```go
func SHA512(data []byte) []byte
```

计算 SHA-512 哈希值。

### SHA512Hex

```go
func SHA512Hex(data []byte) string
```

计算 SHA-512 哈希值，返回十六进制字符串。

### SHA256File

```go
func SHA256File(path string) (string, error)
```

计算文件的 SHA-256 哈希值。

**示例:**

```go
hash, err := crypto.SHA256File("/path/to/file")
```

## HMAC 函数

### HMACMD5

```go
func HMACMD5(key, data []byte) []byte
```

计算 HMAC-MD5 签名。

### HMACMD5Hex

```go
func HMACMD5Hex(key, data []byte) string
```

计算 HMAC-MD5 签名，返回十六进制字符串。

**示例:**

```go
sig := crypto.HMACMD5Hex([]byte("secret"), []byte("message"))
```

### HMACSHA1

```go
func HMACSHA1(key, data []byte) []byte
```

计算 HMAC-SHA1 签名。

### HMACSHA1Hex

```go
func HMACSHA1Hex(key, data []byte) string
```

计算 HMAC-SHA1 签名，返回十六进制字符串。

**示例:**

```go
sig := crypto.HMACSHA1Hex([]byte("secret"), []byte("message"))
```

### HMACSHA256

```go
func HMACSHA256(key, data []byte) []byte
```

计算 HMAC-SHA256 签名。

### HMACSHA256Hex

```go
func HMACSHA256Hex(key, data []byte) string
```

计算 HMAC-SHA256 签名，返回十六进制字符串。

**示例:**

```go
sig := crypto.HMACSHA256Hex([]byte("secret"), []byte("message"))
```

### HMACSHA512

```go
func HMACSHA512(key, data []byte) []byte
```

计算 HMAC-SHA512 签名。

### HMACSHA512Hex

```go
func HMACSHA512Hex(key, data []byte) string
```

计算 HMAC-SHA512 签名，返回十六进制字符串。

## AES 函数

### AESEncrypt

```go
func AESEncrypt(key, data []byte) ([]byte, error)
```

使用 AES-CBC 模式加密数据，密钥长度为 16、24 或 32 字节。使用 PKCS7 填充，随机 IV 前置到密文。

**示例:**

```go
key := []byte("0123456789abcdef") // 16 bytes = AES-128
encrypted, err := crypto.AESEncrypt(key, []byte("Hello, World!"))
```

### AESDecrypt

```go
func AESDecrypt(key, encrypted []byte) ([]byte, error)
```

使用 AES-CBC 模式解密数据。

**示例:**

```go
decrypted, err := crypto.AESDecrypt(key, encrypted)
fmt.Println(string(decrypted))
```

### AESEncryptECB

```go
func AESEncryptECB(key, data []byte) ([]byte, error)
```

使用 AES-ECB 模式加密数据。使用 PKCS7 填充。

**示例:**

```go
encrypted, err := crypto.AESEncryptECB(key, []byte("Hello"))
```

### AESDecryptECB

```go
func AESDecryptECB(key, encrypted []byte) ([]byte, error)
```

使用 AES-ECB 模式解密数据。

### AESEncryptGCM

```go
func AESEncryptGCM(key, data []byte) ([]byte, error)
```

使用 AES-GCM 模式加密数据（认证加密）。随机 nonce 前置到密文。

**示例:**

```go
encrypted, err := crypto.AESEncryptGCM(key, []byte("Hello, World!"))
```

### AESDecryptGCM

```go
func AESDecryptGCM(key, encrypted []byte) ([]byte, error)
```

使用 AES-GCM 模式解密数据。

**示例:**

```go
decrypted, err := crypto.AESDecryptGCM(key, encrypted)
```

### AESEncryptCFB

```go
func AESEncryptCFB(key, data []byte) ([]byte, error)
```

使用 AES-CFB 模式加密数据。随机 IV 前置到密文。

### AESDecryptCFB

```go
func AESDecryptCFB(key, encrypted []byte) ([]byte, error)
```

使用 AES-CFB 模式解密数据。

### AESEncryptHex

```go
func AESEncryptHex(key []byte, data string) (string, error)
```

使用 AES-CBC 加密并返回十六进制字符串。

**示例:**

```go
encrypted, err := crypto.AESEncryptHex([]byte("0123456789abcdef"), "Hello, World!")
fmt.Println(encrypted)
```

### AESDecryptHex

```go
func AESDecryptHex(key []byte, hexStr string) (string, error)
```

解密十六进制编码的 AES-CBC 密文。

**示例:**

```go
decrypted, err := crypto.AESDecryptHex(key, encrypted)
fmt.Println(decrypted)
```

## DES 函数

### DESEncrypt

```go
func DESEncrypt(key, data []byte) ([]byte, error)
```

使用 DES-CBC 模式加密数据。密钥必须为 8 字节。

**示例:**

```go
key := []byte("12345678") // 8 bytes
encrypted, err := crypto.DESEncrypt(key, []byte("Hello"))
```

### DESDecrypt

```go
func DESDecrypt(key, encrypted []byte) ([]byte, error)
```

使用 DES-CBC 模式解密数据。

### TripleDESEncrypt

```go
func TripleDESEncrypt(key, data []byte) ([]byte, error)
```

使用 3DES-CBC 模式加密数据。密钥必须为 24 字节。

**示例:**

```go
key := []byte("123456789012345678901234") // 24 bytes
encrypted, err := crypto.TripleDESEncrypt(key, []byte("Hello"))
```

### TripleDESDecrypt

```go
func TripleDESDecrypt(key, encrypted []byte) ([]byte, error)
```

使用 3DES-CBC 模式解密数据。

## RSA 函数

### GenerateRSAKeyPair

```go
func GenerateRSAKeyPair(bits int) (*RSAPrivateKey, *RSAPublicKey, error)
```

生成 RSA 密钥对。密钥长度至少 2048 位。

**示例:**

```go
priv, pub, err := crypto.GenerateRSAKeyPair(2048)
```

### RSAPrivateKeyFromPEM

```go
func RSAPrivateKeyFromPEM(pemBytes []byte) (*RSAPrivateKey, error)
```

从 PEM 格式解析 RSA 私钥（支持 PKCS1 和 PKCS8）。

**示例:**

```go
priv, err := crypto.RSAPrivateKeyFromPEM(pemData)
```

### RSAPublicKeyFromPEM

```go
func RSAPublicKeyFromPEM(pemBytes []byte) (*RSAPublicKey, error)
```

从 PEM 格式解析 RSA 公钥。

### RSAPrivateKeyToPEM

```go
func RSAPrivateKeyToPEM(key *RSAPrivateKey) ([]byte, error)
```

将 RSA 私钥导出为 PEM 格式。

**示例:**

```go
pemData, err := crypto.RSAPrivateKeyToPEM(priv)
```

### RSAPublicKeyToPEM

```go
func RSAPublicKeyToPEM(key *RSAPublicKey) ([]byte, error)
```

将 RSA 公钥导出为 PEM 格式。

### RSAEncrypt

```go
func RSAEncrypt(pub *RSAPublicKey, data []byte) ([]byte, error)
```

使用 RSA 公钥加密数据（PKCS1v15 填充）。

**示例:**

```go
encrypted, err := crypto.RSAEncrypt(pub, []byte("Hello"))
```

### RSADecrypt

```go
func RSADecrypt(priv *RSAPrivateKey, encrypted []byte) ([]byte, error)
```

使用 RSA 私钥解密数据。

**示例:**

```go
decrypted, err := crypto.RSADecrypt(priv, encrypted)
fmt.Println(string(decrypted))
```

### RSASign

```go
func RSASign(priv *RSAPrivateKey, data []byte) ([]byte, error)
```

使用 RSA 私钥签名数据（SHA-256 摘要）。

**示例:**

```go
signature, err := crypto.RSASign(priv, []byte("message"))
```

### RSAVerify

```go
func RSAVerify(pub *RSAPublicKey, data, signature []byte) bool
```

使用 RSA 公钥验证签名。

**示例:**

```go
valid := crypto.RSAVerify(pub, []byte("message"), signature)
if valid {
    fmt.Println("签名验证通过")
}
```

### RSAEncryptOAEP

```go
func RSAEncryptOAEP(pub *RSAPublicKey, data []byte) ([]byte, error)
```

使用 RSA 公钥加密数据（OAEP 填充，SHA-256 哈希）。

**示例:**

```go
encrypted, err := crypto.RSAEncryptOAEP(pub, []byte("Hello"))
```

### RSADecryptOAEP

```go
func RSADecryptOAEP(priv *RSAPrivateKey, encrypted []byte) ([]byte, error)
```

使用 RSA 私钥解密 OAEP 填充的数据。

## ECDSA 函数

### GenerateECDSAKeyPair

```go
func GenerateECDSAKeyPair(curve elliptic.Curve) (*ECDSAPrivateKey, *ECDSAPublicKey, error)
```

生成 ECDSA 密钥对。支持 elliptic.P256()、elliptic.P384()、elliptic.P521() 等曲线。

**示例:**

```go
priv, pub, err := crypto.GenerateECDSAKeyPair(elliptic.P256())
```

### ECDSASign

```go
func ECDSASign(priv *ECDSAPrivateKey, data []byte) ([]byte, error)
```

使用 ECDSA 私钥签名数据（SHA-256 摘要）。签名格式为 r || s，定长字节。

**示例:**

```go
signature, err := crypto.ECDSASign(priv, []byte("message"))
```

### ECDSAVerify

```go
func ECDSAVerify(pub *ECDSAPublicKey, data, signature []byte) bool
```

使用 ECDSA 公钥验证签名。

**示例:**

```go
valid := crypto.ECDSAVerify(pub, []byte("message"), signature)
```

## 完整示例

```go
package main

import (
    "crypto/elliptic"
    "fmt"
    "log"

    "github.com/gongm/gohutool/crypto"
)

func main() {
    // AES 加密解密
    key := []byte("0123456789abcdef")
    encrypted, err := crypto.AESEncrypt(key, []byte("Hello, World!"))
    if err != nil {
        log.Fatal(err)
    }
    decrypted, err := crypto.AESDecrypt(key, encrypted)
    fmt.Println(string(decrypted)) // Hello, World!

    // HMAC-SHA256
    sig := crypto.HMACSHA256Hex([]byte("secret"), []byte("message"))
    fmt.Println(sig)

    // RSA 加密解密
    priv, pub, err := crypto.GenerateRSAKeyPair(2048)
    if err != nil {
        log.Fatal(err)
    }
    enc, _ := crypto.RSAEncrypt(pub, []byte("RSA message"))
    dec, _ := crypto.RSADecrypt(priv, enc)
    fmt.Println(string(dec)) // RSA message

    // ECDSA 签名
    ecdsaPriv, ecdsaPub, _ := crypto.GenerateECDSAKeyPair(elliptic.P256())
    sigBytes, _ := crypto.ECDSASign(ecdsaPriv, []byte("data"))
    fmt.Println(crypto.ECDSAVerify(ecdsaPub, []byte("data"), sigBytes)) // true
}
```
