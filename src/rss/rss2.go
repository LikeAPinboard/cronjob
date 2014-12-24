package rss

import (
	"encoding/xml"
	_ "fmt"
	_ "time"
)

type Rss2Data struct {
	XMLName xml.Name `xml:"RDF"`
	Title   []string `xml:"channel>item>title"`
	Url     []string `xml:"channel>item>link"`
	Date    []string `xml:"channel>item>pubDate"`
}

type Rss2 struct {
	xml     []byte
	entries []RssEntry

	RssParser
}

func NewRss2Parser(buffer []byte) *Rss2 {
	return &Rss2{
		xml:     buffer,
		entries: []RssEntry{},
	}
}

func (r *Rss2) Parse() {

}

func (r *Rss2) GetEntries() []RssEntry {
	return r.entries
}
