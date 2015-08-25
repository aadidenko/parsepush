package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/facebookgo/parse"
)

var (
	appID     = ""
	masterKey = ""
	alert     = ""
	badge     = 1
	help      = false

	pushURL = &url.URL{
		Scheme: "https",
		Host:   "api.parse.com",
		Path:   "/1/push",
	}
)

// Payload represents notfication payload
type Payload struct {
	Alert string `json:"alert"`
	Badge int    `json:"badge"`
}

// ParsePush represents Parse.com push notification options
type ParsePush struct {
	Channels []string `json:"channels"`
	Data     Payload  `json:"data"`
}

func init() {
	flag.StringVar(&appID, "app-id", "", "Parse.com Application ID")
	flag.StringVar(&masterKey, "master-key", "", "Parse.com Master Key")
	flag.StringVar(&alert, "alert", "", "An alert message to display to the user")
	flag.IntVar(&badge, "badge", 1, "A number to badge the app icon with")
	flag.BoolVar(&help, "help", false, "Help Usage")
}

func usage() {
	fmt.Println(`Usage:

    parsepush [arguments]

Arguments:

    -app-id      Parse.com Application ID (required)
    -master-key  Parse.com Master Key (required)
    -alert       An alert message to display to the user (required)
    -badge       A number to badge the app icon with (defailt: 1)
`)
}

func main() {
	flag.Parse()

	if help {
		usage()
		os.Exit(0)
	}

	if appID == "" {
		fmt.Printf("Error: Cannot use empty -app-id\n\n")
		usage()
		os.Exit(1)
	}

	if masterKey == "" {
		fmt.Printf("Error: Cannot use empty -master-key\n\n")
		usage()
		os.Exit(1)
	}

	if alert == "" {
		fmt.Printf("Error: Cannot use empty -alert\n\n")
		usage()
		os.Exit(1)
	}

	fmt.Println("Sending...")

	credentials := parse.MasterKey{
		ApplicationID: appID,
		MasterKey:     masterKey,
	}

	parseCli := &parse.Client{Credentials: credentials}

	body := &ParsePush{
		Channels: []string{"global"},
		Data:     Payload{Alert: alert, Badge: badge},
	}

	_, err := parseCli.Post(pushURL, body, nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
