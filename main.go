package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	YOUTUBE_API_URL = "https://www.googleapis.com/youtube/v3/search?part=snippet&q=%s&key=%s"
	TIMEOUT_SECONDS = 2
)

type items struct {
	Id struct {
		VideoId string `json:"videoId"`
	}
}

type result struct {
	Items []items `json:"items"`
}

func main() {
	fmt.Println("Grabity")

	query := "how+bizarre"

	url := fmt.Sprintf(YOUTUBE_API_URL, query, os.Getenv("YOUTUBE_API_KEY"))

	spaceClient := http.Client{
		Timeout: time.Second * TIMEOUT_SECONDS, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	resp := result{}
	jsonErr := json.Unmarshal(body, &resp)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(resp.Items)
}
