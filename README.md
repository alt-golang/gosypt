# Gosypt

An idiomatic port of "Jasypt" in Go
=====================================


![Language Badge](https://img.shields.io/github/languages/top/alt-golang/gosypt) <br/>
[release notes](https://github.com/alt-golang/gosypt/blob/main/HISTORY.md)

<a name="intro">Introduction</a>
--------------------------------
An simple encryption and decryption module, using the idiomatic Go crypto/aes
packages.

<a name="usage">Usage</a>
-------------------------

To use the module in your Go modules `go get -d github.com/alt-golang/gosypt`

```go
    package main
	encrytped, _ := EncryptString("1234567890123456", "HelloWorld")
	fmt.Println(encrypted)
	decrytped, _ := DecryptString("1234567890123456", encrytped)
```

To use the module as a command from the command line `go install github.com/alt-golang/gosypt`
(and add the GOPATH/bin dir to your path)

```shell
    go install github.com/alt-golang/gosypt   
    export $PATH=$PATH:`go env GOPATH`\bin
    gosypt -v
    > github.com/alt-lang/gosypt v1.0.0
    gosypt -h
    > Usage of gosypt:
    >  -d string
    >        text to be decrypted
    >  -e string
    >        text to be encrypted
    >  -h	output usage
    >  -p string
    >        the secret key
    >  -v	output the version number
    gosypt -p 1234567890123456 -e HelloWorld
    > G4hvsS/rhncB4cO58UEhz9crxQ9vp93k83l8mlBkz+c=
    gosypt -p 1234567890123456 -d G4hvsS/rhncB4cO58UEhz9crxQ9vp93k83l8mlBkz+c=
    > HelloWorld
```

<a name="license">License</a>
-----------------------------

May be freely distributed under the [MIT license](https://raw.githubusercontent.com/alt-golang/gosypt/main/LICENSE).

Copyright (c) 2022 Craig Parravicini    
