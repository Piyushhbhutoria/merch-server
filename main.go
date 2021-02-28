package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Piyushhbhutoria/merch-server/config"
	"github.com/Piyushhbhutoria/merch-server/db"
	"github.com/Piyushhbhutoria/merch-server/server"
)

func main() {
	environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	server.Init()
}
