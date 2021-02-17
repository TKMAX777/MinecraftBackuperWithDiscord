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
	"../zip"

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
		if t.Hour() == Settings.UploadTime.Hour &&
			t.Minute() == Settings.UploadTime.Minute {
			backup()
		}
		time.Sleep(time.Minute)
	}
}

func backup() {
	var err error
	var currentTime = time.Now()

	err = zip.Compress(TemporaryArchvePath, Settings.FilePath)
	if err != nil {
		ErrorHandle(err, "Compress")
		return
	}

	f, err := os.Open(TemporaryArchvePath)
	if err != nil {
		ErrorHandle(err, "Opening temporary archive file")
		return
	}
	defer f.Close()

	var newFileInfo = &drive.File{
		Name: fmt.Sprintf(
			"WorldData_%04d/%02d/%02d",
			currentTime.Year(),
			currentTime.Month(),
			currentTime.Day(),
		),
		Parents:     []string{Settings.DriveParentDir},
		Description: "World backup data",
	}

	var uploadedFileInfo *drive.File
	uploadedFileInfo, err = Drive.Files.Create(newFileInfo).Media(f).Do()

	os.Remove(TemporaryArchvePath)

	Discord.SendMessage(makeSendMessage(uploadedFileInfo))
}
