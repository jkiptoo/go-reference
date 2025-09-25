//Using SHA256 hashes to compute short identities for binary or text blobs.
//'crypto' packages are used to implement various hash functions

package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	
	print := fmt.Println
	//Start a new hash
	newHash := "Pass this string through sha256"

	hashedString := sha256.New()

	//Transform string into bytes using []byte because '.Write' expects bytes only
	hashedString.Write([]byte(newHash))
	
	//Get final hash result as a byte slice.
	//The argument passed to Sum can be appended to an existing byte slice
	hashResult := hashedString.Sum(nil)

	print(newHash)
	fmt.Printf("%x\n", hashResult)

	//for SHA 512, import crypto/sha512 and use sha512.New()
	//Pay careful attention to hash strength

}