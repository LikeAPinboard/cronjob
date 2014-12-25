package rss

import (
	"encoding/xml"
	"errors"
	"time"
)

type Rss2Data struct {
	XMLName xml.Name `xml:"rss"`
	Title   []string `xml:"channel>item>title"`
	Url     []string `xml:"channel>item>link"`
	Date    []string `xml:"channel>item>pubDate"`
}

type Rss2 struct {
	xml []byte

	RssParser
}

func NewRss2Parser(buffer []byte) *Rss2 {
	return &Rss2{
		xml: buffer,
	}
}

func (r *Rss2) Parse() (feed []RssEntry, err error) {
	data := Rss2Data{}
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

func (r *Rss2) parseDateFormat(date string) string {
	format := "Mon, 02 Jan 2006 15:04:05 -0700"
	t, _ := time.Parse(format, date)

	return t.Format("2006-01-02 15:03:04")
}
