package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
	"strings"
)

const nmeaLat = "0521"
const nmeaLong = "02459"
const logName = "logs/20200628_153027.log"

func main() {
	decryptKnown()
	//bruteforce()
}

func decryptKnown() {
	ciphertext, err := ioutil.ReadFile(logName)
	if err != nil {
		panic(err)
	}
	plaintext := make([]byte, len(ciphertext), len(ciphertext))

	key := strings.Repeat(nmeaLat, 4)
	IV := strings.Repeat(nmeaLong, 3) + "0"
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	cbc := cipher.NewCBCDecrypter(block, []byte(IV))
	cbc.CryptBlocks(plaintext, ciphertext)

	err = ioutil.WriteFile(logName+".dec", plaintext, 0644)
	if err != nil {
		panic(err)
	}
}

func bruteforce() {
	ciphertext, err := ioutil.ReadFile(logName)
	if err != nil {
		panic(err)
	}
	plaintext := make([]byte, len(ciphertext), len(ciphertext))

	for i := 0; i < 9999; i++ {
		keyInput := fmt.Sprintf("%04d", i)
		key := strings.Repeat(keyInput, 4)
		IV := strings.Repeat("00000", 3) + "0"
		block, err := aes.NewCipher([]byte(key))
		if err != nil {
			panic(err)
		}

		cbc := cipher.NewCBCDecrypter(block, []byte(IV))
		cbc.CryptBlocks(plaintext, ciphertext)

		if string(plaintext[28:31]) == ",N," {
			fmt.Printf("Found key: %s\n", keyInput)
			return
		}
	}

}
