package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println("==== XOR Encryption/Decryption ====")

	var text string
	var key string

	fmt.Print("Enter text: ")
	fmt.Scanln(&text)
	fmt.Print("Enter key (single character): ")
	fmt.Scanln(&key)

	// XOR Encrypt
	ciphertext := xorEncrypt(text, key[0])
	cipherHex := hex.EncodeToString([]byte(ciphertext))

	// XOR Decrypt
	decrypted := xorEncrypt(ciphertext, key[0])

	fmt.Println("\n--- Results ---")
	fmt.Println("Plaintext:", text)
	fmt.Println("Key:", string(key[0]))
	fmt.Println("Ciphertext (Hex):", cipherHex)
	fmt.Println("Decrypted Text:", decrypted)
}

// XOR encryption/decryption function
func xorEncrypt(text string, key byte) string {
	result := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		result[i] = text[i] ^ key
	}
	return string(result)
}
