package ciphers

import (
    b64 "encoding/base64"
    "fmt"
)

// Decode String from Base64
func EncodeDecode(data string) {
    fmt.Println("\n-------------Begin With Base64------------------")
    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println("\nEncode message from RSA By Base64 : ")
    fmt.Println(sEnc)

    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println("\nDecode message from RSA By Base64 : ")
    fmt.Println(string(sDec))

    fmt.Println("-------------End With Base64------------------")
    fmt.Println()
}

func Decode(data string) {
    sDec, _ := b64.StdEncoding.DecodeString(data)
    fmt.Println(string(sDec))
}