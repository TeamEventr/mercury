package helpers

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"time"

	"aidanwoods.dev/go-paseto"
)

const (
	RefreshTokenValidTime = time.Hour * 24 * 90
	AuthTokenValidTime    = time.Minute * 30
	privateKeyPath        = "keys/app.rsa"
	publicKeyPath         = "keys/app.rsa.pub"
)

type TokenClaims struct {
	Id        string
	Issuer    string
	IssuedAt  time.Time
	ExpiresAt time.Duration
	Role      string `json:"role"`
	Csrf      string `json:"csrf"`
}

var (
	VerifyKey paseto.V4AsymmetricPublicKey
	SignKey   paseto.V4AsymmetricSecretKey
)

func InitPaseto() error {
	privateKeyHex, err := LoadPrivateKeyHex()
	if err != nil {
		return err
	}
	publicKeyHex, err := LoadPublicKeyHex()
	if err != nil {
		return err
	}
	// Verify using public key
	VerifyKey, err = paseto.NewV4AsymmetricPublicKeyFromHex(publicKeyHex)
	if err != nil {
		return err
	}
	// Sign using private key
	SignKey, err = paseto.NewV4AsymmetricSecretKeyFromHex(privateKeyHex)
	if err != nil {
		return err
	}
	return nil
}

func GenerateCsrfSecret() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}

	csrfToken := base64.StdEncoding.EncodeToString(token)
	return csrfToken, nil
}

func VerifyCsrfSecret(cookieCsrf, headerCsrf string) bool {
	decodedCsrfFromCookie, err := base64.StdEncoding.DecodeString(cookieCsrf)
	if err != nil {
		return false
	}
	decodedCsrfFromHeader, err := base64.StdEncoding.DecodeString(headerCsrf)
	if err != nil {
		return false
	}

	if bytes.Equal(decodedCsrfFromCookie, decodedCsrfFromHeader) {
		return true
	}
	return false
}

// func CreateAuthToken(email, role, csrf string) (auth string, err error) {
//
// }
//
// func CreateRefreshToken(email, role, csrf string) (auth string, err error) {
//
// }
//
// func CreateTokens(email, role string) (auth, refresh, csrf string, err error) {
//
// }
//
// func UpdateAuthToken(oldAuth, refresh string) (newAuth, csrf string, err error) {
//
// }
//
// func UpdateRefreshToken(oldRefresh string) (newRefresh string, err error) {
//
// }
//
// func VerifyAndRefreshTokens(oldauth, oldRefresh, oldCsrf string) (newAuth,
// 	newRefresh, newCsrf string, err error) {
//
// }
//
// func RevokeRefreshToken(refresh string) error {
//
// }
