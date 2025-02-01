package cmd

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"

	"github.com/razorpay/razorpay-go"
)

func ExecuteRazorpay() (string, error) {
	client := razorpay.NewClient("key", "secret")

	data := map[string]interface{}{
		// FIXME: Adding a dummy price tag for now
		"amount":   int(100) * 100,
		"currency": "INR",
		"receipt":  "101", // replace with receipt-id
	}

	body, err := client.Order.Create(data, nil)
	if err != nil {
		return "", errors.New("Payment not initiated")
	}
	razorId, _ := body["id"].(string)
	return razorId, nil
}

func RazorPaymentVerification(sign, orderId, paymentId string) error {
	signature := sign
	secret := "SECRET" // TODO: Add the actual secret
	data := orderId + "|" + paymentId

	h := hmac.New(sha256.New, []byte(secret))
	_, err := h.Write([]byte(data))
	if err != nil {
		// This should never happen
		panic(err)
	}

	sha := hex.EncodeToString(h.Sum(nil))
	if subtle.ConstantTimeCompare([]byte(sha), []byte(signature)) != 1 {
		return errors.New("Payment failed")
	}
	return nil
}
