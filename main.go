package main

import (
	"flag"
	"fmt"
	gosypt "github.com/alt-golang/gosypt.pkg"
	"log"
)

var version = "v1.0.1"

func main() {

	versionPtr := flag.Bool("v", false, "output the version number")
	passwordPtr := flag.String("p", "", "the secret key")
	encryptPtr := flag.String("e", "", "text to be encrypted")
	decryptPtr := flag.String("d", "", "text to be decrypted")
	helpPtr := flag.Bool("h", false, "output usage")
	flag.Parse()

	if *helpPtr == false {
		if *versionPtr {
			fmt.Printf("github.com/alt-lang/gosypt2 %s\n", version)
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
				result, err := gosypt.Decrypt(key, value)
				if err == nil {
					fmt.Printf("%s\n", string(result))
				} else {
					log.Fatalf("Error: %s", err)
				}

			} else if *encryptPtr != "" {
				key := []byte(*passwordPtr)
				value := []byte(*encryptPtr)
				result, err := gosypt.Encrypt(key, value)
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

}
