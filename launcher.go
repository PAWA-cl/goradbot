package main

import(
	"flag"

	"./goradbot"
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



