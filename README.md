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

Commands are words beginning with a '#' sign. 
For example :

```
#clap
```

The available commands are :

- #speak to turn on/off text to speech
- #clap to play the sound of a clapping crowd (embeded sound)
- #bell to ring a bell (embeded sound)

Anything else must have a mp3 file in the media folder with the same name as the command. 
For example :

```
#enya
```

will play the mp3 file named "enya.mp3" from the media folder.

## Cross compile to Windows

GOOS=windows CGO_ENABLED=1 GOARCH=386 go build
