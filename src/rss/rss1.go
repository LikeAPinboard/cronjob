package rss

import (
	"encoding/xml"
	_ "fmt"
	_ "time"
)

type Rss1Data struct {
	XMLName xml.Name `xml:"RDF"`
	Title   []string `xml:"item>title"`
	Url     []string `xml:"item>link"`
	Date    []string `xml:"item>date"`
}

type Rss1 struct {
	xml     []byte
	entries []RssEntry

	RssParser
}

func NewRss1Parser(buffer []byte) *Rss1 {
	return &Rss1{
		xml:     buffer,
		entries: []RssEntry{},
	}
}

func (r *Rss1) Parse() {

}

func (r *Rss1) GetEntries() []RssEntry {
	return r.entries

}
