package ciphers

import (
	"bytes"
	"crypto/rsa"

	"crypto/x509"
	"encoding/pem"
	"fmt"
	"hash"
	"io"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/packet"
)
// Export Public Key as Pem String
func ExportPublicKeyAsPemStr(pubkey *rsa.PublicKey) string {
    pubkey_pem := string(pem.EncodeToMemory(&pem.Block{Type:  "RSA PUBLIC KEY",Bytes: x509.MarshalPKCS1PublicKey(pubkey)}))
    return pubkey_pem
}

// Export Private Key as Pem String
func ExportPrivateKeyAsPemStr(privatekey *rsa.PrivateKey) string {
    privatekey_pem := string(pem.EncodeToMemory(&pem.Block{Type:  "RSA PRIVATE KEY",Bytes: x509.MarshalPKCS1PrivateKey(privatekey)}))
    return privatekey_pem
}

// Export Message as Pem String
func ExportMsgAsPemStr(msg []byte) string {
    msg_pem := string(pem.EncodeToMemory(&pem.Block{Type:  "MESSAGE",Bytes: msg}))
    return msg_pem
}

// Convert Bytes to Private Key
func BytesToPrivateKey(priv []byte, passphrase []byte) *rsa.PrivateKey {
    entityList, err := openpgp.ReadArmoredKeyRing(bytes.NewReader(priv))
	if err != nil {
        fmt.Println("Error reading private key: ", err)
		return nil
	}

    privateKey := entityList[0].PrivateKey
	if privateKey == nil {
        fmt.Println("No private key found in armor block")
		return nil
	}

	err = privateKey.Decrypt(passphrase)
	if err != nil {
        fmt.Println("Error decrypting private key: ", err)
		return nil
	}

	if privateKey.PubKeyAlgo != packet.PubKeyAlgoRSA {
        fmt.Println("PGP key is not RSA private key")
		return nil
	}

	return privateKey.PrivateKey.(*rsa.PrivateKey)
}

// Convert Bytes to Public Key
func BytesToPublicKey(pub []byte) *rsa.PublicKey {
    entityList, err := openpgp.ReadArmoredKeyRing(bytes.NewReader(pub))
	if err != nil {
        fmt.Println("Error reading public key: ", err)
		return nil
	}

	publicKey := entityList[0].PrimaryKey
	if publicKey == nil {
        fmt.Println("Error reading public key: ", err)
		return nil
	}

	if publicKey.PubKeyAlgo != packet.PubKeyAlgoRSA {
		fmt.Println("PGP key is not RSA public key")
		return nil
	}

	return publicKey.PublicKey.(*rsa.PublicKey)
}

// Encrypt too long message
func EncryptOAEP(hash hash.Hash, random io.Reader, public *rsa.PublicKey, msg []byte, label []byte) ([]byte, error) {
	msgLen := len(msg)
	step := public.Size() - 2*hash.Size() - 2
	var encryptedBytes []byte

	for start := 0; start < msgLen; start += step {
		finish := start + step
		if finish > msgLen {
			finish = msgLen
		}

		encryptedBlockBytes, err := rsa.EncryptOAEP(hash, random, public, msg[start:finish], label)
		if err != nil {
			return nil, err
		}

		encryptedBytes = append(encryptedBytes, encryptedBlockBytes...)
	}

	return encryptedBytes, nil
}

// Decrypt too long message
func DecryptOAEP(hash hash.Hash, random io.Reader, private *rsa.PrivateKey, msg []byte, label []byte) ([]byte, error) {
    msgLen := len(msg)
    step := private.PublicKey.Size()
    var decryptedBytes []byte

    for start := 0; start < msgLen; start += step {
        finish := start + step
        if finish > msgLen {
            finish = msgLen
        }

        decryptedBlockBytes, err := rsa.DecryptOAEP(hash, random, private, msg[start:finish], label)
        if err != nil {
            return nil, err
        }

        decryptedBytes = append(decryptedBytes, decryptedBlockBytes...)
    }

    return decryptedBytes, nil
}