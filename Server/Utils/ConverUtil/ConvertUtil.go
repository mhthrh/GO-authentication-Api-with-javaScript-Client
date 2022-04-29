package ConverUtil

import (
	"bytes"
	"encoding/base64"
	"math/rand"
)

func Byte64(msg []byte) string {
	return base64.StdEncoding.EncodeToString(msg)
}

func BytesToString(b []byte) string {
	return bytes.NewBuffer(b).String()
}

func RandomString(n int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
