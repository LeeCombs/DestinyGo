package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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

		statResp := make(map[string]map[string]string)

		for _, c := range chars {
			statResp[c["CharacterID"].(string)] = controllers.GetHistoricalStats(dPlayers[0]["membershipID"], c["CharacterID"].(string))
		}

		c.HTML(http.StatusOK, "searchUser.tmpl.html", gin.H{
			"iconPath":    dPlayers[0]["iconPath"],
			"displayName": dName,
			"gScore":      gScore,
			"chars":       chars,
			"statResp":    statResp,
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
