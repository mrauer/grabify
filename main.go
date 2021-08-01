package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mrauer/grabify/lib"
)

const (
	SOFTWARE_LOGO = `
 ██████  ██████   █████  ██████  ██ ███████ ██    ██ 
██       ██   ██ ██   ██ ██   ██ ██ ██       ██  ██  
██   ███ ██████  ███████ ██████  ██ █████     ████   
██    ██ ██   ██ ██   ██ ██   ██ ██ ██         ██    
 ██████  ██   ██ ██   ██ ██████  ██ ██         ██    
                                                   `
	SOFTWARE_VERSION       = "v0.0.2"
	DOWNLOAD_PATH          = "data"
	PLAYLIST_KW            = "playlist"
	PLAYLIST_PAUSE_SECONDS = 10
)

var songs map[int]string

func main() {
	fmt.Println(SOFTWARE_LOGO)
	fmt.Println(fmt.Sprintf("version %s\n", SOFTWARE_VERSION))
	// Spotify playlist.
	if len(os.Args) == 3 {
		if os.Args[1] == PLAYLIST_KW {
			playlist := os.Args[2]
			token, err := lib.SpotifyLogin()
			if err != nil {
				fmt.Println(err)
			}
			tracks, err := lib.GetPlaylistItems(playlist, token)
			if err != nil {
				fmt.Println(err)
			}
			for _, track := range tracks {
				fmt.Println(fmt.Sprintf("Playlist track %s", track))
				go func(track string) {
					choices, _ := lib.YoutubeSearch(track, false)
					if err != nil {
						return
					}
					lib.DownloadTrack(choices[0])
					lib.AmplifyAudio(choices[0])
				}(track)
				time.Sleep(PLAYLIST_PAUSE_SECONDS * time.Second)
			}
			return
		}
	}
	// Default route.
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What song would you want to download? ")
	query, _ := reader.ReadString('\n')
	fmt.Println("")
	choices, err := lib.YoutubeSearch(query, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("")
	fmt.Print("Choose between these songs: ")
	song_id_str, _ := reader.ReadString('\n')
	song_id, err := strconv.Atoi(strings.TrimSuffix(song_id_str, "\n"))
	fmt.Println("")
	lib.DownloadTrack(choices[song_id])
	lib.AmplifyAudio(choices[song_id])
	return
}
