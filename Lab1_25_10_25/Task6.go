package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func task6() {

	fmt.Println("	📋Task6: XOR Encryption/Decryption")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	fmt.Print("Enter key (single byte): ")
	keyInput, _ := reader.ReadString('\n')
	keyInput = strings.TrimSpace(keyInput)

	var key byte
	if num, err := strconv.Atoi(keyInput); err == nil {
		key = byte(num) 
	} else {
		key = keyInput[0]
	}

	ciphertext := xorEncrypt(text, key)
	cipherHex := hex.EncodeToString([]byte(ciphertext))

	decrypted := xorEncrypt(ciphertext, key)

	fmt.Println("\n--- Results ---")
	fmt.Println("Plaintext:", text)
	fmt.Println("Key:", keyInput)
	fmt.Println("Ciphertext (Hex):", cipherHex)
	fmt.Println("Decrypted Text:", decrypted)
}


func xorEncrypt(text string, key byte) string {
	result := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		result[i] = text[i] ^ key
	}
	return string(result)
}
