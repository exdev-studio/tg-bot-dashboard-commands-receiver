package main

import (
	"flag"
	"log"

	"github.com/exdev-studio/tg-bot-dashboard-commands-receiver/internal/app/receiver"
)

var (
	bindAddr string
)

func init() {
	flag.StringVar(&bindAddr, "bind-addr", "0.0.0.0:8080", "will be used as an address for listening to requests; format 0.0.0.0:8080")
}

func main() {
	flag.Parse()

	c := &receiver.Config{
		BindAddr: bindAddr,
	}

	if err := receiver.Start(c); err != nil {
		log.Fatal(err)
	}
}
