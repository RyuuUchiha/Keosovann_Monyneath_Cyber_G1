package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/sha3"
	"os"
	"strings"
)

func lab0() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("===========================================")
	fmt.Println("🔐💻  Lab 0: Proof the Hash Program  💻🔐")
	fmt.Println("===========================================")

	fmt.Print("🧩 Please input value 1: ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)

	fmt.Print("🧩 Please input value 2: ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)

	fmt.Println("\n⚙️  Processing hashes...")
	proofMe(input1, input2)
}

// hashText computes the hash of text based on the algorithm name
func hashText(text string, algo string) string {
	var hash []byte

	switch algo {
	case "MD5":
		sum := md5.Sum([]byte(text))
		hash = sum[:]
	case "SHA1":
		sum := sha1.Sum([]byte(text))
		hash = sum[:]
	case "SHA256":
		sum := sha256.Sum256([]byte(text))
		hash = sum[:]
	case "SHA512":
		sum := sha512.Sum512([]byte(text))
		hash = sum[:]
	case "SHA3-256":
		sum := sha3.Sum256([]byte(text))
		hash = sum[:]
	default:
		return "❌ Unknown algorithm"
	}

	return hex.EncodeToString(hash)
}

func proofMe(txt1, txt2 string) {
	algorithms := []string{"MD5", "SHA1", "SHA256", "SHA512", "SHA3-256"}

	fmt.Println("=============================")
	fmt.Println("🔒  Hash Comparison Results")
	fmt.Println("=============================")

	for _, algo := range algorithms {
		hashA := hashText(txt1, algo)
		hashB := hashText(txt2, algo)

		fmt.Printf("\n🧮 Algorithm: %s\n", algo)
		fmt.Printf(" 🔹 Hash A: %s\n", hashA)
		fmt.Printf(" 🔹 Hash B: %s\n", hashB)

		if hashA == hashB {
			fmt.Println(" ✅ Match!")
		} else {
			fmt.Println(" ❌ No Match!")
		}
	}
}
