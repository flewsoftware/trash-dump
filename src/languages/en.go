package languages

// "english" language pack
func getPackEN() Language {
	return Language{
		Iso639Code: "en",
		ConsoleOutputMessages: ConsoleOutputMessages{
			CantWriteFile:     "Error cant write file",
			FileAlreadyExists: "File already exists\nUse the flag \\\"--force\\\" to overwrite the file",
			OverWritingFIle:   "Overwriting file because of the usage of the \\\"--force\\\" flag",
			FileWritten:       "File successfully written",
		},
	}
}
