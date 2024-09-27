package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

//func ReadTokenKey(filename string) ([]byte, error) {
//	// Read the PEM file
//	pemData, err := ioutil.ReadFile(filename)
//	if err != nil {
//		return []byte{}, err
//	}
//
//	return pemData, nil
//
//}

func LoadPrivateKeyFromFile(filename string) (*rsa.PrivateKey, error) {
	privateKeyData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(privateKeyData)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	// Use ParsePKCS8PrivateKey instead
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// Type assertion to ensure it's an RSA private key
	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("private key is not an RSA private key")
	}

	return rsaPrivateKey, nil
}

func LoadPublicKeyFromFile(filename string) (*rsa.PublicKey, error) {
	publicKeyData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(publicKeyData)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing public key") // Corrected error message
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return nil,
			fmt.Errorf("key is not an RSA public key")
	}

	return publicKey, nil
}

func ReadSecretKeys(publicKeyFileName, privateKeyFileName string) (*SecretKeys, error) {
	publicKey, err := LoadPublicKeyFromFile(publicKeyFileName)
	if err != nil {
		return nil, err
	}

	privateKey, err := LoadPrivateKeyFromFile(privateKeyFileName)
	if err != nil {
		return nil, err
	}

	return &SecretKeys{
		Public:  publicKey,
		Private: privateKey,
	}, nil

}
