package main

import (
	"bytes"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	err := initAes("abc")
	if err != nil {
		t.Fail()
	}

	bs := []byte("hello world")

	b1, err := encrypt(bs)
	if err != nil {
		t.Fail()
	}

	b2, err := decrypt(b1)
	if err != nil {
		t.Fail()
	}

	if !bytes.Equal(b2, bs) {
		t.Fail()
	}
}
