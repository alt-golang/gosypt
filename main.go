package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	//argsWithProg := os.Args
	//argsWithoutProg := os.Args[1:]

	versionPtr := flag.Bool("v", false, "output the version number")
	passwordPtr := flag.String("p", "", "the secret key")
	encryptPtr := flag.String("e", "", "text to be encrypted")
	decryptPtr := flag.String("d", "", "text to be decrypted")
	helpPtr := flag.Bool("h", false, "output usage")
	flag.Parse()

	if *helpPtr == false {
		if *versionPtr {
			fmt.Println("alt-lang/gosypt v1.0.0")
		} else if *passwordPtr == "" {
			fmt.Println("Error: a secret key (flag -p) is required")
			flag.Usage()
		} else {
			if *encryptPtr == "" && *decryptPtr == "" {
				fmt.Println("Error: text to encrypt/decrypt (flag -e or -d) is required")
				flag.Usage()
			} else if *decryptPtr != "" {
				key := []byte(*passwordPtr)
				value := []byte(*decryptPtr)
				fmt.Printf("Decrypting value %s with key %s\n", value, key)
				result, err := Decrypt(key, value)
				if err == nil {
					fmt.Printf("%s\n", string(result))
				} else {
					log.Fatalf("Error: %s", err)
				}

			} else if *encryptPtr != "" {
				key := []byte(*passwordPtr)
				value := []byte(*encryptPtr)
				fmt.Printf("Encrypting value %s with key %s\n", value, key)
				result, err := Encrypt(key, value)
				if err == nil {
					fmt.Printf("%s\n", string(result))
				} else {
					log.Fatalf("Error: %s", err)
				}
			}
		}
	} else {
		flag.Usage()
	}

	//key := []byte("a very very very very secret key") // 32 bytes
	//plaintext := []byte("hello")
	//fmt.Printf("%s\n", plaintext)
	//ciphertext, err := Encrypt(key, plaintext)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%0x\n", ciphertext)
	//result, err := Decrypt(key, ciphertext)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%s\n", result)
}
