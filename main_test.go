package main

import "testing"

func TestEncryptDecrypt(t *testing.T) {

	encrytped, _ := EncryptString("1234567890123456", "HelloWorld")
	decrytped, _ := DecryptString("1234567890123456", encrytped)

	if decrytped != "HelloWorld" {
		t.Errorf("decrytped != \"HelloWorld\": decrytped is:%s", decrytped)
	}

	_, err := EncryptString("123456789012345", "HelloWorld")

	if err == nil {
		t.Errorf("err == nil: err is:%s", "nil")
	}
}
