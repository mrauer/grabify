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
	SOFTWARE_VERSION = "v0.0.2"
	YOUTUBEDL_CMD    = "(cd %s && youtube-dl --extract-audio --audio-format mp3 --audio-quality 0 %s)"
	AMPLIFY_RATIO    = 4
	FILENAME_CMD     = "find %s/ | grep %s"
	AMPLIFY_CMD      = "ffmpeg -i \"%s\" -filter:a \"volume=%d\" \"%s\" && rm \"%s\""
	DOWNLOAD_PATH    = "data"
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

func amplifyAudio(song_id string) error {
	input_cmd, err := exec.Command("sh", "-c", fmt.Sprintf(FILENAME_CMD, DOWNLOAD_PATH, song_id)).Output()
	if err != nil {
		return err
	}
	input_file := strings.TrimSuffix(string(input_cmd), "\n")
	output_file := strings.ReplaceAll(input_file, ".mp3", fmt.Sprintf("_%dx.mp3", AMPLIFY_RATIO))
	cmd := exec.Command("sh", "-c", fmt.Sprintf(AMPLIFY_CMD, input_file, AMPLIFY_RATIO, output_file, input_file))
	cmd.Run()
	return nil
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
	cmd := exec.Command("sh", "-c", fmt.Sprintf(YOUTUBEDL_CMD, DOWNLOAD_PATH, songs[song_id]))
	cmd.Run()

	// amplifying the audio
	fmt.Println(fmt.Sprintf("Amplifying the audio by a ratio of %d", AMPLIFY_RATIO))

	err = amplifyAudio(songs[song_id])
	if err != nil {
		log.Fatal(err)
	}
}
