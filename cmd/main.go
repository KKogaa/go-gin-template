package main

import (
	"github.com/webtoon/cmd/server"
)

func main() {
	s := server.NewServer()
	s.Start()
}
