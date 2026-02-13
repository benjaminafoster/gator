package main

import (
	"context"
	"fmt"
	"html"
	"time"
)

func handlerAggregate(s *State, cmd Command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage: gator agg\n")
	}
	
	time_between_reqs := cmd.args[0]
	parsedDuration, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return fmt.Errorf("failed to parse duration: %w\n", err)
	}
	
	fmt.Printf("Collecting feeds every %s...\n", parsedDuration)
	
	ticker := time.NewTicker(parsedDuration)
	for ; ; <- ticker.C {
		scrapeFeeds(s)
	}
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

func scrapeFeeds(s *State) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("failed to determine the next feed to fetch: %w\n", err)
	}
	
	nextFeedID := nextFeed.ID
	
	err = s.db.MarkFeedFetched(context.Background(), nextFeedID)
	if err != nil {
		return fmt.Errorf("failed to mark feed as fetched: %w\n", err)
	}
	
	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w\n", err)
	}
	
	unescapedFeed := unescapeHTML(feed)
	
	fmt.Printf("Printing posts found in %s...\n", unescapedFeed.Channel.Title)
	
	for _, item := range unescapedFeed.Channel.Item {
		fmt.Printf("Found post: %s\n", item.Title)
	}
	
	return nil
}
