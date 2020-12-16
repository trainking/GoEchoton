package main

import (
	"GoEchoton/api/server"
)

func main() {
	s := server.New()
	s.Start()
}
