package main

import (
	"fmt"
	"os"
	"bufio"
	"encoding/base64"
	
)


func main() {

  println("Lab6:")

  print("Enter a string to encrypt: ")

  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  plainText := scanner.Text()

  print("Enter a key to start the encryption: ")
  var key byte
  fmt.Scanln(&key)

  encrypt := xorEncryption(plainText, key)
  println("Encrypted:",encrypt)

  de_encrypt, err := base64.StdEncoding.DecodeString(encrypt)
  if err != nil {
    print("error:", err)
    return
  }

  decrypt:= xorEncryption(string(de_encrypt), key)
  de_decrypt, err := base64.StdEncoding.DecodeString(decrypt)
  if err != nil {
    fmt.Print("Decrypt error: ", err, "\n")
    return
  }

  fmt.Println("Decrypted", string(de_decrypt))
}

func xorEncryption(pText string, key byte) string {

  cipher := []byte(pText)
  print(key)
  for i := 0; i < len(pText); i++ {
    cipher[i] = pText[i] ^ key
  }

  return base64.StdEncoding.EncodeToString(cipher)
}
