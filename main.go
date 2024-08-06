package main

import (
	"fmt"
	"net/http"
	"log"
	"strconv"
)

func handleOperations(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Request received\n")
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	num1Str := r.FormValue("num1")
	num2Str := r.FormValue("num2")
	operation := r.FormValue("operation")
    fmt.Printf("num1: %T, num2: %T, operation: %T\n", num1Str, num2Str, operation)
	num1, err := strconv.Atoi(num1Str)
	if err != nil {
		http.Error(w, "Invalid number for num1", http.StatusBadRequest)
		return
	}
	num2, err := strconv.Atoi(num2Str)
	if err != nil {
		http.Error(w, "Invalid number for num2", http.StatusBadRequest)
		return
	}

	var result int
	switch operation {
	case "add":
		result = num1 + num2
	case "sub":
		result = num1 - num2
	case "mul":
		result = num1 * num2
	case "div":
		if num2 == 0 {
			http.Error(w, "Division by zero", http.StatusBadRequest)
			return
		}
		result = num1 / num2
	default:
		http.Error(w, "Unknown operation", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Result: %d", result)
}

func main() {
	http.HandleFunc("/form", handleOperations)
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/",fileServer)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
