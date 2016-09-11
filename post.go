package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type server struct {
	w    io.Writer
	port int
}

func (s *server) processPost(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	url := q.Get("url")
	if url == "" {
		log.Println("empty url")
		w.WriteHeader(401)
		return
	}
	m := &meta{}
	if err := f(m, url); err != io.EOF {
		log.Println("error processing url: %v", err)
		w.WriteHeader(500)
		return
	}
	e := Entry{
		Meta:      m,
		URL:       url,
		Timestamp: time.Now().Unix(),
	}
	d, err := json.Marshal(e)
	if err != nil {
		log.Fatal(err)
		log.Println("error serializing: %v", err)
		w.WriteHeader(500)
		return
	}
	_, err = s.w.Write(d)
	s.w.Write([]byte("\n"))
	if err != nil {
		log.Println("error writing to file: %v", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	log.Println("successful")
	return
}

func startServer(w io.Writer, port int) {
	s := &server{
		w:    w,
		port: port,
	}
	http.HandleFunc("/post", s.processPost)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
