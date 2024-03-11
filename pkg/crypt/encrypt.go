package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"log"
	"os"
)

// Function to encrypt file
func EncryptFile(inputFile string, key []byte) error {
	// Read the original file
	input, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	// Verify If file encrypted
	if isEncrypted(input) {
		log.Println("The file has already been encrypted")
		return nil
	}

	// Add head to the file
	encryptedHeader := []byte("This file has been encrypted\n")
	encryptedData := append(encryptedHeader, input...)

	// Create a new block used key
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Generate new vector
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return err
	}

	// Create a new cipher in CBC mood
	stream := cipher.NewCFBEncrypter(block, iv)
	// Encrypt content file
	stream.XORKeyStream(encryptedData[len(encryptedHeader):], encryptedData[len(encryptedHeader):])

	// Write content in the same file
	if err := os.WriteFile(inputFile, encryptedData, 0644); err != nil {
		return err
	}

	return nil
}

// Function to verify If file is encrypted
func isEncrypted(data []byte) bool {
	encryptedHeader := []byte("This file has been encrypted\n")
	return len(data) > len(encryptedHeader) && string(data[:len(encryptedHeader)]) == string(encryptedHeader)
}
