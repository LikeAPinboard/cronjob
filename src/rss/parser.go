package rss

import (
	"encoding/xml"
	"errors"
)

type RssParser interface {
	Parse()
	GetEntries() (entries []RssEntry)
}

type RssEntry struct {
	Title string
	Url   string
	Date  string
}

type Rss1Type struct {
	XMLName xml.Name `xml:"RDF"`
}

type Rss2Type struct {
	XMLName xml.Name `xml:"rss"`
}
type AtomType struct {
	XMLName xml.Name `xml:"feed"`
}

func NewParser(buffer []byte) (parser RssParser, err error) {
	rss1 := Rss1Type{}
	if err := xml.Unmarshal(buffer, &rss1); err == nil && rss1.XMLName.Space != "" {
		return NewRss1Parser(buffer), nil
	}

	rss2 := Rss2Type{}
	if err := xml.Unmarshal(buffer, &rss2); err == nil && rss2.XMLName.Space != "" {
		return NewRss2Parser(buffer), nil
	}

	atom := AtomType{}
	if err := xml.Unmarshal(buffer, &atom); err == nil && atom.XMLName.Space != "" {
		return NewAtomParser(buffer), nil
	}

	return nil, errors.New("Cannot parse RSS. Invalid XML format?")
}
