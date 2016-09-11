package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type config struct {
	folder  string
	port    int
	baseURL string
}

func main() {
	c := config{}
	flag.StringVar(&c.folder, "dir", ".", "log path")
	flag.StringVar(&c.baseURL, "base", "", "base url override with protocol for generating bookmarklet")
	flag.IntVar(&c.port, "port", 9999, "port of server")
	var showVersionOnly = flag.Bool("version", false, "print version info")
	flag.Parse()

	if *showVersionOnly {
		fmt.Printf("%v\n", version)
		os.Exit(0)
	}
	lf := &logFactory{
		dirPath: c.folder,
		format:  "01-02-2006",
	}
	if err := lf.prepFolder(); err != nil {
		log.Fatal(err)
	}
	startServer(lf, c.port, c.baseURL)

}
