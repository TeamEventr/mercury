package helpers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"os"
)

func GenerateRSAKeyPair(bits int) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return fmt.Errorf("Error generating private key: %w", err)
	}
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	publicKey := &privateKey.PublicKey
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	})

	if err := os.WriteFile("keys/app.rsa", privateKeyPEM, 0600); err != nil {
		return fmt.Errorf("Error saving private key: %w", err)
	}
	if err := os.WriteFile("keys/app.pub.rsa", publicKeyPEM, 0644); err != nil {
		return fmt.Errorf("error saving public key: %w", err)
	}
	fmt.Println("RSA keypair generated and saved successfully")
	return nil
}

func LoadPrivateKeyHex() (string, error) {
	data, err := os.ReadFile("keys/app.rsa")
	if err != nil {
		return "", fmt.Errorf("Failed to read private key file: %w", err)
	}
	block, _ := pem.Decode(data)
	if block == nil {
		return "", fmt.Errorf("Invalid private key PEM data.")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("Failed to parse private key: %w", err)
	}

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyHex := hex.EncodeToString(privateKeyBytes)

	return privateKeyHex, nil
}

func LoadPublicKeyHex() (string, error) {
	data, err := os.ReadFile("keys/app.pub.rsa")
	if err != nil {
		return "", fmt.Errorf("Failed to read public key file: %w", err)
	}
	block, _ := pem.Decode(data)
	if block == nil || block.Type != "PUBLIC KEY" {
		return "", fmt.Errorf("Invalid public key PEM data.")
	}
	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("Failed to parse public key: %w", err)
	}

	publicKeyBytes := x509.MarshalPKCS1PublicKey(publicKey)
	publicKeyHex := hex.EncodeToString(publicKeyBytes)

	return publicKeyHex, nil
}
