package providers

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// DotEnv structs

type DotEnv struct {
	Location   string
	LastLoaded struct {
		Key       string
		Value     string
		Retrieved time.Time
	}
}

var CurrentEnv = DotEnv{}

// Check if Env file updated more recently than last loaded value
func (d DotEnv) fileModMoreRecent() bool {
	file, err := os.Stat(d.Location)

	if err != nil {
		log.Fatalf("Error loading file stats: %s", err)
	}

	modifiedtime := file.ModTime()

	return (modifiedtime.After(d.LastLoaded.Retrieved))
}

// Retrieve value from Environment Variables file
func (d DotEnv) RetrieveValue(key string) string {

	if d.LastLoaded.Key == key && d.fileModMoreRecent() {
		return (d.LastLoaded.Value)
	}

	err := godotenv.Load(d.Location)

	if err != nil {
		log.Println("!!Error loading .env file!!")
		log.Printf("               Location: %s", d.Location)
		log.Printf("          Requested Key: %s", key)
		log.Fatalf("                  Error: %s", err)
	}

	d.LastLoaded.Key = key
	d.LastLoaded.Value = os.Getenv(key)
	d.LastLoaded.Retrieved = time.Now()

	return d.LastLoaded.Value
}
