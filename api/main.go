package main

import (
	"github.com/MahiroWatanabe/reversetodo/db"
	"github.com/MahiroWatanabe/reversetodo/server"
)

func main() {
	db.Init()
	server.Init()
}