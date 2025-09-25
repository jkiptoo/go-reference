//Using built-in support for base64 encoding and decoding

package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

	print := fmt.Println

	//Create string to encode and decode
	dataToEncode := "cde456!?$*&()'-=@~"

	//Encoder requires a byte, so we convert the string then encode using standard encoder
	standardEncode := b64.StdEncoding.EncodeToString([]byte(dataToEncode))
	print(standardEncode)

	//If not sure of the input, check for error during decoding
	standardDecode, _ := b64.StdEncoding.DecodeString(standardEncode)
	print(string(standardDecode))
	print()

	//Encode or decode using URL compatible base64
	urlEncode := b64.URLEncoding.EncodeToString([]byte(dataToEncode))
	print(urlEncode)
	urlDecode, _ := b64.URLEncoding.DecodeString(urlEncode)
	print(string(urlDecode))



}