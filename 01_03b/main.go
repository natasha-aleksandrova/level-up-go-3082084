package main

import (
	"encoding/json"
	// "fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const path = "entries.json"

// raffleEntry is the struct we unmarshal raffle entries into
type raffleEntry struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

// importData reads the raffle entries from file and creates the entries slice.
func importData() []raffleEntry {
	var data []raffleEntry
	raw_data, err := os.ReadFile("01_03b/entries.json")
	if err != nil {
		log.Fatal("Failed with " + err.Error())
	}
	err = json.Unmarshal(raw_data, &data)
	if err != nil {
		log.Fatal("failed to unmarshall " + err.Error())
	}
	return data
}

// getWinner returns a random winner from a slice of raffle entries.
func getWinner(entries []raffleEntry) raffleEntry {
	rand.Seed(time.Now().Unix())
	wi := rand.Intn(len(entries))
	return entries[wi]
}

func main() {
	entries := importData()
	log.Println("And... the raffle winning entry is...")
	winner := getWinner(entries)
	time.Sleep(500 * time.Millisecond)
	log.Println(winner)
}