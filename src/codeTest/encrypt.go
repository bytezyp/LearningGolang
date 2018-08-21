package main

import (
	"encoding/base64"
	"strconv"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main()  {
	cipherHexText :=[]byte("MVBIYnlaZDduYWthdG5YVzM4QzV2c055clVjY3NDeVVaUglVBFdWVQBmOTFkODNm Nw%3D%3D")
	key := []byte("test")
	fmt.Println(bilibiliDecrypt(cipherHexText,key, nil, nil))

}
func bilibiliDecrypt(cipherHexText, ekey, ikey []byte, args ...interface{}) (float64, error)  {
	encodeString,_ := base64.StdEncoding.DecodeString(string(cipherHexText))
	vector := string(encodeString[1:33])
	encryted_price := string(encodeString[33:41])
	m := getMd5Hash(vector+string(ekey))
	price := string(ocx([]rune(encryted_price), []rune(m[0:9])))
	return strconv.ParseFloat(price, 64)
}
func getMd5Hash(text string)string  {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
func ocx(str1, str2 []rune) []rune  {
	result := make([]rune, len(str1))
	for index, _ :=range str1{
		result[index]=str1[index]^str2[index]
	}
	return result
}
