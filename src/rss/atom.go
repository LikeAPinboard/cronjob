package rss

import (
	"encoding/xml"
	"errors"
	"time"
)

type AtomData struct {
	XMLName xml.Name   `xml:"feed"`
	Title   []string   `xml:"entry>title"`
	Url     []AtomAttr `xml:"entry>link"`
	Date    []string   `xml:"entry>updated"`
}

type AtomAttr struct {
	Value string `xml:"href,attr"`
}

type Atom struct {
	xml []byte

	RssParser
}

func NewAtomParser(buffer []byte) *Atom {
	return &Atom{
		xml: buffer,
	}
}

func (r *Atom) Parse() (feed []RssEntry, err error) {
	data := AtomData{}
	if err := xml.Unmarshal(r.xml, &data); err != nil {
		return nil, errors.New("RSS Parse Error")
	}

	//var title, url, date string
	for index, title := range data.Title {
		feed = append(feed, RssEntry{
			Title: title,
			Url:   r.parseUrlAttribute(data.Url[index]),
			Date:  r.parseDateFormat(data.Date[index]),
		})
	}

	return feed, nil
}

func (r *Atom) parseDateFormat(date string) string {
	format := "2006-01-02T15:04:05-07:00"
	t, _ := time.Parse(format, date)

	return t.Format("2006-01-02 15:03:04")
}

func (r *Atom) parseUrlAttribute(attr AtomAttr) string {
	return attr.Value
}
