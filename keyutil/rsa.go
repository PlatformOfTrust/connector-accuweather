package keyutil

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type rsaKeyType int

const (
	rsaKeyInvalid rsaKeyType = iota
	rsaKeyFile
	rsaKeyURL
)

func LoadRsaKeys(keys []string) ([]*rsa.PublicKey, error) {
	res := []*rsa.PublicKey{}

	for _, key := range keys {
		rsaKey, err := LoadRsaKey(key)
		if err != nil {
			return res, err
		}
		res = append(res, rsaKey)
	}

	return res, nil
}

func LoadRsaKey(key string) (*rsa.PublicKey, error) {
	keyType := ParseKeyType(key)
	switch keyType {
	case rsaKeyURL:
		return FetchRsaKey(key)
	case rsaKeyFile:
		return LoadRsaKeyFile(key)
	}

	return nil, errors.New("failed to load specified key")
}

// parseKeyType determines if the provided path for the key is a file or a url
func ParseKeyType(key string) rsaKeyType {
	if strings.HasPrefix(key, "http://") || strings.HasPrefix(key, "https://") {
		return rsaKeyURL
	}

	return rsaKeyFile
}

func LoadRsaKeyFile(path string) (*rsa.PublicKey, error) {
	key, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return ParsePEM(key)
}

func FetchRsaKey(url string) (*rsa.PublicKey, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return ParsePEM(body)
}

func ParsePEM(c []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(c)

	cert, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	rsaPublicKey, ok := cert.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("invalid public key")
	}

	return rsaPublicKey, nil
}
