package main

import (
	"flag"
	"fmt"
	gosypt "github.com/alt-golang/gosypt.pkg"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	rc := m.Run()

	// rc 0 means we've passed,
	// and CoverMode will be non empty if run with -cover
	if rc == 0 && testing.CoverMode() != "" {
		c := testing.Coverage()
		if c < 0.874 {
			fmt.Println("Tests passed but coverage failed at", c)
			rc = -1
		}
	}
	os.Exit(rc)
}

func TestEncryptDecrypt(t *testing.T) {

	encrytped, _ := gosypt.EncryptString("1234567890123456", "HelloWorld")
	decrytped, _ := gosypt.DecryptString("1234567890123456", encrytped)

	if decrytped != "HelloWorld" {
		t.Errorf("decrytped != \"HelloWorld\": decrytped is:%s", decrytped)
	}

	_, err := gosypt.EncryptString("123456789012345", "HelloWorld")

	if err == nil {
		t.Errorf("err == nil: err is:%s", "nil")
	}
}

func TestMainFunc(t *testing.T) {

	fmt.Println(os.Args)
	origArgs := os.Args
	defer func() { os.Args = origArgs }()

	args := make([]string, 0)
	args = append(args, "gosypt")

	flag.CommandLine = flag.NewFlagSet("0", flag.ContinueOnError)
	os.Args = append(args, "-v")
	main()

	flag.CommandLine = flag.NewFlagSet("1", flag.ContinueOnError)
	os.Args = append(args, "-h")
	main()

	flag.CommandLine = flag.NewFlagSet("2", flag.ContinueOnError)
	os.Args = append(args, "-p")
	main()

	flag.CommandLine = flag.NewFlagSet("3", flag.ContinueOnError)
	os.Args = append(args, "-p", "1234567890123456")
	main()

	flag.CommandLine = flag.NewFlagSet("4", flag.ContinueOnError)
	os.Args = append(args, "-p", "1234567890123456", "-e", "HelloWorld")
	main()

	flag.CommandLine = flag.NewFlagSet("5", flag.ContinueOnError)
	os.Args = append(args, "-p", "1234567890123456", "-d", "G4hvsS/rhncB4cO58UEhz9crxQ9vp93k83l8mlBkz+c=")
	main()

	flag.CommandLine = flag.NewFlagSet("6", flag.ContinueOnError)
	os.Args = append(args, "-p", "1234567890123456", "-d")
	main()

	flag.CommandLine = flag.NewFlagSet("7", flag.ContinueOnError)
	os.Args = append(args, "-p", "1234567890123450", "-e", "HelloWorld")
	main()

}
