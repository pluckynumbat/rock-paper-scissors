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
		fmt.Println("Press 1 to play a random choice against the server")
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
			sendServerRequest(serverURL, portNumber, "play-random")

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

func sendPlayRockRequest(serverURLPrefix string) (string, error) {
	res, err := sendServerRequest(serverURLPrefix, "/play-rock")
	if err != nil {
		return "", fmt.Errorf("Error in send play rock request: %v", err)
	}

	return res, nil
}

func createServerURLPrefix(serverAddr, portNumber string) string {
	return "http://" + serverAddr + ":" + portNumber
}

func sendServerRequest(serverURL, endpoint string) (string, error) {
	result := ""

	resp, err := http.Get(serverURL + "/" + endpoint)
	if err != nil {
		return result, fmt.Errorf("error in http request, error: %v", err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		result += fmt.Sprintln(scanner.Text())
	}

	return result, nil
}
