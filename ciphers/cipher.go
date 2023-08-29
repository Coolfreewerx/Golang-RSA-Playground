package ciphers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
)

func CipherWithRSAandBase64(public_key_path string, private_key_path string, msg string, pass string) {
	fmt.Printf("\nMessage : %s\n\n", msg)

	publicKeyData, _ := ioutil.ReadFile(public_key_path)
	privateKeyData, _ := ioutil.ReadFile(private_key_path)
	phassphrase := []byte(pass)

	publicKey := BytesToPublicKey(publicKeyData)
	privateKey := BytesToPrivateKey(privateKeyData, phassphrase)

	fmt.Printf("%s\n", ExportPrivateKeyAsPemStr(privateKey))
	fmt.Printf("%s\n", ExportPublicKeyAsPemStr(publicKey))

	message := []byte(msg)
	label := []byte("")
	hash := sha256.New()

	// Encrypt message with Tam's Public Key
	ciphertext, _ := rsa.EncryptOAEP(hash, rand.Reader, publicKey, message, label)
	fmt.Printf("%s\n", ExportMsgAsPemStr(ciphertext))

	// Encode & Decode message from RSA Encryption by use base64
	data := ExportMsgAsPemStr(ciphertext)
	DecodeString(data)

	// Decrypt message with Tam's Private Key
	plainText, _ := rsa.DecryptOAEP(hash, rand.Reader, privateKey, ciphertext, label)
	fmt.Printf("RSA decrypted : %s", plainText)
}