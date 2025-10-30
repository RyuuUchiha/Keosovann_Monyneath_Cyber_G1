package main

import (
	"Labw3_01_11_25/utils/crack"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)


func lab3() {

	fmt.Println("==========================================================")
	fmt.Println("🔐 🧪  Lab 3: Hashing (SHA512) — Password Cracker  🧪 🔐")
	fmt.Println("==========================================================")

	wordlistPath := "nord_vpn.txt"
	targetHash := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"
	verboseFile := "results/verbose3.txt"

	if v := os.Getenv("WORDLIST"); v != "" {
		wordlistPath = v
	}
	if v := os.Getenv("TARGET_HASH"); v != "" {
		targetHash = v
	}
	if v := os.Getenv("VERBOSE_FILE"); v != "" {
		verboseFile = v
	}

	// Open verbose log file (overwrite each run)
	vf, err := os.Create(verboseFile)
	if err != nil {
		log.Fatalf("❌ failed to create verbose file: %v", err)
	}
	defer vf.Close()

	// Write header
	fmt.Fprintf(vf, "Verbose run started: %s\nWordlist: %s\nTarget SHA512: %s\n\n", time.Now().Format(time.RFC3339), wordlistPath, targetHash)
	fmt.Printf("📁 Starting crack: wordlist=%s target=%s\n", wordlistPath, targetHash)

	// Open wordlist
	f, err := os.Open(wordlistPath)
	if err != nil {
		log.Fatalf("❌ failed to open wordlist '%s': %v", wordlistPath, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// increase buffer
	const maxCapacity = 10 * 1024 * 1024
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, maxCapacity)

	var tried int
	start := time.Now()
	for scanner.Scan() {
		tried++
		password := scanner.Text()
		password = strings.TrimSpace(password)
		if password == "" {
			continue
		}

		h := crack.SHA512Hash(password)

		// Verbose line: timestamp, attempt number, password, sha512
		line := fmt.Sprintf("%s | #%d | %s | %s\n", time.Now().Format(time.RFC3339), tried, password, h)
		if _, err := vf.WriteString(line); err != nil {
			log.Printf("⚠️ warning: failed to write verbose line: %v", err)
		}

		// Periodic progress to console every 10000 tries
		if tried%10000 == 0 {
			elapsed := time.Since(start).Truncate(time.Second)
			fmt.Printf("🔎 Tried: %d passwords (elapsed: %s)\n", tried, elapsed)
		}

		// Compare
		if strings.EqualFold(h, targetHash) {
			foundMsg := fmt.Sprintf("\n✅ FOUND! Password is: %s\nAttempts: %d\nTime: %.3fs\n", password, tried, time.Since(start).Seconds())
			fmt.Print(foundMsg)
			if _, err := vf.WriteString(foundMsg); err != nil {
				log.Printf("⚠️ warning: failed to write found message: %v", err)
			}

			// ensure verbose file is flushed to disk
			if err := vf.Sync(); err != nil {
				log.Printf("⚠️ warning: failed to sync verbose file: %v", err)
			}
			
			fmt.Printf("📁 The output saved to the file: %s\n", verboseFile)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("❌ error scanning wordlist: %v", err)
	}

	fmt.Printf("\n❌ Finished. Tried %d passwords. No match found.\n", tried)
	fmt.Fprintf(vf, "\n❌ Finished at %s. Tried %d passwords. No match found.\n", time.Now().Format(time.RFC3339), tried)

}
