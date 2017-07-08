package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"DestinyGo/controllers"

	"gopkg.in/gin-gonic/gin.v1"
)

var (
	DisplayName = "UserNameHere"
)

func main() {
	fmt.Println("Runnin' the program")

	// Temporary
	// Grab the display name from a local file, for now
	DisplayName, _ = getDName()
	searchChars, _ := controllers.SearchDestinyPlayer(DisplayName)

	// Setup the router
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	// Set up routes
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
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
		// Parse the form and grab the supplied user name
		c.Request.ParseForm()
		dName := c.Request.PostForm["displayName"][0]

		fmt.Println("Searching for:", dName)

		// Search for the player
		dPlayers, _ := controllers.SearchDestinyPlayer(dName)
		if len(dPlayers) <= 0 {
			fmt.Println("No player(s) found")
			c.String(http.StatusNotFound, "Player not found")
			return
		}

		// Grab the account summary and populate the character info
		gScore, chars, _ := controllers.GetAccountSummary(dPlayers[0]["membershipID"], true)

		// stats = map[characterID]map[stat]value
		stats := make(map[string]map[string]string)

		for _, c := range chars {
			charID := c["CharacterID"].(string)
			fmt.Println(charID)

			statResp := make(map[string]string)
			statResp = controllers.GetHistoricalStats(dPlayers[0]["membershipID"], charID)

			// Doing a little additional remapping of character data to this output
			stats[charID] = make(map[string]string)
			stats[charID] = statResp
			stats[charID]["CharacterID"] = charID
			stats[charID]["CharacterLevel"] = strconv.Itoa(c["CharacterLevel"].(int))
			stats[charID]["PowerLevel"] = strconv.Itoa(c["PowerLevel"].(int))
			stats[charID]["EmblemPath"] = c["EmblemPath"].(string)
			stats[charID]["BackgroundPath"] = c["BackgroundPath"].(string)
			stats[charID]["ClassName"] = c["ClassName"].(string)
			stats[charID]["RaceName"] = c["RaceName"].(string)
			stats[charID]["GenderName"] = c["GenderName"].(string)

			stats[charID]["AutoRifle"] = c["AutoRifle"].(string)
			stats[charID]["HandCannon"] = c["HandCannon"].(string)
			stats[charID]["PulseRifle"] = c["PulseRifle"].(string)
			stats[charID]["ScoutRifle"] = c["ScoutRifle"].(string)

			stats[charID]["FusionRifle"] = c["FusionRifle"].(string)
			stats[charID]["Shotgun"] = c["Shotgun"].(string)
			stats[charID]["SideArm"] = c["SideArm"].(string)
			stats[charID]["Sniper"] = c["Sniper"].(string)

			stats[charID]["MachineGun"] = c["MachineGun"].(string)
			stats[charID]["RocketLauncher"] = c["RocketLauncher"].(string)
			stats[charID]["Sword"] = c["Sword"].(string)

			stats[charID]["Grenade"] = c["Grenade"].(string)
			stats[charID]["Melee"] = c["Melee"].(string)
			stats[charID]["Relic"] = c["Relic"].(string)
			stats[charID]["Super"] = c["Super"].(string)

			stats[charID]["SubMachineGun"] = c["SubMachineGun"].(string) // Not used

		}
		fmt.Println("stats", stats)

		c.HTML(http.StatusOK, "searchUser.tmpl.html", gin.H{
			"iconPath":    dPlayers[0]["iconPath"],
			"displayName": dName,
			"gScore":      gScore,
			"stats":       stats,
		})
	}

	return gin.HandlerFunc(fn)
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
