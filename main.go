package main

import (
	"log"
	"os"

	"github.com/znkisoft/zedisDB/server"
)

func main() {
	port := os.Getenv("ZEDIS_PORT")
	if port == "" {
		port = "6379"
	}

	s := server.NewServer()
	log.Fatal(s.ListenAndServe(":" + port))
}
