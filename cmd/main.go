package main

import (
	"flag"
	"log"
)

func main() {
	t := mustToken()
}

func mustToken() string {
	token := flag.String(
		name: "token-bot",
		value: "",
		usage: "token for access to telegram-bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not")
	}

	return *token
}