package dbhandler

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func InitDB() *sql.DB {
	if os.Getenv("VERCEL_ENV") == "" {
		LoadEnv()
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatalf("DB_URL is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	_, err = db.Exec(GetCreateQuery())
	if err != nil {
		log.Printf("Failed to create table: %v", err)
	}

	return db
}

func StoreSecret(db *sql.DB, account string, provider string, secret string) {
	encryptKey := os.Getenv("ENCRYPT_KEY")
	encryptedSecret, err := Encrypt(encryptKey, secret)
	if err != nil {
		log.Printf("Failed to encrypt secret: %v", err)
	}
	_, err = db.Exec(GetInsertQuery(), account, provider, encryptedSecret)
	if err != nil {
		log.Printf("Failed to store secret: %v", err)
	} else {
		fmt.Println("Stored secret successfully")
	}
}

func GetSecret(db *sql.DB, provider string) string {
	var secret string
	encryptKey := os.Getenv("ENCRYPT_KEY")
	err := db.QueryRow(GetSelectQuery(), provider).Scan(&secret)
	if err != nil {
		log.Printf("Failed to get secret: %v", err)
		return ""
	}
	decryptedSecret, err := Decrypt(encryptKey, secret)
	if err != nil {
		log.Printf("Failed to decrypt secret: %v", err)
	}
	return decryptedSecret
}
