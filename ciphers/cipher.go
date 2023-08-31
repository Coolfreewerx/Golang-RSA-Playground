package ciphers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"decode-decrypt-playground/unzips"
	"fmt"
	"io/ioutil"
)


func DecodeDecryptFile(result string, privateKeyData string, phassphrase string, encrypt_file_path string, decrypt_file_path string) {
	privateKey := BytesToPrivateKey([]byte(privateKeyData), []byte(phassphrase))
	
	// decode result by use base64
	Decode(result)

	// write encrypted file
	err := ioutil.WriteFile(encrypt_file_path, []byte(result), 0644)

	label := []byte("")
	hash := sha256.New()

	// decrypt file with Private Key
	fileEnc, _ := ioutil.ReadFile(encrypt_file_path)
 	fileDecrypt, err := DecryptOAEP(hash, rand.Reader, privateKey, fileEnc, label)
	if err != nil {
		fmt.Printf("Error Decrypt: %s", err)
		return
	}

	// write decrypted file
	err = ioutil.WriteFile(decrypt_file_path, fileDecrypt, 0644)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Printf("Decrypted success\n")

	// Unzip file
	err = unzips.UnzipSource(decrypt_file_path, "csv-files")
	if err != nil {
		fmt.Printf("Error unzipfile : %s", err)
		return
	}
}

// Work with csv file
func CipherFileWithRSAandBase64(public_key_path string, private_key_path string, file_zip_path string, pass string, encrypt_file_path string, decrypt_file_path string) {
	fileRaw, _ := ioutil.ReadFile(file_zip_path)

	publicKeyData, _ := ioutil.ReadFile(public_key_path)
	privateKeyData, _ := ioutil.ReadFile(private_key_path)
	phassphrase := []byte(pass)

	publicKey := BytesToPublicKey(publicKeyData)
	privateKey := BytesToPrivateKey(privateKeyData, phassphrase)

	fmt.Printf("%s\n", ExportPrivateKeyAsPemStr(privateKey))
	fmt.Printf("%s\n", ExportPublicKeyAsPemStr(publicKey))
	
	label := []byte("")
	hash := sha256.New()

	// Encrypt file with Tam's Public Key
	ciphertext, err := EncryptOAEP(hash, rand.Reader, publicKey, fileRaw, label)
	if err != nil {
		fmt.Printf("Error Encrypt: %s", err)
		return
	}

	// Write encrypted file
	err = ioutil.WriteFile("encrypt-file/covid-thailand-2021.pgp", ciphertext, 0644)
	 if err != nil {
		 fmt.Printf("Error %s", err)
		 return
	}
 	fmt.Printf("%s\n",ExportMsgAsPemStr(ciphertext))

	// Encode & Decode zip file from RSA Encryption by use base64
	data := ExportMsgAsPemStr(ciphertext)
	EncodeDecode(data)

 	// Decrypt file with Tam's Private Key
	fileEnc, _ := ioutil.ReadFile(encrypt_file_path)
 	fileDecrypt, err := DecryptOAEP(hash, rand.Reader, privateKey, fileEnc, label)
	if err != nil {
		fmt.Printf("Error Decrypt: %s", err)
		return
	}

	err = ioutil.WriteFile(decrypt_file_path, fileDecrypt, 0644)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Printf("Decrypted success\n")

	// Unzip file
	err = unzips.UnzipSource(decrypt_file_path, "csv-files")
	if err != nil {
		fmt.Printf("Error unzipfile : %s", err)
		return
	}
}

// Work with message string
func CipherMessageWithRSAandBase64(public_key_path string, private_key_path string, msg string, pass string) {

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
	EncodeDecode(data)

	// Decrypt message with Tam's Private Key
	plainText, _ := rsa.DecryptOAEP(hash, rand.Reader, privateKey, ciphertext, label)
	fmt.Printf("RSA decrypted : %s", plainText)
}