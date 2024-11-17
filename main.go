package main

import (
	"fmt"
	"os"
)

func xorStreamCipher(data, key []byte) []byte {
	ciphered := make([]byte, len(data))
	keyLen := len(key)

	for i := 0; i < len(data); i++ {
		ciphered[i] = data[i] ^ key[i%keyLen]
	}

	return ciphered
}

func printHelp() {
	fmt.Println("Usage: stream-cipher <input_file> <output_file> <key> <mode>")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  <input_file>    Tên file đầu vào chứa văn bản cần mã hóa hoặc giải mã")
	fmt.Println("  <output_file>   Tên file đầu ra để lưu trữ kết quả sau khi mã hóa hoặc giải mã")
	fmt.Println("  <key>           Khóa dùng để mã hóa và giải mã (nhập dưới dạng chuỗi)")
	fmt.Println("  <mode>          Chế độ hoạt động: 'encrypt' để mã hóa hoặc 'decrypt' để giải mã")
	fmt.Println()
	fmt.Println("Example:")
	fmt.Println("  stream-cipher plaintext.txt encrypted.txt mysecretkey encrypt")
	fmt.Println("  stream-cipher encrypted.txt decrypted.txt mysecretkey decrypt")
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--help" {
		printHelp()
		return
	}

	if len(os.Args) < 5 {
		fmt.Println("Missing parameters. Use 'stream-cipher --help' to see instructions.")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]
	key := []byte(os.Args[3])
	mode := os.Args[4]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	var result []byte
	if mode == "encrypt" {
		result = xorStreamCipher(data, key)
		fmt.Println("Encryption completed.")
	} else if mode == "decrypt" {
		result = xorStreamCipher(data, key)
		fmt.Println("Decryption completed.")
	} else {
		fmt.Println("Invalid mode. Use 'encrypt' or 'decrypt'.")
		return
	}

	err = os.WriteFile(outputFile, result, 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		return
	}

	fmt.Printf("Output written to %s\n", outputFile)
}