package helpers

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"

	"github.com/o1egl/paseto/v2"
)

type Payload struct {
	Username  string
	Email     string
	TokenType string
	ExpiryAt  time.Time
}

func GenerateTokens(username string, email string) (string, string, error) {
	// Access token valid for 1 hour
	accessToken, err := CreateToken(username, email, "access", 60*time.Minute)
	if err != nil {
		return "", "", fmt.Errorf("Failed to generate access token.\n%w", err)
	}
	// Refresh token valid for 150 days
	refreshToken, err := CreateToken(username, email, "access", 150*24*time.Hour)
	if err != nil {
		return "", "", fmt.Errorf("Failed to generate refresh token.\n%w", err)
	}
	return accessToken, refreshToken, nil
}

func CreateToken(username string, email string,
	tokenType string, expiresIn time.Duration) (string, error) {
	payload := Payload{
		Username:  username,
		Email:     email,
		TokenType: tokenType,
		ExpiryAt:  time.Now().Add(expiresIn),
	}

	// Accessign private key
	privateKey, err := LoadPrivateKey("private_key.pem")
	if err != nil {
		return "", err
	}
	privPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	token, err := paseto.NewV2().Encrypt(privPEM, payload, nil)
	if err != nil {
		return "", fmt.Errorf("Failed to encode token: %w", err)
	}

	return token, nil
}

func VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	// Accessing public key
	publicKey, err := LoadPublicKey("public_key.pem")
	if err != nil {
		return payload, fmt.Errorf("Failed to load public key: %w", err)
	}
	pubPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	})

	err = paseto.NewV2().Decrypt(token, pubPEM, payload, nil)
	if err != nil {
		return payload, err
	}
	return payload, nil
}
