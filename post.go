package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type writerFactory interface {
	Get(time.Time) (io.Writer, error)
}

type server struct {
	writerFactory writerFactory
	port          int
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
	logWriter, err := s.writerFactory.Get(time.Now())
	if err != nil {
		log.Println("error writing to file: %v", err)
		w.WriteHeader(500)
		return
	}
	_, err = logWriter.Write(d)
	if err != nil {
		log.Println("error writing to file: %v", err)
		w.WriteHeader(500)
		return
	}
	logWriter.Write([]byte("\n"))
	w.WriteHeader(200)
	log.Println("successful")
	return
}

func startServer(wf writerFactory, port int) {
	s := &server{
		writerFactory: wf,
		port:          port,
	}
	http.HandleFunc("/post", s.processPost)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
