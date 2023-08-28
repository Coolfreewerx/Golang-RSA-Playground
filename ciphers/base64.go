package ciphers

import (
    b64 "encoding/base64"
    "fmt"
)

// Decode String from Base64
func DecodeString(data string) {
    fmt.Println("\n-------------Begin With Base64------------------")
    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println("\nEncode message from RSA By Base64 : ")
    fmt.Println(sEnc)

    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println("\nDecode message from RSA By Base64 : ")
    fmt.Println(string(sDec))

    // uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    // fmt.Println(uEnc)
    // uDec, _ := b64.URLEncoding.DecodeString(uEnc)
    // fmt.Println(string(uDec))

    fmt.Println("-------------End With Base64------------------")
    fmt.Println()
}