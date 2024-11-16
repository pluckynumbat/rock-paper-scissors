package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("welcome to the rock-paper-client...")

	fmt.Println("Please enter the server URL: ")
	serverURL := ""
	_, err := fmt.Scanln(&serverURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Please enter the port number: ")
	portNumber := ""
	_, err = fmt.Scanln(&portNumber)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	option := ""
	for {
		fmt.Println("Options:")
		fmt.Println("Press 1 to play a random game")
		fmt.Println("Press any other key to exit")

		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if option != "1" {
			return
		}

		resp, err := http.Get("http://" + serverURL + ":" + portNumber + "/random")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		defer resp.Body.Close()

		printServerResponseDetails(resp)
	}
}

func printServerResponseDetails(resp *http.Response) {
	fmt.Println("Response Status: ", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
