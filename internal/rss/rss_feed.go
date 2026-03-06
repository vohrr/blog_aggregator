package rss

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	var rss_feed RSSFeed
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(data, &rss_feed)
	if err != nil {
		return nil, err
	}
	html.UnescapeString(rss_feed.Channel.Title)
	html.UnescapeString(rss_feed.Channel.Description)
	for _, item := range rss_feed.Channel.Item {
		html.UnescapeString(item.Title)
		html.UnescapeString(item.Description)
	}
	return &rss_feed, nil
}
