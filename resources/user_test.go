package resources

import (
	"encoding/base32"
	"testing"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func TestGenerateOTPCode(t *testing.T) {
	secret := base32.StdEncoding.EncodeToString([]byte("JBSWY3DPEHPK3PXP"))
	passcode, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
		Period:    30,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA512,
	})
	if err != nil {
		panic(err)
	}
	println(passcode)
}
