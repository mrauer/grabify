# Grabify

<!-- PROJECT LOGO -->
<p align="center">
  <a href="https://github.com/mrauer/grabify">
    <img src="images/logo.png" alt="Logo">
  </a>

  <h3 align="center">Grabify</h3>

  <p align="center">
    Download Spotify playlists as MP3s.
    <br />
    <br />
    <a href="https://github.com/mrauer/grabify/issues">Report Bug</a>
    ·
    <a href="https://github.com/mrauer/grabify/issues">Request Feature</a>
  </p>
</p>

<!-- TABLE OF CONTENTS -->
## Table of Contents

* [Commands](#commands)
  * [Configuration](#configuration)
  * [Download a single song](#download-a-single-song)
  * [Download a playlist from Spotify](#download-a-playlist-from-spotify)
* [License](#license)

<!-- COMMANDS -->
## Commands

<!-- CONFIGURATION -->
### Configuration

In order to properly work, the software needs at least the following environment variable to perform searches on YouTube:

```sh
YOUTUBE_API_KEY=AIzaSyC6cy-xQwZPsqZINASpj2MrBgBWxxxxxxx

```

**youtube-dl** should be installed as well. See [youtube-dl](https://ytdl-org.github.io/youtube-dl/download.html)


If you would like to download a playlist from Spotify, then you also need these additional keys:

```sh
SPOTIFY_CLIENT_ID=e41349c5b8fb4d698389128fxxxxxxxx
SPOTIFY_SECRET_KEY=536f9cb066f8415cb13de240xxxxxxxx

```

Those would be provided to you if you create a Spotify developer account.

<!-- DOWNLOAD A SINGLE SONG -->
### Download a single song

Grabify gives you the ability to download a single song by default when running the software. Here is the type of output you will get.

```sh
./grabify

 ██████  ██████   █████  ██████  ██ ███████ ██    ██ 
██       ██   ██ ██   ██ ██   ██ ██ ██       ██  ██  
██   ███ ██████  ███████ ██████  ██ █████     ████   
██    ██ ██   ██ ██   ██ ██   ██ ██ ██         ██    
 ██████  ██   ██ ██   ██ ██████  ██ ██         ██    
                                                   
version v0.0.3

What song would you want to download? toto africa live

0 - Toto - Africa (Live)
1 - Toto - Africa
2 - TOTO - Africa live 1982
3 - Toto - Africa Live 35th Anniversary
4 - Toto - Africa (1991) [HD] (With CC)

Choose between these songs: 2

Downloading CgUhp0vyz-Q

```

By picking a song from the list, the song will be downloaded to your computer.

<!-- DOWNLOAD A PLAYLIST FROM SPOTIFY -->
### Download a playlist from Spotify

Now it is also possible to download multiple songs at once from a Spotify playlist. We assume you know where to find a playlist ID.

Here is how you perform a bulk download:

```sh
./grabify playlist 6GWZfZSfrJ1JvsLMg4QVa5

 ██████  ██████   █████  ██████  ██ ███████ ██    ██ 
██       ██   ██ ██   ██ ██   ██ ██ ██       ██  ██  
██   ███ ██████  ███████ ██████  ██ █████     ████   
██    ██ ██   ██ ██   ██ ██   ██ ██ ██         ██    
 ██████  ██   ██ ██   ██ ██████  ██ ██         ██    
                                                   
version v0.0.2

Playlist track Tryo Pas pareil XXV
Downloading pl4X0EMxAjQ
Playlist track Oldelaf et Monsieur D. Le café
Downloading 5Y7ZZsOS4O4
```

This is currently synchronous so it might take a little bit of time depending on the size of your playlist.

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.
