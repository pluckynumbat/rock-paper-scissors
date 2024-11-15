package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("welcome to the rock-paper-client...")

	resp, err := http.Get("http://localhost:8080/random")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status: ", resp.Status)
}
