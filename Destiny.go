package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func main() {
	destinyUrl := "http://www.bungie.net/Platform/Destiny/1/Stats/GetMembershipIdByDisplayName/UserNameHere/"

	// Build the request
	req, reqErr := http.NewRequest("GET", destinyUrl, nil)
	if reqErr != nil {
		log.Fatal("reqErr: ", reqErr)
		return
	}
	req.Header.Add("X-API-Key", "apiKeyHere")

	// Build the client and send the request
	client := &http.Client{}
	resp, respErr := client.Do(req)
	if respErr != nil {
		log.Fatal("respErr: ", respErr)
		return
	}

	// Close the body
	defer resp.Body.Close()

	// Read the content into a byte array
	body, bodyErr := ioutil.ReadAll(resp.Body)
	if bodyErr != nil {
		log.Fatal("bodyErr: ", bodyErr)
		return
	}

	// Convert the body into something usable
	var bodyData map[string]interface{}
	jsonErr := json.Unmarshal(body, &bodyData)
	if jsonErr != nil {
		log.Fatal("jsonErr: ", jsonErr)
		return
	}

	// Show off some stuff
	fmt.Println(body)
	fmt.Println(string(body))

	fmt.Println(bodyData)
	fmt.Println(bodyData["Response"])
	fmt.Println(bodyData["Message"])
}