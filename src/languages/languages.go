package languages

import (
	"github.com/cloudfoundry-attic/jibber_jabber"
	"log"
)

// Returns the system language
func GetLanguage() (string, error) {
	userLanguage, err := jibber_jabber.DetectLanguage()
	return userLanguage, err
}

// Returns console output messages according to the language
func GetLanguagePack(name string) ConsoleOutputMessages {
	switch name {
	case "en":
		return getPackEN().ConsoleOutputMessages
	case "si":
		return GetPackSI().ConsoleOutputMessages
	default:
		log.Fatal("Invalid language code")
		return getPackEN().ConsoleOutputMessages
	}
}

// language struct
type Language struct {
	Iso639Code            string
	ConsoleOutputMessages ConsoleOutputMessages
}

// console output messages struct
type ConsoleOutputMessages struct {
	CantWriteFile     string
	FileAlreadyExists string
	OverWritingFIle   string
	FileWritten       string
}
