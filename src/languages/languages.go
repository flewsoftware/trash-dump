package languages

import (
	"github.com/cloudfoundry-attic/jibber_jabber"
	"log"
)

func GetLanguage() (string, error) {
	userLanguage, err := jibber_jabber.DetectLanguage()
	println("System Language: ", userLanguage)
	return userLanguage, err
}

func GetLanguagePack(name string) ConsoleOutputMessages {
	switch name {
	case "en":
		return GetPackEN().ConsoleOutputMessages
	case "si":
		return GetPackSI().ConsoleOutputMessages
	default:
		log.Fatal("Invalid language code")
		return GetPackEN().ConsoleOutputMessages
	}
}

type Language struct {
	Iso639Code            string
	ConsoleOutputMessages ConsoleOutputMessages
}

type ConsoleOutputMessages struct {
	CantWriteFile     string
	FileAlreadyExists string
	OverWritingFIle   string
	FileWritten       string
}
