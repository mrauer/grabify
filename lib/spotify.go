package lib

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const (
	SPOTIFY_AUTH_URL           = "https://accounts.spotify.com/api/token"
	SPOTIFY_PLAYLIST_ITEMS_URL = "https://api.spotify.com/v1/playlists/%s/tracks"
	BASIC_HEADER               = "Basic %s"
	BEARER_AUTHORIZATION       = "Bearer %s"
)

type Auth struct {
	Token string `json:"access_token"`
}

type Songs struct {
	Items []struct {
		Track struct {
			Artists []struct {
				Name string `json:"name"`
			} `json:"artists"`
			Name string `json:"name"`
		} `json:"track"`
	} `json:"items"`
}

func SpotifyLogin() (string, error) {
	token := ""

	client_id := os.Getenv("SPOTIFY_CLIENT_ID")
	secret_key := os.Getenv("SPOTIFY_SECRET_KEY")

	authorization := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", client_id, secret_key)))

	client := &http.Client{}
	data := url.Values{}
	data.Add("grant_type", "client_credentials")
	req, err := http.NewRequest("POST", SPOTIFY_AUTH_URL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return token, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	req.Header.Set("Authorization", fmt.Sprintf(BASIC_HEADER, authorization))
	resp, err := client.Do(req)
	if err != nil {
		return token, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return token, err
	}
	var auth Auth
	err = json.Unmarshal(body, &auth)
	if err != nil {
		return token, err
	}
	return auth.Token, nil
}

func GetPlaylistItems(playlist, token string) ([]string, error) {
	ret := []string{}
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf(SPOTIFY_PLAYLIST_ITEMS_URL, playlist), nil)
	if err != nil {
		return ret, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf(BEARER_AUTHORIZATION, token))
	resp, err := client.Do(req)
	if err != nil {
		return ret, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ret, err
	}
	var songs Songs
	err = json.Unmarshal(body, &songs)
	if err != nil {
		return ret, err
	}
	for _, item := range songs.Items {
		ret = append(ret, fmt.Sprintf("%s %s", item.Track.Artists[0].Name, item.Track.Name))
	}
	return ret, nil
}
