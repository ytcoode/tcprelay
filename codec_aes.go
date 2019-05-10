package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"errors"
	"io"
)

var (
	aead cipher.AEAD
)

func initAes(key string) error {
	k := md5.Sum([]byte(key))
	b, err := aes.NewCipher(k[:])
	if err != nil {
		return err
	}
	gcm, err := cipher.NewGCM(b)
	if err != nil {
		return err
	}
	aead = gcm
	return nil
}

func encrypt(bs []byte) ([]byte, error) {
	if aead == nil {
		return bs, nil
	}
	nonce := make([]byte, aead.NonceSize())
	_, err := io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}
	return aead.Seal(nonce, nonce, bs, nil), nil
}

func decrypt(bs []byte) ([]byte, error) {
	if aead == nil {
		return bs, nil
	}
	if len(bs) < aead.NonceSize() {
		return nil, errors.New("decrypt: illegal bytes")
	}
	nonce := bs[:aead.NonceSize()]
	bs = bs[aead.NonceSize():]
	return aead.Open(nil, nonce, bs, nil)
}
