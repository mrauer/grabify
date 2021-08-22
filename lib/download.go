package lib

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	YOUTUBE_API_URL = "https://www.googleapis.com/youtube/v3/search?part=snippet&q=%s&type=video&key=%s"
	YOUTUBEDL_CMD   = "(cd %s && youtube-dl --extract-audio --audio-format mp3 --audio-quality 0 %s)"
	DOWNLOAD_PATH   = "data"
	AMPLIFY_RATIO   = 1
	TIMEOUT_SECONDS = 2
	FILENAME_CMD    = "find %s/ | grep \"%s\" | grep '.mp3'"
	AMPLIFY_CMD     = "ffmpeg -i \"%s\" -filter:a \"volume=%d\" \"%s\" && rm \"%s\""
)

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

func YoutubeSearch(query string, verbose bool) (map[int]string, error) {
	choices := make(map[int]string)
	url := fmt.Sprintf(YOUTUBE_API_URL, url.QueryEscape(query), os.Getenv("YOUTUBE_API_KEY"))
	client := http.Client{
		Timeout: time.Second * TIMEOUT_SECONDS, // Timeout after 2 seconds.
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return choices, err
	}

	res, err := client.Do(req)

	if err != nil {
		return choices, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return choices, err
	}

	resp := result{}
	err = json.Unmarshal(body, &resp)

	if err != nil {
		return choices, err
	}

	for idx, song := range resp.Items {
		if verbose {
			song_title := html.UnescapeString(song.Snippet.Title)
			fmt.Printf(fmt.Sprintf("%d - %s\n", idx, song_title))
		}

		choices[idx] = song.Id.VideoId
	}

	return choices, nil
}

func DownloadTrack(track_id string) error {
	fmt.Printf(fmt.Sprintf("Downloading %s\n", track_id))
	cmd := exec.Command("sh", "-c", fmt.Sprintf(YOUTUBEDL_CMD, DOWNLOAD_PATH, track_id))
	_ = cmd.Run()

	return nil
}

func AmplifyAudio(track_id string) error {
	input_cmd, err := exec.Command("sh", "-c", fmt.Sprintf(FILENAME_CMD, DOWNLOAD_PATH, track_id)).Output()
	if err != nil {
		return err
	}

	input_file := strings.TrimSuffix(string(input_cmd), "\n")
	output_file := strings.ReplaceAll(input_file, ".mp3", fmt.Sprintf("_%dx.mp3", AMPLIFY_RATIO))
	cmd := exec.Command("sh", "-c", fmt.Sprintf(AMPLIFY_CMD, input_file, AMPLIFY_RATIO, output_file, input_file))
	_ = cmd.Run()

	return nil
}
