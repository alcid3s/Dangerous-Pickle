/*
	Author: Alcid3s

	This software is strictly for educational purposes only. The author is not responsible for any damage caused by this software.
	A simple decryptor for windows that decrypts the files on the home folder of the user.
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/common-nighthawk/go-figure"
	"ransomware.com/main/internal/colors"
	"ransomware.com/main/internal/decryptor"
)

const Version = "1.0"
const Disclaimer = colors.ColorRed + "Disclaimer: " + colors.ColorBlue + "\nThis code is a decryptor and is for educational purposes only. " +
	"Do not use it for malicious purposes. the creator not responsible for any damage caused by this code." + colors.ColorReset

func titleScreen() {
	title := figure.NewFigure("Optimistic  Pickle  "+Version, "puffy", true)
	title.Print()
	fmt.Println(Disclaimer)
}

func main() {
	titleScreen()
	if runtime.GOOS != "windows" {
		fmt.Println("Windows is required to run this decryptor. Exiting...")
		return
	}
	key := make([]byte, 32)
	fmt.Println("Key: ", key)
	decryptor.ExecuteDecrypt(filepath.Join(os.Getenv("USERPROFILE")), key)
}
