package main

import (
	"fmt"

	"google.golang.org/api/drive/v3"
)

func makeSendMessage(f *drive.File) string {
	var message string

	message += "定期バックアップを行いました。\n"
	message += fmt.Sprintf("FileName: %s\n", f.Name)

	fmt.Printf("%#v\n", f)

	return message
}
