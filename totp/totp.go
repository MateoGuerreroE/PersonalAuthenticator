package totp

import (
	"fmt"
	"log"
	"strings"
	"time"

	"Personal2FA/dbhandler"
	"github.com/pquerna/otp/totp"
)

func GetSecret(appName string) (string, error) {
	db := dbhandler.InitDB()
	secret := dbhandler.GetSecret(db, appName)

	return strings.TrimSpace(secret), nil
}

func RegisterApp(appName, accountName string) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      appName,
		AccountName: accountName,
	})
	if err != nil {
		log.Fatal("Error generating key:", err)
	}

	// Save secret to database
	db := dbhandler.InitDB()
	dbhandler.StoreSecret(db, accountName, appName, key.Secret())

	fmt.Printf("App %s registered.\n", appName)
}

func GenerateTOTP(appName string) (string, error) {
	secret, err := GetSecret(appName)
	if err != nil {
		return "", fmt.Errorf("error getting secret for app %s: %w", appName, err)
	}

	token, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		return "", fmt.Errorf("error generating totp code: %w", err)
	}

	return token, nil
}
