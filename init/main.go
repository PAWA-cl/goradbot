package main

import (
	"os"

	"github.com/PAWA-cl/goradbot"
)

var token string

func init() {
	token = os.Getenv("DISCORD_TOKEN")
}

func main() {
	goradbot.Start(token)
}
