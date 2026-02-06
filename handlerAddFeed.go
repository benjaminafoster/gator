package main

import (
	"fmt"
	"time"
	"context"
	"github.com/google/uuid"
	"github.com/benjaminafoster/gator/internal/database"
)

func handlerAddFeedv1(s *State, cmd Command, user database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("Usage: addFeed <name> <url>\n")
	}

	name := cmd.args[0]
	url := cmd.args[1]

	feed, err := addFeed(s, name, url)
	if err != nil {
		return fmt.Errorf("failed to add feed: %w\n", err)
	}
	
	fmt.Printf("New feed added: %s\n", feed.Name)
	
	followFeed, err := followFeed(s, url)
	if err != nil {
		return err
	}
	
	fmt.Printf("Now following feed: %s\n", followFeed.FeedName)

	return nil
}

func handlerAddFeedv2(s *State, cmd Command, user database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("Usage: addFeed <name> <url>\n")
	}

	feedName := cmd.args[0]
	feedUrl := cmd.args[1]
	
	feedParams := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: feedName,
		Url: feedUrl,
		UserID: user.ID,
	}
	
	feedRow, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("failed to create feed: %w\n", err)
	}
	
	fmt.Printf("New feed added: %s\n", feedRow.Name)
	
	feedFollowParams := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feedRow.ID,
	}
	
	_, err = s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("failed to follow feed: %w\n", err)
	}
	
	fmt.Printf("Current user: %s\n", user.Name)

	return nil
}
