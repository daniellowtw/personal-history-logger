package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type config struct {
	folder string
	port   int
}

func main() {
	c := config{}
	flag.StringVar(&c.folder, "dir", ".", "log path")
	flag.IntVar(&c.port, "port", 9999, "port of server")
	var showVersionOnly = flag.Bool("version", false, "print version info")
	flag.Parse()

	if *showVersionOnly {
		fmt.Printf("%v\n", version)
		os.Exit(0)
	}
	w, err := prepFolder(c.folder)
	if err != nil {
		log.Fatal(err)
	}
	startServer(w, c.port)

}
