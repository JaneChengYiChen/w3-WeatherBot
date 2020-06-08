package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	checkEnvFile()
	GetWeatherData("宜蘭縣")
}

func checkEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetWeatherData(place string) {
	var client http.Client

	req, err := http.NewRequest("GET", os.Getenv("weather"), nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()

	// add parameters
	q.Add("Authorization", os.Getenv("weather_auth"))
	q.Add("locationName", place)
	req.URL.RawQuery = q.Encode()

	res, _ := client.Do(req)

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}
}
