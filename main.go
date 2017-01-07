package main

import (
	"log"
	"net/http"

	"github.com/FriendsOfApi/fake-twitter-api/api"
	"github.com/sagikazarmark/bingo"
	flag "github.com/spf13/pflag"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8080", "Custom port")
}

func main() {
	bin, _ := bingo.New("fake-twitter-api", "Fake Twitter API", "Provides a simple example for FriendsOfApi/Boilerplate")

	api.Stats(bin)

	log.Println("Starting server on *:" + port)
	log.Fatal(http.ListenAndServe(":"+port, bingo.NewHandler(bin)))
}
