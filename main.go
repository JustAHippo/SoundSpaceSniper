package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	println("\n██╗  ███╗░░██╗███████╗███████╗██████╗░  ░█████╗░  ██╗░░░░░██╗███████╗███████╗\n██║  ████╗░██║██╔════╝██╔════╝██╔══██╗  ██╔══██╗  ██║░░░░░██║██╔════╝██╔════╝\n██║  ██╔██╗██║█████╗░░█████╗░░██║░░██║  ███████║  ██║░░░░░██║█████╗░░█████╗░░\n██║  ██║╚████║██╔══╝░░██╔══╝░░██║░░██║  ██╔══██║  ██║░░░░░██║██╔══╝░░██╔══╝░░\n██║  ██║░╚███║███████╗███████╗██████╔╝  ██║░░██║  ███████╗██║██║░░░░░███████╗\n╚═╝  ╚═╝░░╚══╝╚══════╝╚══════╝╚═════╝░  ╚═╝░░╚═╝  ╚══════╝╚═╝╚═╝░░░░░╚══════╝")
	println("Made by Drowzee")
	print("TXT List: ")
	var filename string
	fmt.Scanln(&filename)
	var requestedPassword string
	print("Password for account: ")
	fmt.Scanln(&requestedPassword)
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		print("Trying username ", scanner.Text())
		status, responseStr, response := sendRequest(scanner.Text(), requestedPassword)
		if status == 200 && response.Success == true {
			println(" | Successfully grabbed!")
			passCombo := scanner.Text() + ":" + requestedPassword
			println("User Pass combo is " + passCombo)
			successFile, _ := os.OpenFile("success.txt", os.O_WRONLY|os.O_APPEND, 0644)
			defer successFile.Close()
			successFile.WriteString(passCombo + "\n")
			println("Response to request is as follows(debug purposes):")
			println(responseStr)
			println("Added to success.txt successfully! The next time you snipe, please use a VPN or different VPN IP due to now being blacklisted from making another account on the same IP.")
			os.Exit(0)
		} else if status == 200 && response.Success == false {
			println(" | Failed with message", response.Message)
		} else if status == 429 {
			println("\nError Code 429 shown, halting execution")
			os.Exit(1)
		} else {
			println(" | Unsuccessful ", status)
		}

	}

}

func sendRequest(username string, password string) (int, string, apiResp) {
	form := url.Values{"username": {username}, "password": {password}, "password2": {password}}
	requestURL := "https://soundspaceplus.dev/api/auth/register"
	req, _ := http.NewRequest("POST", requestURL, strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Host", "soundspaceplus.dev")
	req.Header.Set("Origin", "https://soundspaceplus.dev")
	resp, _ := http.DefaultClient.Do(req)
	shit, _ := ioutil.ReadAll(resp.Body)
	var apiResponse apiResp
	json.Unmarshal(shit, &apiResponse)
	return resp.StatusCode, string(shit), apiResponse
}

type apiResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
