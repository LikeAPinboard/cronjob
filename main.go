package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"rss"
)

func main() {
	var (
		resp   *http.Response
		err    error
		buffer []byte
		parser rss.RssParser
	)

	// test RSS
	if resp, err = http.Get("http://blog.wnotes.net/blog/feed/rss"); err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	defer resp.Body.Close()

	if buffer, err = ioutil.ReadAll(resp.Body); err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	if parser, err = rss.NewParser(buffer); err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	parser.Parse()

	// get entries
	for _, item := range parser.GetEntries() {
		// TODO: Implement
		fmt.Printf("%v\n", item)
	}
}
