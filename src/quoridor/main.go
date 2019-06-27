package main

import (
	"quoridor/server"
	"quoridor/storage"
)

func main() {
	storage.Init()
	server.Start()
}
