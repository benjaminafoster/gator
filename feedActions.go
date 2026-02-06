package main

import (
	"fmt"
	"encoding/xml"
	"net/http"
	"io"
	"context"
	"time"
	"github.com/benjaminafoster/gator/internal/database"
	"github.com/google/uuid"
)


type RSSFeed struct {
	Channel struct {
		Title string `xml:"title"`
		Link string `xml:"link"`
		Description string `xml:"description"`
		Item []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title string `xml:"title"`
	Link string `xml:"link"`
	Description string `xml:"description"`
	PubDate string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w\n", err)
	}
	
	req.Header.Set("User-Agent", "gator")
	c := http.Client{
		Timeout: time.Second * 10,
	}
	
	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch feed: %w\n", err)
	}
	
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d\n", resp.StatusCode)
	}
	
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w\n", err)
	}
	
	feed := RSSFeed{}
	
	err = xml.Unmarshal(data, &feed)
	if err != nil {
		return nil, fmt.Errorf("failed to parse XML: %w\n", err)
	}
	
	return &feed, nil
}

func addFeed(s *State, name string, url string) (database.Feed, error) {
	currentUser := s.cfg.CurrentUser
	dbUser, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return database.Feed{}, fmt.Errorf("failed to get user: %w\n", err)
	}
	feedParams := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
		Url: url,
		UserID: dbUser.ID,
	}
	
	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return database.Feed{},fmt.Errorf("failed to create feed: %w\n", err)
	}
	
	
	return feed, nil
}

func followFeed(s *State, url string) (database.CreateFeedFollowRow, error) {
	currentUsername := s.cfg.CurrentUser
	currentUser, err := s.db.GetUser(context.Background(), currentUsername)
	if err != nil {
		return database.CreateFeedFollowRow{}, fmt.Errorf("Failed to get current user: %w", err)
	}
	
	currentUserId := currentUser.ID
	
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return database.CreateFeedFollowRow{}, fmt.Errorf("Failed to get feed by URL: %w", err)
	}
	
	feedID := feed.ID
	
	followParams := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: currentUserId,
		FeedID: feedID,
	}
	
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), followParams)
	if err != nil {
		return database.CreateFeedFollowRow{}, fmt.Errorf("Failed to create feed follow: %w", err)
	}
	
	return feedFollow, nil
}