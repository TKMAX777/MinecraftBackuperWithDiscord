package main

import "fmt"

// ErrorHandle export error information
func ErrorHandle(err error, function string) {
	var errMessage string = fmt.Sprintf("Error at %s: %s", function, err.Error())
	Discord.SendError(fmt.Errorf(errMessage))
	fmt.Println(errMessage)
	return
}
