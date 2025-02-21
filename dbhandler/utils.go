package dbhandler

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"github.com/joho/godotenv"
	"io"
	"log"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetCreateQuery() string {
	return `
	CREATE TABLE IF NOT EXISTS secrets (
    	id SERIAL PRIMARY KEY,
    	account TEXT NOT NULL,
    	provider TEXT NOT NULL UNIQUE,
    	secret TEXT NOT NULL
	);
	`
}

func GetInsertQuery() string {
	return `INSERT INTO secrets (account, provider, secret) VALUES ($1, $2, $3);`
}

func GetSelectQuery() string {
	return `SELECT secret FROM secrets WHERE provider = $1;`
}

func Encrypt(key, text string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(text))
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(key, encryptedText string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil
}
