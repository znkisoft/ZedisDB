package main

import (
	"os"

	"github.com/znkisoft/zedisDB/server"
)

func main() {
	port := os.Getenv("ZEDIS_PORT")
	if port == "" {
		port = "6379"
	}

	server.ListenAndServe(":" + port)
}
