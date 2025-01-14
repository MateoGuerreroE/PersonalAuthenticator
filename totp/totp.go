package totp

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
)

func GetSecret(appName string) (string, error) {
	fileName := fmt.Sprintf("%s.secret", appName)

	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("error reading secret file %s: %w", fileName, err)
	}

	return strings.TrimSpace(string(data)), nil
}

func RegisterApp(appName, accountName string) string {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      appName,
		AccountName: accountName,
	})
	if err != nil {
		log.Fatal("Error generating key:", err)
	}

	// Save secret to a file or database
	file, _ := os.Create(appName + ".secret")
	defer file.Close()
	file.WriteString(key.Secret())

	// Generate QR code for app to scan
	qrFile := appName + "-qrcode.png"
	qrcode.WriteFile(key.URL(), qrcode.Medium, 256, qrFile)

	fmt.Printf("App %s registered. Scan the QR code: %s\n", appName, qrFile)
	return key.Secret()
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
