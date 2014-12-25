package rss

import (
	"encoding/xml"
	"errors"
	"time"
)

type Rss1Data struct {
	XMLName xml.Name `xml:"RDF"`
	Title   []string `xml:"item>title"`
	Url     []string `xml:"item>link"`
	Date    []string `xml:"item>date"`
}

type Rss1 struct {
	xml []byte

	RssParser
}

func NewRss1Parser(buffer []byte) *Rss1 {
	return &Rss1{
		xml: buffer,
	}
}

func (r *Rss1) Parse() (feed []RssEntry, err error) {
	data := Rss1Data{}
	if err := xml.Unmarshal(r.xml, &data); err != nil {
		return nil, errors.New("RSS Parse Error")
	}

	//var title, url, date string
	for index, title := range data.Title {
		feed = append(feed, RssEntry{
			Title: title,
			Url:   data.Url[index],
			Date:  r.parseDateFormat(data.Date[index]),
		})
	}

	return feed, nil
}

func (r *Rss1) parseDateFormat(date string) string {
	format := "2006-01-02T15:04:05-07:00"
	t, _ := time.Parse(format, date)

	return t.Format("2006-01-02 15:03:04")
}
