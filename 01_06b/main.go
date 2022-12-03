package main

import (
	"encoding/json"
	"log"
	"math"
	"os"
)

// User represents a user record.
type User struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

const path = "01_06b/users.json"

// getBiggestMarket takes in the slice of users and
// returns the biggest market.
func getBiggestMarket(users []User) (string, int) {
	var countryCnt = make(map[string]int)
	var maxCountry string
	var maxCount int = int(math.Inf(-1))
	for _, user := range users {
		countryCnt[user.Country]++
	}
	for k, v := range countryCnt {
		if v > maxCount{
			maxCount = v
			maxCountry = k
		}
	}
	return maxCountry, maxCount
}

func main() {
	users := importData()
	country, count := getBiggestMarket(users)
	log.Printf("The biggest user market is %s with %d users.\n",
		country, count)
}

// importData reads the raffle entries from file and
// creates the entries slice.
func importData() []User {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data []User
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
