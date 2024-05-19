package decryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func ExecuteDecrypt(dirPath string, key []byte) error {
	return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			decrypt(path, key)
		}

		return nil
	})
}

func decrypt(fileName string, key []byte) error {
	fmt.Println("Decrypting", fileName)

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

	nonceSize := gcm.NonceSize()
	if len(contents) < nonceSize {
		fmt.Println("Error nonce size", file.Name())
		return err
	}

	nonce, contents := contents[:nonceSize], contents[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, contents, nil)
	if err != nil {
		fmt.Println("Error decrypting file", file.Name())
		return err
	}

	err = os.WriteFile(file.Name(), plaintext, 0777)
	if err != nil {
		fmt.Println("Error writing to file", file.Name())
		return err
	}

	return nil
}
