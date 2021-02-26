package main

import (
	"fmt"
	"os"
	"time"

	"../zip"

	"google.golang.org/api/drive/v3"
)

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
			"WorldData_%04d/%02d/%02d_%02d:%02d",
			currentTime.Year(),
			currentTime.Month(),
			currentTime.Day(),
			currentTime.Hour(),
			currentTime.Minute(),
		),
		Parents:     []string{Settings.DriveParentDir},
		Description: "World backup data",
	}

	var uploadedFileInfo *drive.File
	uploadedFileInfo, err = Drive.Files.Create(newFileInfo).Media(f).Do()

	os.Remove(TemporaryArchvePath)

	Discord.SendMessage(makeSendMessage(uploadedFileInfo))
}
