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

## Cross compile to Windows

GOOS=windows CGO_ENABLED=1 GOARCH=386 go build
