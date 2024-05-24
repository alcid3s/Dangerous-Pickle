package obfuscator

import (
	"fmt"
	"io"
	"os"
)

func XorFile(inputFilename string, key []byte, outputFilename string) error {
	inputFile, err := os.Open(inputFilename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open input file: %v\n", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputFilename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open output file: %v\n", err)
		inputFile.Close()
		os.Exit(1)
	}
	defer outputFile.Close()

	fileInfo, err := inputFile.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get file info: %v\n", err)
		os.Exit(1)
	}
	fileSize := fileInfo.Size()

	buffer, err := io.ReadAll(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read input file: %v\n", err)
		os.Exit(1)
	}

	keyLength := len(key)
	for i := 0; i < int(fileSize); i++ {
		buffer[i] ^= key[i%keyLength]
	}

	_, err = outputFile.Write(buffer)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write to output file: %v\n", err)
		os.Exit(1)
	}

	return nil
}
