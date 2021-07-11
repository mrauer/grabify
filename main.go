package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	YOUTUBE_API_URL = "https://www.googleapis.com/youtube/v3/search?part=snippet&q=%s&key=%s"
	TIMEOUT_SECONDS = 2
	SOFTWARE_LOGO   = `
 ██████  ██████   █████  ██████  ██ ███████ ██    ██ 
██       ██   ██ ██   ██ ██   ██ ██ ██       ██  ██  
██   ███ ██████  ███████ ██████  ██ █████     ████   
██    ██ ██   ██ ██   ██ ██   ██ ██ ██         ██    
 ██████  ██   ██ ██   ██ ██████  ██ ██         ██    
                                                   `
	SOFTWARE_VERSION = "v0.0.1"
	YOUTUBEDL_CMD    = "youtube-dl --extract-audio --audio-format mp3 %s"
)

var songs map[int]string

type items struct {
	Id struct {
		VideoId string `json:"videoId"`
	}
	Snippet struct {
		Title string `json:"title"`
	}
}

type result struct {
	Items []items `json:"items"`
}

func main() {
	fmt.Println(SOFTWARE_LOGO)
	fmt.Println(fmt.Sprintf("version %s\n", SOFTWARE_VERSION))

	// ask the user.
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What song would you want to download? ")
	query, _ := reader.ReadString('\n')
	fmt.Println("")

	url := fmt.Sprintf(YOUTUBE_API_URL, url.QueryEscape(query), os.Getenv("YOUTUBE_API_KEY"))

	spaceClient := http.Client{
		Timeout: time.Second * TIMEOUT_SECONDS, // Timeout after 2 seconds.
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

	songs := make(map[int]string)
	for idx, song := range resp.Items {
		fmt.Println(fmt.Sprintf("%d - %s", idx, song.Snippet.Title))
		songs[idx] = song.Id.VideoId
	}

	fmt.Println("")
	fmt.Print("Choose between these songs: ")
	song_id_str, _ := reader.ReadString('\n')
	song_id, err := strconv.Atoi(strings.TrimSuffix(song_id_str, "\n"))
	fmt.Println("")

	// download the song.
	cmd := exec.Command("sh", "-c", fmt.Sprintf(YOUTUBEDL_CMD, songs[song_id]))
	cmd.Run()

}
