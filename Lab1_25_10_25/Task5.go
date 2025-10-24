package main

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func task5() {

	fmt.Println("	📋Task5: Binary, Hex, and Base64 Encoding")

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n') 
	text = strings.TrimSpace(text)         

	binary := toBinary(text)

	hexValue := hex.EncodeToString([]byte(text))

	base64Value := base64.StdEncoding.EncodeToString([]byte(text))

	fmt.Println("\n--- Encoded Representations ---")
	fmt.Printf("Original Text: %s\n", text)
	fmt.Printf("Binary: %s\n", binary)
	fmt.Printf("Hexadecimal: %s\n", hexValue)
	fmt.Printf("Base64: %s\n", base64Value)
}


func toBinary(s string) string {
	binary := ""
	for _, c := range s {
		binary += fmt.Sprintf("%08b ", c)
	}
	return binary
}
