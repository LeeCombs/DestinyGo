package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "errors"
    "./models"
)

const (
    ApiUrl = "https://www.bungie.net/Platform/Destiny/"
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

    // Get the membership ID and display the account's character summaries
    memID, _ := getMembershipIdByDisplayName(DisplayName)
    characterSummary(memID)
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
func getBody(url string) ([]byte, error) {
    // Build the request
    req, reqErr := http.NewRequest("GET", url, nil)
    if reqErr != nil {
        log.Fatal("reqErr: ", reqErr)
        return nil, errors.New("Error building NewRequest")
    }

    // Grab the API key to add to the request
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

    return body, nil
}

// Retrieve the membership ID associated with the displayName
func getMembershipIdByDisplayName(displayName string) (string, error) {
    url := ApiUrl + MembershipType + "/Stats/GetMembershipIdByDisplayName/" + displayName

    // Get and parse the response body
    var body, bodyErr = getBody(url)
    if bodyErr != nil {
        log.Fatal("bodyErr: ", bodyErr)
        return "", errors.New("Error getting body")
    }

    var bodyData map[string]interface{}
    jsonErr := json.Unmarshal(body, &bodyData)
    if jsonErr != nil {
        log.Fatal("jsonErr: ", jsonErr)
        return "", errors.New("Error Unmarshalling body")
    }

    // Return the body's response value, as that's the membership ID
    return bodyData["Response"].(string), nil
}

// Display the account's character summary
func characterSummary(destinyMembershipID string) {
    url := ApiUrl + MembershipType + "/Account/" + destinyMembershipID + "/Summary/"

    // Get and parse the response body
    var body, bodyErr = getBody(url)
    if bodyErr != nil {
        log.Fatal("bodyErr: ", bodyErr)
        return 
    }

    var dRes models.AccountSummaryResponse
    jErr := json.Unmarshal(body, &dRes)
    if jErr != nil {
        log.Fatal("jErr: ", jErr)
        return
    }

    // Display all Character IDs for now
    fmt.Println("==========")
    fmt.Println(dRes.Response.Data.MembershipID)
    fmt.Println(dRes.Response.Data.MembershipType)
    for i, e := range dRes.Response.Data.Characters {
        fmt.Println(i, e.CharacterBase.CharacterID)
    }
    fmt.Println("==========")

}
