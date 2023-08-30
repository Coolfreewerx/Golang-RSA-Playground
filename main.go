package main

import (
	"crypto/rand"
	// "crypto/rsa"
	"crypto/sha256"
	"decode-decrypt-playground/ciphers"
	"decode-decrypt-playground/unzips"
	"fmt"
	"io/ioutil"
)

const pubKey = "ciphers/tam-public-key.asc"
const fileToEnc = "/tmp/data.txt"

func main() {

	/*/ 
	 	In practice, someone want to send a message to Tam will use Tam's public key to encrypt the message.
		Then, Bring encrypted message to Tam by use Base64 to encode the message. 
		After that, Tam will decode the message by use Base64 and decrypt the message by use Tam's private key.
		Finaly, Tam will get the message.
	/*/

 	m := "Test-RSA Decryption with Tam's Private Key"
	fmt.Printf("\nMessage : %s\n\n", m)


	tamPublicKeyData, _ := ioutil.ReadFile("ciphers/tamtam-public-key.asc")
	tamPrivateKeyData, _ := ioutil.ReadFile("ciphers/tamtam-private-key.asc")
	phassphrase := []byte("tamtam1122")

	tamPublicKey := ciphers.BytesToPublicKey(tamPublicKeyData)
	tamPrivateKey := ciphers.BytesToPrivateKey(tamPrivateKeyData, phassphrase)

 	fmt.Printf("%s\n",  ciphers.ExportPrivateKeyAsPemStr(tamPrivateKey))
 	fmt.Printf("%s\n", ciphers.ExportPublicKeyAsPemStr(tamPublicKey))

 	// message := []byte(m)
 	label := []byte("")
 	hash := sha256.New()

	fileRaw, _ := ioutil.ReadFile("ciphers/covid-thailand-2021.csv.zip")
	// Encrypt message with Tam's Public Key
 	ciphertext, err := ciphers.EncryptOAEP(hash, rand.Reader, tamPublicKey, fileRaw, label)
	if err != nil {
		fmt.Printf("Error Encrypt: %s", err)
		return
	}
	err = ioutil.WriteFile("encrypt-file/covid-thailand-2021.pgp", ciphertext, 0644)
	 if err != nil {
		 fmt.Printf("Error %s", err)
		 return
	 }
 	fmt.Printf("%s\n",ciphers.ExportMsgAsPemStr(ciphertext))

	// Encode & Decode message from RSA Encryption by use base64
	data := ciphers.ExportMsgAsPemStr(ciphertext)
	ciphers.DecodeString(data)

 	// Decrypt message with Tam's Private Key
	fileEnc, _ := ioutil.ReadFile("encrypt-file/covid-thailand-2021.pgp")
 	fileDecrypt, err := ciphers.DecryptOAEP(hash, rand.Reader, tamPrivateKey, fileEnc, label)
	if err != nil {
		fmt.Printf("Error Decrypt: %s", err)
		return
	}
	err = ioutil.WriteFile("decrypt-file/covid-thailand-2021.zip", fileDecrypt, 0644)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Printf("Decrypted success\n")
 	// fmt.Printf("RSA decrypted : %s", plainText)

	// Unzip file
	err = unzips.UnzipSource("decrypt-file/covid-thailand-2021.zip", "csv-files")
	if err != nil {
		fmt.Printf("Error unzipfile : %s", err)
		return
	}
}