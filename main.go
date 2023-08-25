package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"decode-decrypt-playground/ciphers"
	"fmt"
)

func main() {

	/*/ 
	 	In practice, someone want to send a message to Tam will use Tam's public key to encrypt the message.
		and Tam will use his private key to decrypt the message.
	/*/

 	bits := 1024
 	m := "Test-RSA Decryption with Tam's Private Key"


 	tamPrivateKey, _ := rsa.GenerateKey(rand.Reader,bits)
 
 	tamPublicKey := &tamPrivateKey.PublicKey

 	fmt.Printf("%s\n",  ciphers.ExportPrivateKeyAsPemStr(tamPrivateKey))

 	fmt.Printf("%s\n", ciphers.ExportPublicKeyAsPemStr(tamPublicKey))

 	message := []byte(m)
 	label := []byte("")
 	hash := sha256.New()

	// Encrypt message with Tam's Public Key
 	ciphertext, _ := rsa.EncryptOAEP(hash, rand.Reader, tamPublicKey, message,label)
 	fmt.Printf("%s\n",ciphers.ExportMsgAsPemStr(ciphertext))

 	// Decrypt message with Tam's Private Key
 	plainText, _:= rsa.DecryptOAEP(hash, rand.Reader, tamPrivateKey, ciphertext, label)

 	fmt.Printf("RSA decrypted to [%s]", plainText)

}