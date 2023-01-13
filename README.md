# GoTwitch

## About

Go bot that speaks messages from twitch chat.

## Usage

```sh
gotwitch -channel <channel id> -speak -lang <lang id>
```

- channel id : the twitch channel id (default theoldcoder)
- speak : set if you want the bot to speak the chat messages (default false)
- lang : language code if 'speak' is set (default fr)

For example : 

```sh
gotwitch -channel theoldcoder -speak -lang fr
```

## Chat commands

The commands are word beginning with a '#' sign. For example :

```
#clap
```

The available commands are :

- #clap to play a clapping crowd (embeded sound)
- #bell to ring a bell (embeded sound)

Anything else muse have a mp3 file in the media folder with the same name as the command. For example :

```
#enya
```

will play the mp3 file named "enya.mp3" in tue media folder.

## Cross compile to Windows

GOOS=windows CGO_ENABLED=1 GOARCH=386 go build
