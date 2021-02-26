package main

// Setting put profile info
type Setting struct {
	DiscordHookURI      string
	ErrorDiscordHookURI string

	FilePath   string
	UploadTime []struct {
		Hour   int
		Minute int
	}
	DriveParentDir string

	MessageInfo      DiscordInfo
	ErrorMessageInfo DiscordInfo
}

// DiscordInfo put discord message information
type DiscordInfo struct {
	AvaterURI string
	UserName  string
}
