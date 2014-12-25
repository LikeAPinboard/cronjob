package main

import (
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
	if resp, err = http.Get("http://blog.wnotes.net/blog/feed/atom"); err != nil {
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

	if feeds, err := parser.Parse(); err != nil {
		fmt.Printf("%v\n", err)
		return
	} else {
		fmt.Printf("%v\n", feeds)
	}
}
