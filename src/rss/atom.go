package rss

import (
	"encoding/xml"
	_ "fmt"
	_ "time"
)

type AtomData struct {
	XMLName xml.Name   `xml:"feed"`
	Title   []string   `xml:"entry>title"`
	Url     []AtomAttr `xml:"entry>link"`
	Date    []string   `xml:"entry>updated"`
}

type AtomAttr struct {
	Key   string `xml:"href,attr"`
	Value string `xml:",chardata"`
}

type Atom struct {
	xml     []byte
	entries []RssEntry

	RssParser
}

func NewAtomParser(buffer []byte) *Atom {
	return &Atom{
		xml:     buffer,
		entries: []RssEntry{},
	}
}

func (a *Atom) Parse() {

}

func (a *Atom) GetEntries() []RssEntry {
	return a.entries
}
