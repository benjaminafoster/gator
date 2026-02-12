package main

import (
	"fmt"
	"context"
	"html"
)

func handlerAggregate(s *State, cmd Command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Usage: gator agg\n")
	}
	
	feedUrl := "https://www.wagslane.dev/index.xml"
	
	feed, err := fetchFeed(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w\n", err)
	}
	
	fmt.Println(unescapeHTML(feed))
	
	return nil
}

func unescapeHTML(feed *RSSFeed) *RSSFeed {
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	
	for i := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
	}

	return feed
}