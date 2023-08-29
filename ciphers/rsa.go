package ciphers

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/packet"
)

func ExportPublicKeyAsPemStr(pubkey *rsa.PublicKey) string {
    pubkey_pem := string(pem.EncodeToMemory(&pem.Block{Type:  "RSA PUBLIC KEY",Bytes: x509.MarshalPKCS1PublicKey(pubkey)}))
    return pubkey_pem
}

func ExportPrivateKeyAsPemStr(privatekey *rsa.PrivateKey) string {
    privatekey_pem := string(pem.EncodeToMemory(&pem.Block{Type:  "RSA PRIVATE KEY",Bytes: x509.MarshalPKCS1PrivateKey(privatekey)}))
    return privatekey_pem
}

func ExportMsgAsPemStr(msg []byte) string {
    msg_pem := string(pem.EncodeToMemory(&pem.Block{Type:  "MESSAGE",Bytes: msg}))
    return msg_pem
}

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