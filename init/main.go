package main

import (
	"flag"

	"github.com/PAWA-cl/goradbot"
)

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
}

var token string
var buffer = make([][]byte, 0)

func main() {
	goradbot.Start(token)
}
