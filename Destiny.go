package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
)

const (
	ApiUrl = "http://www.bungie.net/Platform/Destiny/"
)

var (
	MembershipType = "1" // Xbox = 1, PSN = 2
	DisplayName = "UserNameHere"
)

func main() {
	// Temporary
	// Grab the display name from a local file, for now
	readKey, readErr := ioutil.ReadFile("DestinyName.txt")
	if readErr != nil {
		log.Fatal("readErr: ", readErr)
		return 
	}
	DisplayName = string(readKey)

	// Build the request
	destinyUrl := ApiUrl + MembershipType + "/Stats/GetMembershipIdByDisplayName/" + DisplayName

	// Convert the body into something usable
	var bodyData, bodyErr = getBodyData(destinyUrl)
	if bodyErr != nil {
		log.Fatal("bodyErr: ", bodyErr)
		return 
	}

	// Show off information
	for index, element := range bodyData {
		fmt.Printf("Type %T ", element)
		fmt.Println(index, element)
	}
}

// Grab an API key from a local file
func getAPIKey() (string, error) {
	readKey, readErr := ioutil.ReadFile("DestinyKey.txt")
	if readErr != nil {
		log.Fatal("readErr: ", readErr)
		return "", errors.New("Error reading DestinyKey.txt")
	}
	return string(readKey), nil
}

// Make a GET request to the supplied url
func getBodyData(url string) (map[string]interface{}, error) {
	req, reqErr := http.NewRequest("GET", url, nil)
	if reqErr != nil {
		log.Fatal("reqErr: ", reqErr)
		return nil, errors.New("Error building NewRequest")
	}

	apiKey, apiErr := getAPIKey()
	if apiErr != nil {
		log.Fatal("apiErr: ", apiErr)
		return nil, errors.New("Error getting API key")
	}

	req.Header.Add("X-API-Key", apiKey)

	// Build the client and send the request
	client := &http.Client{}
	resp, respErr := client.Do(req)
	if respErr != nil {
		log.Fatal("respErr: ", respErr)
		return nil, errors.New("Error making response")
	}

	// Close the body
	defer resp.Body.Close()

	// Read the content into a byte array
	body, bodyErr := ioutil.ReadAll(resp.Body)
	if bodyErr != nil {
		log.Fatal("bodyErr: ", bodyErr)
		return nil, errors.New("Error reading response body")
	}

	// Convert the body into something usable
	var bodyData map[string]interface{}
	jsonErr := json.Unmarshal(body, &bodyData)
	if jsonErr != nil {
		log.Fatal("jsonErr: ", jsonErr)
		return nil, errors.New("Error Unmarshalling body")
	}

	// Show off some stuff
	fmt.Println(body)
	fmt.Println(string(body))
	fmt.Println(bodyData)

	return bodyData, nil
}
