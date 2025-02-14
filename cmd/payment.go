package cmd

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"fmt"

	"github.com/razorpay/razorpay-go"
)

type RzpConfig struct {
	Key    string
	Secret string
	Client *razorpay.Client
}

func NewRzpConfig(key, secret string) (*RzpConfig, error) {
	if key == "" || secret == "" {
		return nil, fmt.Errorf("Key or Secret is unavailable.")
	}

	return &RzpConfig{
		Key:    key,
		Secret: secret,
		Client: razorpay.NewClient(key, secret),
	}, nil
}

func (rzp *RzpConfig) ObtainKey() string {
	return rzp.Key
}

func (rzp *RzpConfig) ExecuteRazorpay(amount int, eventId, name, contact,
	email, desc, txnId string) (string, error) {

	options := map[string]interface{}{
		"amount":          amount * 100,
		"currency":        "INR",
		"receipt":         txnId,
		"payment_capture": true, // automatically capture payment
		"notes": map[string]interface{}{
			"event_id": eventId,
		},
		"customer": map[string]interface{}{
			"name":  name,
			"email": email,
		},
		"description": desc,
	}

	order, err := rzp.Client.Order.Create(options, nil)
	if err != nil {
		return "", fmt.Errorf("Payment not initiated")
	}
	razorId, _ := order["id"].(string)
	return razorId, nil
}

func RazorPaymentVerification(sign, orderId, paymentId string) error {
	signature := sign
	secret := EnvVars.RzpSecret
	data := orderId + "|" + paymentId

	h := hmac.New(sha256.New, []byte(secret))
	_, err := h.Write([]byte(data))
	if err != nil {
		// This should never happen
		panic(err)
	}

	sha := hex.EncodeToString(h.Sum(nil))
	if subtle.ConstantTimeCompare([]byte(sha), []byte(signature)) != 1 {
		return fmt.Errorf("Payment failed")
	}
	return nil
}
