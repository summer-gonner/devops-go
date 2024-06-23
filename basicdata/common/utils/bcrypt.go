package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"log"

	"github.com/tjfoc/gmsm/sm2"
	x509g "github.com/tjfoc/gmsm/x509"
	"golang.org/x/crypto/bcrypt"
)

func GenerateFromPassword(pwd string) (hashedPassword string, err error) {
	password := []byte(pwd)
	// Hashing the password with the default cost of 10
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	hashedPassword = string(hashedPasswordBytes)
	return
}

func CompareHashAndPassword(hashedPwd, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		return false
	}
	return true
}

// EncryptSm2 加密
func EncryptSm2(privateKey, content string) string {
	// 从十六进制导入公私钥
	priv, err := x509g.ReadPrivateKeyFromHex(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 公钥加密部分
	msg := []byte(content)
	pub := &priv.PublicKey
	cipherTxt, err := sm2.Encrypt(pub, msg, rand.Reader, sm2.C1C2C3) // sm2加密
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("加密文字:%s\n加密结果:%x\n", msg, cipherTxt)
	encodeRes := fmt.Sprintf("%x", cipherTxt)
	return encodeRes
}

// DecryptSm2 解密
func DecryptSm2(privateKey, encryptData string) (string, error) {
	// 从十六进制导入公私钥
	priv, err := x509g.ReadPrivateKeyFromHex(privateKey)
	if err != nil {
		return "", err
	}
	// 私钥解密部分
	hexData, err := hex.DecodeString(encryptData)
	if err != nil {
		return "", err
	}
	plainTxt, err := sm2.Decrypt(priv, hexData, sm2.C1C2C3) // sm2解密
	if err != nil {
		return "", err
	}
	// fmt.Printf("解密后的明文：%s\n私钥：%s \n 匹配一致", plainTxt, x509.WritePrivateKeyToHex(priv))
	return string(plainTxt), nil
}

// EncryptAndDecrypt 加密/解密
func EncryptAndDecrypt(privateKey, content string) {
	// 从十六进制导入公私钥
	priv, err := x509g.ReadPrivateKeyFromHex(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 公钥加密部分
	msg := []byte(content)
	pub := &priv.PublicKey
	cipherTxt, err := sm2.Encrypt(pub, msg, rand.Reader, sm2.C1C2C3) // sm2加密
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("加密文字:%s\n加密结果:%x\n", msg, cipherTxt)

	// 私钥解密部分
	plainTxt, err := sm2.Decrypt(priv, cipherTxt, sm2.C1C2C3) // sm2解密
	if err != nil {
		log.Fatal(err)
	}
	if !bytes.Equal(msg, plainTxt) {
		log.Fatal("原文不匹配:", msg)
	}
	fmt.Printf("解密后的明文：%s\n私钥：%s \n 匹配一致", plainTxt, x509g.WritePrivateKeyToHex(priv))
}

// EncryptRSA 加密
func EncryptRSA(content, publicKey string) (encryptStr string, err error) {
	// 	var publicKey = `-----BEGIN PUBLIC KEY-----
	// MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDaIWAL13RU+bJN2hfmTSyOBotf
	// 71pq8jc2ploPBHtN3smTUkYPbX2MIbO9TrRj3u67s/kGQZrz6tyQ68oexpukPN4/
	// ypzp64UA5CQENSA41ZxTpYADbFQsiX9Spv6aDHhHzUlZtWRru9ptcFO3tDKq0ACT
	// OAR1ZEHFwQGhzwaAowIDAQAB
	// -----END PUBLIC KEY-----`
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return "", fmt.Errorf("failed to parse public key PEM")
	}
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	//  类型断言
	rsaPublicKey := publicKeyInterface.(*rsa.PublicKey)
	// 加密数据
	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, []byte(content))
	if err != nil {
		return "", fmt.Errorf("error encrypting data:%v", err)
	}

	return base64.StdEncoding.EncodeToString(encryptedData), err

}

// DecryptRSA 解密
func DecryptRSA(encryptStr, privateKey string) (content string, err error) {
	// 	var privateKey = `-----BEGIN PRIVATE KEY-----
	// MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBANohYAvXdFT5sk3a
	// F+ZNLI4Gi1/vWmryNzamWg8Ee03eyZNSRg9tfYwhs71OtGPe7ruz+QZBmvPq3JDr
	// yh7Gm6Q83j/KnOnrhQDkJAQ1IDjVnFOlgANsVCyJf1Km/poMeEfNSVm1ZGu72m1w
	// U7e0MqrQAJM4BHVkQcXBAaHPBoCjAgMBAAECgYA/aJJN/uyvQwKlBPALn4WDJ73e
	// PmrvScfpGAR39xqM8WVxcOoy0+Y6FRX1wupHWefWIqQSQIH1w+EoM5LGzX8yflSo
	// lG3E0mgJzrMAOTs5FVkdN4tV6rKYq/vA9R67AD0a9nq7yOFeTqjGzWj4l7Vptvu4
	// prK5GWV+i0+mpB2kKQJBAP0n1EMAHQSW38zOngfaqC6cvnjEbX4NnhSPRZVzlu3y
	// ZkitiA/Y96yCCybCWD0TkF43Z1p0wIGuXSJ1Igku6bcCQQDclMziUz1RnQDl7RIN
	// 449vbmG2mGLoXp5HTD9QP0NB46w64WwXIX7IZL2GubndTRFUFTTPLZZ80XbhFtp6
	// 19B1AkEAnIgjJGaOisbrjQz5BCw8r821rKDwfu/WninUwcteOLUYb7n1Fq92vZEP
	// aiDjRKizLL6fRnxIiCcTaXn52KnMUwJBAJaKOxYPRx8G7tD8rcCq2H5tL+TFNWNv
	// B8iTAfbLZiR2tFlu9S0IIBW1ox9qa63b5gKjgmoOq9C9x8swpKUH2u0CQAKDHqwh
	// aH6lVtV8cw55Ob8Dsh3PgFUazuM1+e5PjmZku3/2jeQQJrecu/S6LooPdeUf+EtV
	// OB/5HvFhGpEu2/E=
	// -----END PRIVATE KEY-----`
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return "", fmt.Errorf("failed to parse private key PEM")
	}
	privateKeyData, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	privateKeyInterface := privateKeyData.(*rsa.PrivateKey)

	// 解密数据

	byt, err := base64.StdEncoding.DecodeString(encryptStr)
	if err != nil {
		return "", fmt.Errorf("base64 DecodeString err:%v", err)
	}

	decryptedData, err := rsa.DecryptPKCS1v15(rand.Reader, privateKeyInterface, byt)
	if err != nil {
		return "", fmt.Errorf("error decrypting data:%v", err)
	}

	return string(decryptedData), nil

}

func Md5(s []byte) string {
	m := md5.New()
	m.Write(s)

	return hex.EncodeToString(m.Sum(nil))
}
