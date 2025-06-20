package utils

import "encoding/base64"

func ArmorValue(value []byte) string {
	return base64.StdEncoding.EncodeToString(value)
}

func DearmorValue(value string) []byte {
	bt, _ := base64.StdEncoding.DecodeString(value)
	return bt
}
