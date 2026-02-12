package main

import (
	"fmt"
	"time"
	"context"
	"github.com/google/uuid"
	"github.com/benjaminafoster/gator/internal/database"
)


func handlerFollow(s *State, cmd Command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage: gator follow <feed_url>\n")
	}
	
	feedUrl := cmd.args[0]
	
	feedFollow, err := followFeed(s, feedUrl)
	if err != nil {
		return err
	}
	
	fmt.Printf("Feed name: %s\n", feedFollow.FeedName)
	fmt.Printf("Current user: %s\n", feedFollow.UserName)
	
	return nil
} 

func handlerFollowv2(s *State, cmd Command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage: gator follow <feed_url>\n")
	}
	
	feedUrl := cmd.args[0]
	
	feedRow, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("Error fetching feed: %w\n", err)
	}
	
	followParams := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feedRow.ID,
	}
	
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), followParams)
	if err != nil {
		return err
	}
	
	fmt.Printf("Feed name: %s\n", feedFollow.FeedName)
	fmt.Printf("Current user: %s\n", feedFollow.UserName)
	
	return nil
}