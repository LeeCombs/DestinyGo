package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "errors"
    "./models"
    "strconv"
    "./constants"
)

var (
    ApiUrl = "https://www.bungie.net/Platform/Destiny/"
    MembershipType = 1 // Xbox = 1, PSN = 2
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
    memID, _ := GetMembershipIdByDisplayName(DisplayName, false)

    // Getting an Account Summary
    /*
    grimScore, characters, accSumErr := GetAccountSummary(memID, false)
    if accSumErr != nil {
        log.Fatal("accSumErr", accSumErr)
        return
    }
    fmt.Println("Grimoire Score", grimScore)
    for _, char := range characters {
        fmt.Println()
        fmt.Println("Character:", char.CharacterBase.CharacterID)
        fmt.Println("Level:", char.CharacterLevel)
        fmt.Println("Light:", char.CharacterBase.PowerLevel)
        fmt.Println("EmblemPath:", char.EmblemPath)
        fmt.Println("BackgroundPath:", char.BackgroundPath)
    }
    */

    // Getting a Character Summary
    /*
    char, charErr := GetCharacterSummary(memID, "2305843009333417637", false)
    if charErr != nil {
        log.Fatal("charErr", charErr)
        return
    }
    fmt.Println("Character:", char.CharacterBase.CharacterID)
    fmt.Println("Level:", char.CharacterLevel)
    fmt.Println("Light:", char.CharacterBase.PowerLevel)
    fmt.Println("EmblemPath:", char.EmblemPath)
    fmt.Println("BackgroundPath:", char.BackgroundPath)
    */

    // Searching for a player
    /*
    dPlayers, searchErr :=  SearchDestinyPlayer(DisplayName)
    if searchErr != nil {
        log.Fatal("searchErr: ", searchErr)
        return 
    }
    fmt.Println(dPlayers)
    */

    // GetCharacterProgression(memID, "2305843009333417637", false)
    // GetCharacterInventorySummary(memID, "2305843009333417637", false)
    // GetAllItemsSummary(memID, false)

    GetDestinySingleDefinition(int(constants.DefinitionTypeInventoryItem), "2878029263", false)
    fmt.Println(memID)

}

// Search for Destiner players by display name
// PUBLIC Endpoint
// Returns a 2D array of [[IconPath, MembershipID]]
func SearchDestinyPlayer(displayName string) ([][]string, error) {
    // Build the uri
    // SearchDestinyPlayer/{membershipType}/{displayName}/
    uri := "SearchDestinyPlayer/" + strconv.Itoa(MembershipType) + "/" + displayName

    // Get and parse the response body
    var dRes models.SearchDestinyPlayer
    dataErr := getData(uri, &dRes)
    if dataErr != nil {
        log.Fatal("dataErr: ", dataErr)
        return nil, errors.New("Error retrieving dRes")
    }

    // Build the return (Icon Path, Member ID)
    retVal := [][]string{};
    for _, e := range dRes.Response {
        retVal = append(retVal, []string{e.IconPath, e.MembershipID})
    }
    return retVal, nil
}

// Retrieve the membership ID associated with the displayName
// PUBLIC Endpoint
// Returns a string of the Membership ID
func GetMembershipIdByDisplayName(displayName string, ignoreCase bool) (string, error) {
    // Build the uri
    // {membershipType}/Stats/GetMembershipIdByDisplayName/{displayName}
    uri := strconv.Itoa(MembershipType) + "/Stats/GetMembershipIdByDisplayName/" + displayName
    if ignoreCase {
        uri += "?ignorecase=true"
    }

    // Get and parse the response
    var dRes models.MembershipIdByDisplayName
    dataErr := getData(uri, &dRes)
    if dataErr != nil {
        log.Fatal("dataErr: ", dataErr)
        return "", errors.New("Error retrieving data")
    }

    // Return the body's response value, as that's the membership ID
    return dRes.Response, nil
}

// Retrieve the summary information for a given membership id
// PUBLIC Endpoint
// Returns Account's Grimoire Score and all Characters
func GetAccountSummary(destinyMembershipId string, definitions bool) (int, []models.Character, error) {
    // Build the uri
    // {membershipType}/Account/{destinyMembershipId}/Summary
    uri := strconv.Itoa(MembershipType) + "/Account/" + destinyMembershipId + "/Summary"
    if definitions {
        uri += "?definitions=true"
    }

    // Get and parse the response
    var dRes models.AccountSummaryResponse
    dataErr := getData(uri, &dRes)
    if dataErr != nil {
        log.Fatal("dataErr: ", dataErr)
        return -1, nil, errors.New("Error retrieving data")
    }

    // Build the return
    // TODO: Only interested in the following; think about only returning these instead of entire characters
    // Account Grimoire Score
    // Character's CharacterID, CharacterLevel, PowerLevel, EmblemPath, BackgroundPath
    grimoireScore := dRes.Response.Data.GrimoireScore
    characters := dRes.Response.Data.Characters
    return grimoireScore, characters, nil
}

