package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"DestinyGo/models"
)

var (
	ApiUrl         = "https://www.bungie.net/Platform/Destiny/"
	MembershipType = 1 // Xbox = 1, PSN = 2
	DisplayName    = "UserNameHere"
	MiniManifest   models.MiniMani
)

// Read and load the local manifest file
func init() {
	file, fileErr := ioutil.ReadFile("./static/MiniMani.json")
	if fileErr != nil {
		log.Fatal("fileErr:", fileErr)
		return
	}

	jsonErr := json.Unmarshal(file, &MiniManifest)
	if jsonErr != nil {
		fmt.Println("jsonErr:", jsonErr)
	}
}

/////////////////
// Destiny API //
/////////////////

// Search for Destiner players by display name
// PUBLIC Endpoint
// Returns a 2D array of [[IconPath, MembershipID]]
func SearchDestinyPlayer(displayName string) ([]map[string]string, error) {
	// Build the uri
	// SearchDestinyPlayer/{membershipType}/{displayName}/
	uri := "SearchDestinyPlayer/" + strconv.Itoa(MembershipType) + "/" + displayName

	// Get and parse the response body
	var dRes models.SearchDestinyPlayer
	dataErr := getData(uri, &dRes)
	if dataErr != nil {
		log.Fatal("dataErr:", dataErr)
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

	// Build the character info
	chars := []map[string]interface{}{}
	for _, e := range dRes.Response.Data.Characters {
		chars = append(chars, map[string]interface{}{
			"CharacterID":    e.CharacterBase.CharacterID,
			"CharacterLevel": e.BaseCharacterLevel,
			"PowerLevel":     e.CharacterBase.PowerLevel,
			"EmblemPath":     e.EmblemPath,
			"BackgroundPath": e.BackgroundPath,
			"ClassName":      MiniManifest.DestinyClassDefinition[e.CharacterBase.ClassHash].ClassName,
			"RaceName":       MiniManifest.DestinyRaceDefinition[e.CharacterBase.RaceHash].RaceName,
			"GenderName":     MiniManifest.DestinyGenderDefinition[e.CharacterBase.GenderHash].GenderName,
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
	return dRes.Response.Data, nil
}

// Get a singular character inventory summary information
// PUBLIC Endpoint
// TODO: Actually return the information
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
// TODO: Actually return the information
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
		log.Fatal("dataErr:", dataErr)
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
		log.Fatal("dataErr:", dataErr)
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
		log.Fatal("dataErr:", dataErr)
		return
	}

	fmt.Println(dRes)

	// Note: dRes's data changes based on definitionType

	// For Inventory Items, we're iterested in:
	// ItemName, ItemDescription, Icon, HasIcon, SecondaryIcon
}

// Returns the specific item from the current manifest a json object
func GetHistoricalStats(destinyMembershipId string, characterID string) map[string]string {
	// Build the uri
	// Manifest/{definitionType}/{definitionId}/
	// Stats/{membershipType}/{destinyMembershipId}/{characterId}/
	uri := "/Stats/" + strconv.Itoa(MembershipType) + "/" + destinyMembershipId + "/" + characterID

	// Additional Query String params
	// ?periodType (0-4) PeriodType Enum
	// ?modes (0-36) DestinyActivityModeType Enum
	// ?groups (1-4) Some DestinyStatsGroupType Enum
	// ?monthstart - YYYY-MM
	// ?monthend - YYYY-MM
	// ?daystart - YYYY-MM-DD
	// ?dayend - YYYY-MM-DD

	// Get and parse the response body
	var dRes models.HistoricalResponse
	dataErr := getData(uri, &dRes)
	if dataErr != nil {
		log.Fatal("dataErr: ", dataErr)
		return nil
	}

	retVal := make(map[string]string)

	pvpStats := dRes.Response.AllPvP.AllTime

	retVal["Kills"] = pvpStats.Kills.Basic.DisplayValue
	retVal["Deaths"] = pvpStats.Deaths.Basic.DisplayValue
	retVal["Score"] = pvpStats.Score.Basic.DisplayValue
	retVal["CombatRating"] = pvpStats.CombatRating.Basic.DisplayValue
	retVal["LongestSpree"] = pvpStats.LongestKillSpree.Basic.DisplayValue
	retVal["LongestLife"] = pvpStats.LongestSingleLife.Basic.DisplayValue
	retVal["BestKills"] = pvpStats.BestSingleGameKills.Basic.DisplayValue
	retVal["BestScore"] = pvpStats.BestSingleGameScore.Basic.DisplayValue
	retVal["BestWeapon"] = pvpStats.WeaponBestType.Basic.DisplayValue
	retVal["OrbsDropped"] = pvpStats.OrbsDropped.Basic.DisplayValue
	retVal["Assists"] = pvpStats.Assists.Basic.DisplayValue

	// PvP stats
	retVal["AutoRifle"] = pvpStats.WeaponKillsAutoRifle.Basic.DisplayValue
	retVal["FusionRifle"] = pvpStats.WeaponKillsFusionRifle.Basic.DisplayValue
	retVal["Grenade"] = pvpStats.WeaponKillsGrenade.Basic.DisplayValue
	retVal["HandCannon"] = pvpStats.WeaponKillsHandCannon.Basic.DisplayValue
	retVal["MachineGun"] = pvpStats.WeaponKillsMachinegun.Basic.DisplayValue
	retVal["Melee"] = pvpStats.WeaponKillsMelee.Basic.DisplayValue
	retVal["PulseRifle"] = pvpStats.WeaponKillsPulseRifle.Basic.DisplayValue
	retVal["Relic"] = pvpStats.WeaponKillsRelic.Basic.DisplayValue
	retVal["RocketLauncher"] = pvpStats.WeaponKillsRocketLauncher.Basic.DisplayValue
	retVal["ScoutRifle"] = pvpStats.WeaponKillsScoutRifle.Basic.DisplayValue
	retVal["Shotgun"] = pvpStats.WeaponKillsShotgun.Basic.DisplayValue
	retVal["SideArm"] = pvpStats.WeaponKillsSideArm.Basic.DisplayValue
	retVal["Sniper"] = pvpStats.WeaponKillsSniper.Basic.DisplayValue
	retVal["SubMachineGun"] = pvpStats.WeaponKillsSubmachinegun.Basic.DisplayValue
	retVal["Super"] = pvpStats.WeaponKillsSuper.Basic.DisplayValue
	retVal["Sword"] = pvpStats.WeaponKillsSword.Basic.DisplayValue

	return retVal
}

//////////////////////////////
// Private helper functions //
//////////////////////////////

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

	ioutil.WriteFile("output.json", body, 0644)

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
