package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	//count fot break
	count := 0
	for range time.Tick(15 * time.Second) {
		//break loop
		if count >= 5 {
			break
		}

		//call post req funv
		postReq()

		//increment for break
		count++

	}
}

func postReq() {
	//initiliaze data
	data := map[string]int{}
	// initiliaze range number
	min := 1
	max := 20
	randWater := (rand.Intn(max-min) + min)
	randWind := (rand.Intn(max-min) + min)

	//assign data
	data["water"] = randWater
	data["wind"] = randWind

	client := &http.Client{}

	jsonByte, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}

	buff := bytes.NewBuffer(jsonByte)

	//prepare
	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", buff)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	//do request
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	statWater := ""
	statWind := ""

	//water status
	switch {
	case randWater <= 5:
		statWater = "aman"
	case randWater > 8:
		statWater = "bahaya"
	default:
		statWater = "siaga"
	}

	//wind status
	switch {
	case randWind <= 6:
		statWind = "aman"
	case randWind > 15:
		statWind = "bahaya"
	default:
		statWind = "siaga"
	}

	log.Printf("%s \nstatus water : %s\nstatus wind : %s\n==============\n", string(body), statWater, statWind)
}
