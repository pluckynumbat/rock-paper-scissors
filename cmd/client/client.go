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
		fmt.Println("Press 'R' or 'r' to play Rock against the server")
		fmt.Println("Press 'P' or 'p' to play Paper against the server")
		fmt.Println("Press 'S' or 's' to play Scissors against the server")
		fmt.Println("Press any other key to exit")

		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		switch option {
		case "1":
			sendServerRequest(serverURL, portNumber, "random")

		case "R", "r":
			sendServerRequest(serverURL, portNumber, "play-rock")

		case "P", "p":
			sendServerRequest(serverURL, portNumber, "play-paper")

		case "S", "s":
			sendServerRequest(serverURL, portNumber, "play-scissors")

		default:
			return
		}
	}
}

func sendServerRequest(serverURL, portNumber, endpoint string) {
	resp, err := http.Get("http://" + serverURL + ":" + portNumber + "/" + endpoint)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer resp.Body.Close()

	printServerResponseDetails(resp)
}

func printServerResponseDetails(resp *http.Response) {
	fmt.Println("Response Status: ", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
