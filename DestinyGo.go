package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "DestinyGo/constants"
	"DestinyGo/models"

	"gopkg.in/gin-gonic/gin.v1"
)

var (
	ApiUrl         = "https://www.bungie.net/Platform/Destiny/"
	MembershipType = 1 // Xbox = 1, PSN = 2
	DisplayName    = "UserNameHere"
)

func main() {
	fmt.Println("Runnin' the program")

	// Temporary
	// Grab the display name from a local file, for now
	DisplayName, _ = getDName()
	searchChars, _ := SearchDestinyPlayer(DisplayName)

	// Setup the router
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	// Set up routes
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":       "DestinyGo",
			"displayName": DisplayName,
			"chars":       searchChars,
		})
	})

	router.POST("/searchUser", handleSearch())

	// Serve
	port := os.Getenv("PORT")
	if port == "" {
		port = "8787"
	}
	router.Run(":" + port)

}

//////////////
// Handlers //
//////////////

func handleSearch() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.Request.ParseForm()

		fmt.Println(c.Request.PostForm)

		for k, v := range c.Request.PostForm {
			fmt.Println("kv", k, v)
		}

		dName := c.Request.PostForm["displayName"][0]
		fmt.Println("dName", dName)

		dPlayers, _ := SearchDestinyPlayer(dName)
		fmt.Println("dPlayers", dPlayers)

		if len(dPlayers) <= 0 {
			fmt.Println("No player(s) found")
			c.String(http.StatusNotFound, "Player not found")
			return
		}

		gScore, chars, _ := GetAccountSummary(dPlayers[0]["membershipID"], false)
		for i, e := range chars {
			fmt.Println(i, e["CharacterID"])
		}

		c.HTML(http.StatusOK, "searchUser.tmpl", gin.H{
			"iconPath":    dPlayers[0]["iconPath"],
			"displayName": dName,
			"gScore":      gScore,
			"chars":       chars,
		})
	}

	return gin.HandlerFunc(fn)
}

/////////////////
// Destiny API //
/////////////////

// Search for Destiner players by display name
// PUBLIC Endpoint
// Returns a 2D array of [[IconPath, MembershipID]]
// TODO: Convert this to a map return
func SearchDestinyPlayer(displayName string) ([]map[string]string, error) {
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
	retVal := []map[string]string{}
	for _, e := range dRes.Response {
		retVal = append(retVal, map[string]string{"iconPath": e.IconPath, "membershipID": e.MembershipID})
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
// Returns Account's Grimoire Score and basic Character information
func GetAccountSummary(destinyMembershipId string, definitions bool) (int, []map[string]interface{}, error) {
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
	// Currently interested in: CharacterID, CharacterLevel, PowerLevel, EmblemPath, BackgroundPath
	chars := []map[string]interface{}{}
	for _, e := range dRes.Response.Data.Characters {
		chars = append(chars, map[string]interface{}{
			"CharacterID":    e.CharacterBase.CharacterID,
			"CharacterLevel": e.BaseCharacterLevel,
			"PowerLevel":     e.CharacterBase.PowerLevel,
			"EmblemPath":     e.EmblemPath,
			"BackgroundPath": e.BackgroundPath,
		})
	}

	grimoireScore := dRes.Response.Data.GrimoireScore
	return grimoireScore, chars, nil
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
func GetCharacterInventorySummary(destinyMembershipId string, characterId string, definitions bool) {
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

// TEMP - Grab a display name from a local file
func getDName() (string, error) {
	dName := os.Getenv("DISPLAYNAME")

	// Temp - If unable to get OS env, just read from local text
	if dName == "" {
		readKey, readErr := ioutil.ReadFile("DestinyName.txt")
		if readErr != nil {
			log.Fatal("readErr: ", readErr)
			return "", errors.New("Error reading DestinyName.txt")
		}
		return string(readKey), nil
	}

	return dName, nil
}

// Grab an API key from a local file
func getAPIKey() (string, error) {
	aKey := os.Getenv("APIKEY")

	// Temp - If unable to get OS env, just read from local text
	if aKey == "" {
		readKey, readErr := ioutil.ReadFile("DestinyKey.txt")
		if readErr != nil {
			log.Fatal("readErr: ", readErr)
			return "", errors.New("Error reading DestinyKey.txt")
		}
		return string(readKey), nil
	}

	return aKey, nil
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
