package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if len(os.Args) < 2 {
		fmt.Println("You need to pass a param")
		os.Exit(1)
	}
	arg1 := os.Args[1]

	baseURL := os.Getenv("BASE_URL")

	s := spinner.New(spinner.CharSets[8], 100*time.Millisecond)
	urlEncodedString := url.QueryEscape(arg1)
	myQueryString := baseURL + urlEncodedString

	req, err := http.NewRequest("GET", myQueryString, nil)
	if err != nil {
		log.Fatalf("Error creating the request: %v", err)
	}
	s.Start()

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(req)
	if err != nil {
		s.Stop()
		log.Fatalf("Error making the request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		s.Stop()
		log.Fatalf("Received non-200 response code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		s.Stop()
		log.Fatalf("Error reading the response: %v", err)
	}

	s.Stop()
	myString := string(body)
	fmt.Println(myString)

}
