package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type Entry struct {
	Meta      *meta
	Timestamp int64
	URL       string
}

func f(m *meta, url string) error {
	if url == "" {
		return fmt.Errorf("url cannot be empty")
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	z := html.NewTokenizer(resp.Body)
	defer resp.Body.Close()
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return z.Err()
		case html.StartTagToken:
			t := z.Token()
			g(&t, m)
		}
	}
	return nil
}

type meta struct {
	Author      string
	Title       string
	Description string
	Keywords    []string
}

func g(t *html.Token, m *meta) {
	if t.Data == "meta" {
		md := make(map[string]string)
		for _, d := range t.Attr {
			md[d.Key] = d.Val
		}
		if md["name"] == "" {
			return
		}
		switch md["name"] {
		case "author":
			m.Author = md["content"]
		case "description":
			m.Description = md["content"]
		case "title":
			m.Title = md["content"]
		case "keywords":
			kws := strings.Split(md["content"], ",")
			for _, kw := range kws {
				m.Keywords = append(m.Keywords, strings.TrimSpace(kw))
			}
		}
	}
}
