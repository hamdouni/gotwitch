# GoTwitch

## About

Go bot that speaks messages from twitch chat.

## Usage

```sh
gotwitch -channel <channel id> -lang <lang id>
```

For example : 

```sh
gotwitch -channel theoldcoder -lang fr
```

## Cross compile to Windows

GOOS=windows CGO_ENABLED=1 GOARCH=386 go build
