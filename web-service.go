package main

import (
	"golang-verse/code"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/palindrom", services)

	fmt.Println("server started at localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func services(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // The Response content will be on JSon

	data := []struct {
		Palindrom string
		Count     string
	}{
		// This 3 lines are refers to the unit test on Palindrome code
		{"1 10", "9"},
		{"11 20", "1"},
		{"21 30", "1"},
	}

	dataJson, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Error response
		return
	}

	// In this Web Service, there are 2 Methods that used on, POST & GET Method
	switch r.Method {
	case "POST":
		r.ParseForm()
		postInput := r.FormValue("palindrom")
		result := strconv.Itoa(code.Utama(postInput))
		w.Write([]byte("There are " + result + " Palindrom counted")) // POST Method will respons on String format
	case "GET":
		w.Write(dataJson) // GET Method will respons on Json format
	default:
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

}

