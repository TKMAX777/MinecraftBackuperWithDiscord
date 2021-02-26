package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"../discord"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

// Drive Googl Drive handler
var Drive *drive.Service

// Discord put discord webhook handler
var Discord *discord.Handler

// Settings put program profile
var Settings Setting

// TemporaryArchvePath put tmporary path
var TemporaryArchvePath string

func init() {
	var err error

	TemporaryArchvePath = filepath.Join("tmp", "archive.zip")
	os.Mkdir("tmp", 0777)

	b, err := ioutil.ReadFile("settings.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	err = json.Unmarshal(b, &Settings)
	if err != nil {
		log.Fatalf("JSON parse error. Unable to read settings")
	}

	if Settings.FilePath == "" {
		log.Fatalf("Error: FilePath not specified.")
	}

	b, err = ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, drive.DriveScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	var client = getClient(config)

	Drive, err = drive.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	Discord = discord.NewHandler(Settings.DiscordHookURI)
	Discord.SetErrorHookURI(Settings.ErrorDiscordHookURI)

	Discord.SetErrorProfile(
		Settings.ErrorMessageInfo.AvaterURI,
		Settings.ErrorMessageInfo.UserName,
	)

	Discord.SetProfile(
		Settings.MessageInfo.AvaterURI,
		Settings.MessageInfo.UserName,
	)

	fmt.Printf("Start backup to Google Drive...\n")
}

func main() {
	var t time.Time
	for {
		t = time.Now()
		for _, ut := range Settings.UploadTime {
			if t.Hour() == ut.Hour && t.Minute() == ut.Minute {
				backup()
			}
		}

		time.Sleep(time.Minute)
	}
}
