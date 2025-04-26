package main

import (
	"context"
	"encoding/xml"
	"fmt"
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

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("feed fetch request failed: %s", err)
	}

	request.Header.Set("User-Agent", "gator")

	client := http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error completing http request: %w", err)
	}

	//fmt.Printf("response: %v", res)

	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error reading response body: %w", err)
	}


	feed := RSSFeed{}
	err = xml.Unmarshal(dat, &feed)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error unmarshaling json response: %w", err)
	}
	
	//fmt.Printf("%+v", feed)

	// escape html in title and description fields of general feed
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	// escape html in title and description fields of each item within the feed
	for i := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
	}

	return &feed, nil
}