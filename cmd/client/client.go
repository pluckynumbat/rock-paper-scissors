package main

import (
	"bufio"
	"fmt"
	"net/http"
)

const escapeString string = "exit"

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

	serverURLPrefix := createServerURLPrefix(serverURL, portNumber)

	done := make(chan bool)
	go runGameLoop(serverURLPrefix, done)
	<-done
}

func runGameLoop(serverURLPrefix string, finished chan bool) {

	for {
		result, err := provideOptions(serverURLPrefix, "")
		if err != nil {
			fmt.Println(err)
			break
		} else if result != escapeString {
			fmt.Println(result)
		} else {
			break
		}
	}

	finished <- true
}

func provideOptions(serverURLPrefix string, currentInput string) (string, error) {

	fmt.Println("Options:")
	fmt.Println("Press 1 to play a random choice against the server")
	fmt.Println("Press 'R' or 'r' to play Rock against the server")
	fmt.Println("Press 'P' or 'p' to play Paper against the server")
	fmt.Println("Press 'S' or 's' to play Scissors against the server")
	fmt.Println("Press any other key to exit")

	option := currentInput
	if option == "" {
		_, err := fmt.Scanln(&option)
		if err != nil {
			return "", fmt.Errorf("Scan Line Error %v:", err)
		}
	}

	result := ""
	var err error

	switch option {
	case "1":
		result, err = sendPlayRandomRequest(serverURLPrefix)

	case "R", "r":
		result, err = sendPlayRockRequest(serverURLPrefix)

	case "P", "p":
		result, err = sendPlayPaperRequest(serverURLPrefix)

	case "S", "s":
		result, err = sendPlayScissorsRequest(serverURLPrefix)

	default:
		result = escapeString
	}

	if err != nil {
		return "", fmt.Errorf("Request Error: %v", err)
	} else {
		return result, nil
	}
}

func sendPlayRockRequest(serverURLPrefix string) (string, error) {
	res, err := sendServerRequest(serverURLPrefix, "play-rock")
	if err != nil {
		return "", fmt.Errorf("Error in send play rock request: %v", err)
	}

	return res, nil
}

func sendPlayPaperRequest(serverURLPrefix string) (string, error) {
	res, err := sendServerRequest(serverURLPrefix, "play-paper")
	if err != nil {
		return "", fmt.Errorf("Error in send play paper request: %v", err)
	}

	return res, nil
}

func sendPlayScissorsRequest(serverURLPrefix string) (string, error) {
	res, err := sendServerRequest(serverURLPrefix, "play-scissors")
	if err != nil {
		return "", fmt.Errorf("Error in send play scissors request: %v", err)
	}

	return res, nil
}

func sendPlayRandomRequest(serverURLPrefix string) (string, error) {
	res, err := sendServerRequest(serverURLPrefix, "play-random")
	if err != nil {
		return "", fmt.Errorf("Error in send play random request: %v", err)
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
