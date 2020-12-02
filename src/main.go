package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"sync"
	"trash-dump/languages"
)

var M languages.ConsoleOutputMessages

func main() {
	var force bool
	var seed int64
	var noBytes uint64
	var fileName string
	var flagLanguage string
	tokenChan := make(chan []byte)

	// flags
	flag.Uint64Var(&noBytes, "noBytes", 10, "No of bytes to generate")
	flag.Int64Var(&seed, "seed", 1115, "Seed to generate random bytes")
	flag.BoolVar(&force, "force", false, "Overwrites file if it already exists")
	flag.StringVar(&fileName, "filename", "trash-dump.tdp", "Name of the file")
	flag.StringVar(&flagLanguage, "language", "", "Language")
	flag.Parse()

	//language
	language, langErr := languages.GetLanguage()
	println("System Language: ", language)

	if flagLanguage == "" {
		if langErr != nil {
			println("Language couldn't be detected")
		} else {
			M = languages.GetLanguagePack(language)
		}
	} else {
		M = languages.GetLanguagePack(flagLanguage)
		println("Using Language: " + flagLanguage)
	}

	// generates a byte array
	token := make([]byte, noBytes)
	fileName = "./" + fileName + ".tdp"

	var wg sync.WaitGroup
	wg.Add(1)
	go generateRandomBytes(seed, token, tokenChan, &wg)
	token = <-tokenChan

	// waits till random bytes are generated
	wg.Wait()

	writeFileMain := func() {
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(M.CantWriteFile + " " + err.Error())
		}
		defer file.Close()
		writeRandomBytesToFile(file, token)
	}

	// Checks if file exists
	val := Exists(fileName)
	if !val {
		writeFileMain()
	} else {
		if force {
			println(M.OverWritingFIle)
			writeFileMain()
		} else {
			println(M.FileAlreadyExists)
		}
	}

}

func writeRandomBytesToFile(file *os.File, token []byte) {
	_, err := file.Write(token)
	if err != nil {
		log.Fatal(M.CantWriteFile + " " + err.Error())
	}
	defer println(M.FileWritten)
}

func generateRandomBytes(seed int64, token []byte, tokenChan chan []byte, wg *sync.WaitGroup) {
	rand.Seed(seed)
	rand.Read(token)
	defer wg.Done()
	tokenChan <- token
}

func Exists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
