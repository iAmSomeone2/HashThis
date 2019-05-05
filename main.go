package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

// md5Hash computes and returns the MD5 hash of the file the filePath string
// specifies.
func md5Hash(filePath string) (string, error) {
	// Init the return string in case things break
	var md5Str string

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return md5Str, err
	}
	defer file.Close()

	hash := md5.New()

	// Copy contents of file into the hash interface
	if _, err := io.Copy(hash, file); err != nil {
		return md5Str, err
	}

	// Get the 16 byte hash
	byteHash := hash.Sum(nil)[:16]

	// Convert to string
	md5Str = hex.EncodeToString(byteHash)
	return md5Str, nil
}

// sha256Hash computes and returns the sha256 hash of the file the filePath string
// specifies.
func sha256Hash(filePath string) (string, error) {
	// Init the return string in case things break
	var sha256Str string

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return sha256Str, err
	}
	defer file.Close()

	hash := sha256.New()

	// Copy contents of file into the hash interface
	if _, err := io.Copy(hash, file); err != nil {
		return sha256Str, err
	}

	// Get the 16 byte hash
	byteHash := hash.Sum(nil)[:32]

	// Convert to string
	sha256Str = hex.EncodeToString(byteHash)
	return sha256Str, nil
}

func main() {
	useSHA256 := flag.Bool("sha256", false, "Compute SHA256 hash.")
	useMD5 := flag.Bool("md5", false, "Compute MD5 hash.")
	fileName := flag.String("file", "", "File to check.")
	flag.Parse()

	if *fileName == "" {
		log.Fatalf("No file name provided!\n")
	}

	if *useSHA256 {
		shaHash, err := sha256Hash(*fileName)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		fmt.Printf("SHA256 Hash: %s\n", shaHash)
	}

	if *useMD5 {
		md5Val, err := md5Hash(*fileName)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		fmt.Printf("MD5 Hash: %s\n", md5Val)
	}

	if !*useMD5 && !*useSHA256 {
		log.Fatalf("No hashing algorithm selected!\n")
	}
}
