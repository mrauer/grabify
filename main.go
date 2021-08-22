package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

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
	SOFTWARE_VERSION       = "v0.0.3"
	PLAYLIST_KW            = "playlist"
	PLAYLIST_ROUTE_NUM_ARG = 3
)

var songs map[int]string

func main() {
	fmt.Println(SOFTWARE_LOGO)
	fmt.Println(fmt.Sprintf("version %s\n", SOFTWARE_VERSION))
	// Spotify playlist.
	if len(os.Args) == PLAYLIST_ROUTE_NUM_ARG {
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

				choices, err := lib.YoutubeSearch(track, false)
				if err != nil {
					return
				}

				if len(choices) == 0 {
					return
				}

				lib.DownloadTrack(choices[0])
				lib.AmplifyAudio(choices[0])
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

	if len(choices) == 0 {
		fmt.Println("You must have reached the YouTube API limit. Please try again later.")

		return
	}

	fmt.Println("")
	fmt.Print("Choose between these songs: ")

	song_id_str, _ := reader.ReadString('\n')
	song_id, _ := strconv.Atoi(strings.TrimSuffix(song_id_str, "\n"))

	fmt.Println("")
	lib.DownloadTrack(choices[song_id])
	lib.AmplifyAudio(choices[song_id])

	return
}