// Get a singular character inventory summary information
// PUBLIC Endpoint -- Note: To get a more detailed overview, use GetDestinyAccountCharacterComplete
// Returns a Character
func GetCharacterSummary(destinyMembershipId string, characterId string, definitions bool) (models.Character, error) {
    // Build the uri
    // {membershipType}/Account/{destinyMembershipId}/Character/{characterId}/
    uri := strconv.Itoa(MembershipType) + "/Account/" + destinyMembershipId + "/Character/" + characterId
    if definitions {
        uri += "?definitions=true"
    }

    // Get and parse the response
    var dRes models.CharacterSummary
    dataErr := getData(uri, &dRes)
    if dataErr != nil {
        log.Fatal("dataErr: ", dataErr)
        return models.Character{}, errors.New("Error retrieving data")
    }

    // Return the Character
    // TODO: Only interested in the following; think about only returning these instead an entire character
    // CharacterID, CharacterLevel, PowerLevel, EmblemPath, BackgroundPath
    return dRes.Response.Data, nil
}

// Get a singular character inventory summary information
// PUBLIC Endpoint
// TODO: Actually return the information?
func GetCharacterInventorySummary (destinyMembershipId string, characterId string, definitions bool) {
    // Build the uri
    // {membershipType}/Account/{destinyMembershipId}/Character/{characterId}/Inventory/Summary
    uri := strconv.Itoa(MembershipType) + "/Account/" + destinyMembershipId + "/Character/" + characterId + "/Inventory/Summary"
    if definitions {
        uri += "?definitions=true"
    }

    // Get and parse the response
    var dRes models.CharacterInventorySummary
    dataErr := getData(uri, &dRes)
    if dataErr != nil {
        log.Fatal("dataErr: ", dataErr)
        return
    }

    // Show off some stuff
    fmt.Println("===CharInvSum===")
    fmt.Println("Char:", characterId)
    for i, e := range dRes.Response.Data.Items {
        fmt.Println("Item", i, e.ItemHash, e.ItemID)
    }
    for i, e := range dRes.Response.Data.Currencies {
        fmt.Println("Cur", i, e.ItemHash, e.Value)
    }
    fmt.Println("================")
}

// Get a singular character progression information
// PUBLIC Endpoint
// TODO: Actually return the information?
func GetCharacterProgression(destinyMembershipId string, characterId string, definitions bool) {
    // Build the uri
    // {membershipType}/Account/{destinyMembershipId}/Character/{characterId}/Progression/
    uri := strconv.Itoa(MembershipType) + "/Account/" + destinyMembershipId + "/Character/" + characterId + "/Progression"
    if definitions {
        uri += "?definitions=true"
    }

    // Get and parse the response body
    var dRes models.CharacterProgression
    dataErr := getData(uri, &dRes)
    if dataErr != nil {
        log.Fatal("dataErr: ", dataErr)
        return
    }

    // Show off some stuff
    fmt.Println("===CharPro===")
    fmt.Println("Char:", characterId)
    fmt.Println("Index, Hash, Current")
    for i, e := range dRes.Response.Data.Progressions {
        fmt.Println(i, e.ProgressionHash, e.CurrentProgress)
    }
    fmt.Println("=============")
}

// Returns all items for a given account
func GetAllItemsSummary(destinyMembershipId string, definitions bool) {
    // Build the uri
    // {membershipType}/Account/{destinyMembershipId}/Items
    uri := strconv.Itoa(MembershipType) + "/Account/" + destinyMembershipId + "/Items"
    if definitions {
        uri += "?definitions=true"
    }

    // Get and parse the response body
    var dRes models.AllItemsSummary
    dataErr := getData(uri, &dRes)
    if dataErr != nil {
        log.Fatal("dataErr: ", dataErr)
        return
    }

    fmt.Println("===AllItemSum===")
    for i, e := range dRes.Response.Data.Items {
        fmt.Println(i, e.ItemHash)
    }
    fmt.Println("================")
}

// Returns the specific item from the current manifest a json object
func GetDestinySingleDefinition(definitionType int, definitionID string, definitions bool) {
    // Build the uri
    // Manifest/{definitionType}/{definitionId}/
    uri := "/Manifest/" + strconv.Itoa(definitionType) + "/" + definitionID
    if definitions {
        uri += "?definitions=true"
    }

    // Get and parse the response body
    var dRes interface{}
    dataErr := getData(uri, &dRes)
    if dataErr != nil {
        log.Fatal("dataErr: ", dataErr)
        return
    }

    fmt.Println(dRes)

    // Note: dRes's data changes based on definitionType

    // For Inventory Items, we're iterested in:
    // ItemName, ItemDescription, Icon, HasIcon, SecondaryIcon
}

//////////////////////////////
// Private helper functions //
//////////////////////////////


// Grab an API key from a local file
func getAPIKey() (string, error) {
    readKey, readErr := ioutil.ReadFile("DestinyKey.txt")
    if readErr != nil {
        log.Fatal("readErr: ", readErr)
        return "", errors.New("Error reading DestinyKey.txt")
    }
    return string(readKey), nil
}

// Make a GET request using the supplied uri
func getBody(uri string) ([]byte, error) {
    // Build the request
    url := ApiUrl + uri
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

// Retrieve the stored data of a given uri, and populate the the supplied interface
func getData(uri string, dRes interface{}) error {
    // Get the body from the uri
    var body, bodyErr = getBody(uri)
    if bodyErr != nil {
        log.Fatal("bodyErr: ", bodyErr)
        return errors.New("Error retrieving body")
    }

    // Populate the interface with the retrieved body data
    jErr := json.Unmarshal(body, &dRes)
    if jErr != nil {
        log.Fatal("jErr: ", jErr)
        return errors.New("Error unmarshalling body")
    }

    return nil
}
